package portsusecase

import (
	"context"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

type CreateOrderUseCase interface {
	CreateOrder(ctx context.Context, order *entity.Order) error
}
