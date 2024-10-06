package usecase

import (
	"context"

	ports "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/repository"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

type CustomerFindByCPFUseCase struct {
	repository ports.CustomerRepository
}

func NewCustomerFindByCPF(repository ports.CustomerRepository) *CustomerFindByCPFUseCase {
	return &CustomerFindByCPFUseCase{
		repository: repository,
	}
}

// RegisterCustomer registra um novo cliente.
func (cr *CustomerFindByCPFUseCase) FindCustomerByCPF(ctx context.Context, cpf string) (*entity.Customer, error) {

	customer, err := cr.repository.Find(ctx, cpf)
	if err != nil {
		return nil, err
	}

	return customer, nil
}
