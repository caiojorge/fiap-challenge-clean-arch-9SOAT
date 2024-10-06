package usecase

import (
	"context"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

type FindAllProductsUseCase interface {
	FindAllProducts(ctx context.Context) ([]*entity.Product, error)
}
