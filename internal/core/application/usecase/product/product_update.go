package usecase

import (
	"context"
	"errors"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
	portsrepository "github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/repository"
)

type ProductUpdateUseCase struct {
	repository portsrepository.ProductRepository
}

func NewProductUpdate(repository portsrepository.ProductRepository) *ProductUpdateUseCase {
	return &ProductUpdateUseCase{
		repository: repository,
	}
}

// UpdateProduct atualiza um novo produto.
func (cr *ProductUpdateUseCase) UpdateProduct(ctx context.Context, product entity.Product) error {

	prd, err := cr.repository.Find(ctx, product.GetID())
	if err != nil {
		return err
	}

	if prd == nil {
		return errors.New("product not found")
	}

	err = cr.repository.Update(ctx, &product)
	if err != nil {
		return err
	}

	return nil
}
