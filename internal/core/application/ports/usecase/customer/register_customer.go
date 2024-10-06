package portsusecase

import (
	"context"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

type RegisterCustomerUseCase interface {
	RegisterCustomer(ctx context.Context, customer entity.Customer) error
}
