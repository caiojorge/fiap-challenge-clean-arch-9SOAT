package usecase

import (
	"context"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

type FindProductByCategoryUseCase interface {
	FindProductByCategory(ctx context.Context, category string) ([]*entity.Product, error)
}
