package dto

import "github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"

type CreateProductDTO struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
}

type ProductDTO struct {
	ID          string  `json:"id,omitempty"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
}

func (dto *ProductDTO) ToEntity() (*entity.Product, error) {
	return &entity.Product{
		ID:          dto.ID,
		Name:        dto.Name,
		Description: dto.Description,
		Category:    dto.Category,
		Price:       dto.Price,
	}, nil
}

func (dto *ProductDTO) FromEntity(product entity.Product) {
	dto.ID = product.ID
	dto.Name = product.Name
	dto.Description = product.Description
	dto.Category = product.Category
	dto.Price = product.Price
}
