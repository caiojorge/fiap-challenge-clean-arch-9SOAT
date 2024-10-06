package converter

import (
	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driven/model"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

type ProductConverter struct{}

func NewProductConverter() *ProductConverter {
	return &ProductConverter{}
}

func (pc *ProductConverter) FromEntity(entity *entity.Product) *model.Product {

	return &model.Product{
		ID:          entity.GetID(),
		Name:        entity.GetName(),
		Description: entity.GetDescription(),
		Price:       entity.GetPrice(),
		Category:    entity.GetCategory(),
	}
}

// TODO: voltar aqui para avaliar se é melhor retornar um erro tbm
func (pc *ProductConverter) ToEntity(model *model.Product) *entity.Product {
	product, _ := entity.ConvertProduct(model.ID, model.Name, model.Description, model.Category, model.Price)
	return product
}
