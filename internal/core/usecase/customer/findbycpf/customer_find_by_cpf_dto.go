package usecase

import (
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/valueobject"
)

// DTO
type CustomerFindByCpfInputDTO struct {
	CPF   string `json:"cpf"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (dto *CustomerFindByCpfInputDTO) ToEntity() *entity.Customer {
	return &entity.Customer{
		CPF: valueobject.CPF{
			Value: dto.CPF,
		},
		Name:  dto.Name,
		Email: dto.Email,
	}
}

type CustomerFindByCpfOutputDTO struct {
	CPF   string `json:"cpf"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (dto *CustomerFindByCpfOutputDTO) FromEntity(customer entity.Customer) {
	dto.CPF = customer.CPF.Value
	dto.Name = customer.Name
	dto.Email = customer.Email
}

// ...
