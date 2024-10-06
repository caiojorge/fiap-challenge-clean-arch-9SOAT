package portsusecase

import (
	"context"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

type FindAllOrderUseCase interface {
	FindAllOrders(ctx context.Context) ([]*entity.Order, error)
}
