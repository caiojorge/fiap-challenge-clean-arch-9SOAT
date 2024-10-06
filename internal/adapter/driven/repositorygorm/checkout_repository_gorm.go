package repositorygorm

import (
	"context"
	"errors"

	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driven/model"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type CheckoutRepositoryGorm struct {
	DB *gorm.DB
}

func NewCheckoutRepositoryGorm(db *gorm.DB) *CheckoutRepositoryGorm {
	return &CheckoutRepositoryGorm{
		DB: db,
	}
}

// Create creates a new checkcout. It returns an error if something goes wrong.
func (r *CheckoutRepositoryGorm) Create(ctx context.Context, entity *entity.Checkout) error {
	var model model.Checkout
	err := copier.Copy(&model, entity)
	if err != nil {
		return err
	}

	if err := r.DB.Create(&model).Error; err != nil {
		return err
	}

	return nil
}

// Update updates the checkout. It returns an error if something goes wrong.
func (r *CheckoutRepositoryGorm) Update(ctx context.Context, entity *entity.Checkout) error {

	var model model.Checkout
	copier.Copy(&model, entity)

	return r.DB.Save(model).Error
}

// Find checkout by id
func (r *CheckoutRepositoryGorm) Find(ctx context.Context, id string) (*entity.Checkout, error) {
	var orderModel model.Checkout
	result := r.DB.Find(&orderModel, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	var entity *entity.Checkout
	err := copier.Copy(&entity, &orderModel)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

// FindAll not implemented
func (r *CheckoutRepositoryGorm) FindAll(ctx context.Context) ([]*entity.Checkout, error) {
	// var mOrders []model.Order

	// result := r.DB.Preload("Items").Order("created_at desc").Find(&mOrders)
	// if result.Error != nil {
	// 	return nil, result.Error
	// }

	// var eOrders []*entity.Order

	// for _, mOrder := range mOrders {
	// 	eOrder := r.converter.ToEntity(&mOrder)
	// 	eOrders = append(eOrders, eOrder)
	// }

	// return eOrders, nil
	return nil, nil
}

// Delete not implemented
func (r *CheckoutRepositoryGorm) Delete(ctx context.Context, id string) error {
	// var orderModel model.Order
	// result := r.DB.Model(&model.Order{}).Where("id = ?", id).First(&orderModel)
	// if result.Error != nil {
	// 	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
	// 		fmt.Println("repositorygorm: order not found")
	// 		return nil
	// 	}
	// 	return result.Error
	// }

	// result = r.DB.Delete(&orderModel)
	// if result.Error != nil {
	// 	return result.Error
	// }

	return nil
}

func (r *CheckoutRepositoryGorm) FindbyOrderID(ctx context.Context, id string) (*entity.Checkout, error) {
	var orderModel model.Checkout
	result := r.DB.Find(&orderModel, "order_id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	if orderModel.ID == "" {
		return nil, errors.New("checkout not found")
	}

	var entity *entity.Checkout
	err := copier.Copy(&entity, &orderModel)
	if err != nil {
		return nil, err
	}

	return entity, nil
}
