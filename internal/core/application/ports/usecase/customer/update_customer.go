package portsusecase

import (
	"context"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

type UpdateCustomerUseCase interface {
	UpdateCustomer(ctx context.Context, customer entity.Customer) error
}
