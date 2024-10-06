package usecase

import (
	"context"

	ports "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/repository"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

type CustomerFindAllUseCase struct {
	repository ports.CustomerRepository
}

func NewCustomerFindAll(repository ports.CustomerRepository) *CustomerFindAllUseCase {
	return &CustomerFindAllUseCase{
		repository: repository,
	}
}

// RegisterCustomer registra um novo cliente.
func (cr *CustomerFindAllUseCase) FindAllCustomers(ctx context.Context) ([]*entity.Customer, error) {

	customers, err := cr.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return customers, nil
}
