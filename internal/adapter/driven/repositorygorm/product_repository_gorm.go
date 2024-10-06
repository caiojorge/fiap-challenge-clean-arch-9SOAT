package repositorygorm

import (
	"context"
	"errors"
	"fmt"

	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driven/converter"
	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driven/model"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
	"gorm.io/gorm"
)

type ProductRepositoryGorm struct {
	DB        *gorm.DB
	converter converter.Converter[entity.Product, model.Product]
}

func NewProductRepositoryGorm(db *gorm.DB, converter converter.Converter[entity.Product, model.Product]) *ProductRepositoryGorm {
	return &ProductRepositoryGorm{
		DB:        db,
		converter: converter,
	}
}

// Create creates a new product. It returns an error if something goes wrong.
func (r *ProductRepositoryGorm) Create(ctx context.Context, entity *entity.Product) error {
	fmt.Println("repositorygorm: Criando produto: " + entity.GetID())
	model := r.converter.FromEntity(entity)

	// o id é gerado no momento da gravação
	// id := shared.NewIDGenerator()
	// entity.RedifneID(id)
	// model.ID = id

	return r.DB.Create(model).Error
}

func (r *ProductRepositoryGorm) Update(ctx context.Context, entity *entity.Product) error {

	result := r.DB.Save(r.converter.FromEntity(entity))
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *ProductRepositoryGorm) Find(ctx context.Context, id string) (*entity.Product, error) {
	var productModel model.Product
	fmt.Println("repositorygorm: Find product: " + id)
	result := r.DB.Model(&model.Product{}).Where("id = ?", id).First(&productModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	entity := r.converter.ToEntity(&productModel)

	return entity, nil
}

func (r *ProductRepositoryGorm) FindAll(ctx context.Context) ([]*entity.Product, error) {
	var mProducts []model.Product
	result := r.DB.Find(&mProducts)
	if result.Error != nil {
		return nil, result.Error
	}

	var eProducts []*entity.Product

	for _, mProduct := range mProducts {
		eProduct := r.converter.ToEntity(&mProduct)
		eProducts = append(eProducts, eProduct)
	}

	return eProducts, nil
}

func (r *ProductRepositoryGorm) Delete(ctx context.Context, id string) error {
	var productModel model.Product
	result := r.DB.Model(&model.Product{}).Where("id = ?", id).First(&productModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Println("repositorygorm: product not found")
			return nil
		}
		return result.Error
	}

	result = r.DB.Delete(&productModel)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *ProductRepositoryGorm) FindByName(ctx context.Context, name string) (*entity.Product, error) {
	var productModel model.Product
	result := r.DB.Model(&model.Product{}).Where("name = ?", name).First(&productModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	entity := r.converter.ToEntity(&productModel)

	return entity, nil
}

func (r *ProductRepositoryGorm) FindByCategory(ctx context.Context, category string) ([]*entity.Product, error) {
	var productModel []model.Product
	result := r.DB.Model(&model.Product{}).Where("category = ?", category).Find(&productModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	var entities []*entity.Product
	for _, product := range productModel {
		entity := r.converter.ToEntity(&product)
		entities = append(entities, entity)
	}

	return entities, nil
}
