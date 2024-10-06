package usecase

import (
	"context"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
	ports "github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/repository"
)

type ProductFindByCategoryUseCase struct {
	repository ports.ProductRepository
}

func NewProductFindByCategory(repository ports.ProductRepository) *ProductFindByCategoryUseCase {
	return &ProductFindByCategoryUseCase{
		repository: repository,
	}
}

func (cr *ProductFindByCategoryUseCase) FindProductByCategory(ctx context.Context, category string) ([]*entity.Product, error) {

	products, err := cr.repository.FindByCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	return products, nil
}
