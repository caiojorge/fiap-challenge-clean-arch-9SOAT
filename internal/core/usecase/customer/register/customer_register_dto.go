package usecase

import (
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/valueobject"
)

// RegisterProductInputDTO para atender o use case de registro de produto
type CustomerRegisterInputDTO struct {
	CPF   string `json:"cpf"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (dto *CustomerRegisterInputDTO) ToEntity() *entity.Customer {
	return &entity.Customer{
		CPF: valueobject.CPF{
			Value: dto.CPF,
		},
		Name:  dto.Name,
		Email: dto.Email,
	}
}

type CustomerRegisterOutputDTO struct {
	CPF   string `json:"cpf"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (dto *CustomerRegisterOutputDTO) FromEntity(customer entity.Customer) {
	dto.CPF = customer.CPF.Value
	dto.Name = customer.Name
	dto.Email = customer.Email
}

// ...
