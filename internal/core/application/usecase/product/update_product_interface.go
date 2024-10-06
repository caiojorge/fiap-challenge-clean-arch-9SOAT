package usecase

import (
	"context"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

type UpdateProductUseCase interface {
	UpdateProduct(ctx context.Context, customer entity.Product) error
}
