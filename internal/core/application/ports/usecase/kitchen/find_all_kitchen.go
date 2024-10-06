package portsusecase

import (
	"context"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

type FindAllKitchenUseCase interface {
	FindAllKitchen(ctx context.Context) ([]*entity.Kitchen, error)
}
