package repositorygorm

import (
	"context"
	"errors"
	"fmt"

	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driven/converter"
	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driven/model"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type OrderRepositoryGorm struct {
	DB        *gorm.DB
	converter converter.Converter[entity.Order, model.Order]
}

func NewOrderRepositoryGorm(db *gorm.DB, converter converter.Converter[entity.Order, model.Order]) *OrderRepositoryGorm {
	return &OrderRepositoryGorm{
		DB:        db,
		converter: converter,
	}
}

// Create creates a new product. It returns an error if something goes wrong.
func (r *OrderRepositoryGorm) Create(ctx context.Context, entity *entity.Order) error {
	var model model.Order

	err := copier.Copy(&model, entity)
	if err != nil {
		return err
	}

	if *model.CustomerCPF == "" {
		model.CustomerCPF = nil
	}

	err = r.DB.Create(model).Error
	if err != nil {
		fmt.Println("gorm: " + err.Error())
		return err
	}

	return nil
}

func (r *OrderRepositoryGorm) Update(ctx context.Context, entity *entity.Order) error {

	result := r.DB.Save(r.converter.FromEntity(entity))
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Find retrieves a product by its ID. It returns an error if something goes wrong.
func (r *OrderRepositoryGorm) Find(ctx context.Context, id string) (*entity.Order, error) {
	var orderModel model.Order
	result := r.DB.Preload("Items").Order("created_at desc").Find(&orderModel, "id = ?", id)
	//result := r.DB.Model(&model.Order{}).Where("id = ?", id).First(&orderModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	entity := r.converter.ToEntity(&orderModel)

	return entity, nil
}

func (r *OrderRepositoryGorm) FindAll(ctx context.Context) ([]*entity.Order, error) {
	var mOrders []model.Order

	result := r.DB.Preload("Items").Order("created_at desc").Find(&mOrders)
	if result.Error != nil {
		return nil, result.Error
	}

	var eOrders []*entity.Order

	for _, mOrder := range mOrders {
		eOrder := r.converter.ToEntity(&mOrder)
		eOrders = append(eOrders, eOrder)
	}

	return eOrders, nil
}

func (r *OrderRepositoryGorm) Delete(ctx context.Context, id string) error {
	var orderModel model.Order
	result := r.DB.Model(&model.Order{}).Where("id = ?", id).First(&orderModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Println("repositorygorm: order not found")
			return nil
		}
		return result.Error
	}

	result = r.DB.Delete(&orderModel)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *OrderRepositoryGorm) FindByParams(ctx context.Context, params map[string]interface{}) ([]*entity.Order, error) {

	var orders []*entity.Order
	var models []*model.Order

	query := r.DB.Preload("Items").Order("created_at desc")
	//query := r.DB.Model(&model.Order{})

	// Adiciona condições dinâmicas com base nos parâmetros
	if status, ok := params["status"]; ok {
		query = query.Where("status = ?", status)
	}
	if customerCPF, ok := params["customer_cpf"]; ok {
		query = query.Where("customer_cpf = ?", customerCPF)
	}
	if startDate, ok := params["start_date"]; ok {
		query = query.Where("created_at >= ?", startDate)
	}
	if endDate, ok := params["end_date"]; ok {
		query = query.Where("created_at <= ?", endDate)
	}

	// Executa a consulta
	err := query.Find(&models).Error
	if err != nil {
		return nil, err
	}

	copier.Copy(&orders, &models)

	return orders, err

}
