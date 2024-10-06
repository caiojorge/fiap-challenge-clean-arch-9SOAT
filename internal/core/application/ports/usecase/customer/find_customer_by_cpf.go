package portsusecase

import (
	"context"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

type FindCustomerByCPFUseCase interface {
	FindCustomerByCPF(ctx context.Context, cpf string) (*entity.Customer, error)
}
