package usecase

import (
	"context"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
	ports "github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/repository"
)

type ProductFindAllUseCase struct {
	repository ports.ProductRepository
}

func NewProductFindAll(repository ports.ProductRepository) *ProductFindAllUseCase {
	return &ProductFindAllUseCase{
		repository: repository,
	}
}

// FindAllProducts busca todos os produtos.
func (cr *ProductFindAllUseCase) FindAllProducts(ctx context.Context) ([]*entity.Product, error) {

	products, err := cr.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}
