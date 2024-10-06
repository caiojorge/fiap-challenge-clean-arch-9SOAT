package usecase

import (
	"context"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

type RegisterProductUseCase interface {
	RegisterProduct(ctx context.Context, customer *entity.Product) error
}
