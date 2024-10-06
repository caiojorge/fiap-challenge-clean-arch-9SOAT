package repositorygorm

import (
	"context"
	"errors"

	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driven/model"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
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

	if err := r.DB.Create(&model).Error; err != nil {
		return err
	}

	return nil
}

// Update updates the checkout. It returns an error if something goes wrong.
func (r *KitchenRepositoryGorm) Update(ctx context.Context, entity *entity.Kitchen) error {

	var model model.Checkout
	copier.Copy(&model, entity)

	return r.DB.Save(model).Error
}

// Find not implemented
func (r *KitchenRepositoryGorm) Find(ctx context.Context, id string) (*entity.Kitchen, error) {

	return nil, nil
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
