package portsusecase

import (
	"context"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

type FindOrderByIDUseCase interface {
	FindOrderByID(ctx context.Context, id string) (*entity.Order, error)
}
