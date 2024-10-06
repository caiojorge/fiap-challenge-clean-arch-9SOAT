package portsusecase

import (
	"context"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

type FindOrderByParamsUseCase interface {
	FindOrdersByParams(ctx context.Context, params map[string]interface{}) ([]*entity.Order, error)
}
