package portsusecase

import (
	"context"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

type FindAllCustomersUseCase interface {
	FindAllCustomers(ctx context.Context) ([]*entity.Customer, error)
}
