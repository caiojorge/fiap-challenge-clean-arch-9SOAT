package usecase

import (
	"context"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

type FindProductByIDUseCase interface {
	FindProductByID(ctx context.Context, id string) (*entity.Product, error)
}
