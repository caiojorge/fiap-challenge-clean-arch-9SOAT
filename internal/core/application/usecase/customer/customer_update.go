package usecase

import (
	"context"
	"fmt"

	portsrepository "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/repository"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

type CustomerUpdateUseCase struct {
	repository portsrepository.CustomerRepository
}

func NewCustomerUpdate(repository portsrepository.CustomerRepository) *CustomerUpdateUseCase {
	return &CustomerUpdateUseCase{
		repository: repository,
	}
}

// RegisterCustomer registra um novo cliente.
func (cr *CustomerUpdateUseCase) UpdateCustomer(ctx context.Context, customer entity.Customer) error {

	fmt.Println("usecase: verifica se o cliente existe: " + customer.GetCPF().Value)
	c, err := cr.repository.Find(ctx, customer.GetCPF().Value)
	if err != nil {
		fmt.Println("usecase: err: " + err.Error())
		return err
	}

	if c == nil {
		return fmt.Errorf("customer not found")
	}

	fmt.Println("usecase: atualizando cliente: " + customer.GetCPF().Value)
	// Cria o cliente
	err = cr.repository.Update(ctx, &customer)
	if err != nil {
		return err
	}

	return nil
}
