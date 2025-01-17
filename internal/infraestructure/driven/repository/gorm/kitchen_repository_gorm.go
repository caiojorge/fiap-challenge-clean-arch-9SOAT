package repositorygorm

import (
	"context"
	"errors"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/model"
	sharedDate "github.com/caiojorge/fiap-challenge-ddd/internal/shared/time"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type KitchenRepositoryGorm struct {
	DB *gorm.DB
}

func NewKitchenRepositoryGorm(db *gorm.DB) *KitchenRepositoryGorm {
	return &KitchenRepositoryGorm{
		DB: db,
	}
}

// Create creates a new checkcout. It returns an error if something goes wrong.
func (r *KitchenRepositoryGorm) Create(ctx context.Context, entity *entity.Kitchen) error {
	var model model.Kitchen
	err := copier.Copy(&model, entity)
	if err != nil {
		return err
	}

	model.CreatedAt = sharedDate.GetBRTimeNow()

	if err := r.DB.Create(&model).Error; err != nil {
		return err
	}

	return nil
}

// Find not implemented
func (r *KitchenRepositoryGorm) Update(ctx context.Context, entity *entity.Kitchen) error {
	return nil
}

// Find not implemented
func (r *KitchenRepositoryGorm) Find(ctx context.Context, id string) (*entity.Kitchen, error) {
	var model model.Kitchen
	result := r.DB.Preload("Order").First(&model, "id = ?", id)
	//result := r.DB.Model(&model.Order{}).Where("id = ?", id).First(&orderModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	entity := &entity.Kitchen{
		ID:            model.ID,
		OrderID:       model.OrderID,
		Queue:         model.Queue,
		EstimatedTime: model.EstimatedTime,
		CreatedAt:     model.CreatedAt,
	}

	return entity, nil
}

// FindAll not implemented
func (r *KitchenRepositoryGorm) FindAll(ctx context.Context) ([]*entity.Kitchen, error) {

	var models []model.Kitchen

	result := r.DB.Find(&models)
	if result.Error != nil {
		return nil, result.Error
	}

	if len(models) == 0 {
		return nil, errors.New("kitchens not found")
	}

	var entities []*entity.Kitchen

	copier.Copy(&entities, &models)

	return entities, nil

}

// Delete not implemented
func (r *KitchenRepositoryGorm) Delete(ctx context.Context, id string) error {

	return nil
}

func (r *KitchenRepositoryGorm) FindByParams(ctx context.Context, params map[string]interface{}) ([]*entity.Kitchen, error) {

	var kitchen []*entity.Kitchen
	var models []*model.Kitchen

	query := r.DB.Order("created_at desc")
	//query := r.DB.Model(&model.Order{})

	// Adiciona condições dinâmicas com base nos parâmetros
	if id, ok := params["id"]; ok {
		query = query.Where("id = ?", id)
	}
	if orderID, ok := params["order_id"]; ok {
		query = query.Where("order_id = ?", orderID)
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

	for _, model := range models {
		entity := entity.Kitchen{
			ID:            model.ID,
			OrderID:       model.OrderID,
			Queue:         model.Queue,
			EstimatedTime: model.EstimatedTime,
			CreatedAt:     model.CreatedAt,
		}
		kitchen = append(kitchen, &entity)
	}

	return kitchen, err

}
