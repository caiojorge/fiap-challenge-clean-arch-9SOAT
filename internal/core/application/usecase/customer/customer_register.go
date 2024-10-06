package usecase

import (
	"context"
	"errors"
	"fmt"

	portsrepository "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/repository"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

type CustomerRegisterUseCase struct {
	repository portsrepository.CustomerRepository
}

func NewCustomerRegister(repository portsrepository.CustomerRepository) *CustomerRegisterUseCase {
	return &CustomerRegisterUseCase{
		repository: repository,
	}
}

// RegisterCustomer registra um novo cliente.
func (cr *CustomerRegisterUseCase) RegisterCustomer(ctx context.Context, customer entity.Customer) error {

	fmt.Println("usecase: verifica se o cliente existe: " + customer.GetCPF().Value)
	customerFound, err := cr.repository.Find(ctx, customer.GetCPF().Value)
	if err != nil && err.Error() != "customer not found" {
		fmt.Println("usecase: err: " + err.Error())
		return err
	}

	if customerFound != nil {
		fmt.Println("usecase: Cliente já existe: " + customer.GetCPF().Value)
		return errors.New("customer already exists")
	}

	fmt.Println("usecase: Criando cliente: " + customer.GetCPF().Value)
	// Cria o cliente
	err = cr.repository.Create(ctx, &customer)
	if err != nil {
		return err
	}

	return nil
}
