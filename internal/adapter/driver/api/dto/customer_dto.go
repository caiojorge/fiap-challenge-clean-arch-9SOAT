package dto

import (
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/valueobject"
)

type UpdateCustomerDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CustomerDTO struct {
	CPF   string `json:"cpf"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (c *CustomerDTO) FromEntity(customer entity.Customer) {
	c.CPF = customer.CPF.Value
	c.Name = customer.Name
	c.Email = customer.Email
}

func (c *CustomerDTO) ToEntity() (*entity.Customer, error) {
	cpf, err := valueobject.NewCPF(c.CPF)
	if err != nil {
		return nil, err
	}

	customer := entity.Customer{
		CPF:   *cpf,
		Name:  c.Name,
		Email: c.Email,
	}

	return &customer, nil
}
