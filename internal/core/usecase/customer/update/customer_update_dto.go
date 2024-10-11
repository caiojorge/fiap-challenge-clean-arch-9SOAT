package usecase

import (
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/valueobject"
)

// DTO
type CustomerUpdateInputDTO struct {
	CPF   string `json:"cpf"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (dto *CustomerUpdateInputDTO) ToEntity() *entity.Customer {
	return &entity.Customer{
		CPF: valueobject.CPF{
			Value: dto.CPF,
		},
		Name:  dto.Name,
		Email: dto.Email,
	}
}

type CustomerUpdateOutputDTO struct {
	CPF   string `json:"cpf"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (dto *CustomerUpdateOutputDTO) FromEntity(customer entity.Customer) {
	dto.CPF = customer.CPF.Value
	dto.Name = customer.Name
	dto.Email = customer.Email
}

// ...
