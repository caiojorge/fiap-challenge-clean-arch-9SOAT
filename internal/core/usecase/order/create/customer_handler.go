package usecase

import (
	"context"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
	domainRepository "github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/repository"
)

type CustomerHandler struct {
	customerRepository domainRepository.CustomerRepository
}

func NewCustomerHandler(customerRepository domainRepository.CustomerRepository) *CustomerHandler {
	return &CustomerHandler{
		customerRepository: customerRepository,
	}
}

func (ch *CustomerHandler) Handle(ctx context.Context, cpf string) (*entity.Customer, error) {

	// cliente não informado
	if cpf == "" {
		return nil, nil
	}

	// cliente foi informado, e vamos verificar se ele já esta cadastrado
	customer, err := ch.customerRepository.Find(ctx, cpf)
	if err != nil {
		return nil, err
	}

	// Se o cliente já existe, nada mais a ser feito
	if customer != nil {
		return nil, nil
	}

	// Se o cliente não existe, vamos cadastra-lo
	customer, err = entity.NewCustomerWithCPFInformed(cpf)
	if err != nil {
		return nil, err
	}

	// persiste o cliente
	if err = ch.customerRepository.Create(ctx, customer); err != nil {
		return nil, err
	}

	return customer, nil
}
