package entity

import (
	"errors"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/valueobject"
	"github.com/caiojorge/fiap-challenge-ddd/internal/shared/validator"
)

// Customer representa a entidade Cliente
type Customer struct {
	CPF   valueobject.CPF
	Name  string
	Email string
}

// NewCustomerWithCPFOnly identifica um cliente pelo CPF; na verdade, cria pelo CPF
func NewCustomerWithCPFOnly(cpf *valueobject.CPF) (*Customer, error) {
	if cpf == nil {
		return nil, errors.New("CPF is required")

	}

	return &Customer{
		CPF: *cpf,
	}, nil
}

func NewCustomer(cpf valueobject.CPF, name, email string) (*Customer, error) {

	customer := &Customer{
		CPF:   cpf,
		Name:  name,
		Email: email,
	}

	// Validate customer
	err := customer.Validate()
	if err != nil {
		return nil, err
	}

	return customer, nil
}

// RegisterCustomer caso o cliente queira se registrar, informando os atributos um a um
func (c *Customer) RegisterCustomer(name, email string) error {
	c.Name = name
	c.Email = email

	// Validate customer
	err := c.Validate()
	if err != nil {
		return err
	}

	return nil
}

// NewCustomer caso o cliente queira se registrar, informando todo os atributos ao mesmo tempo

func (c *Customer) GetCPF() *valueobject.CPF {
	return &c.CPF
}

func (c *Customer) GetName() string {
	return c.Name
}

func (c *Customer) GetEmail() string {
	return c.Email
}

// Validate valida todos os campos do cliente
func (c *Customer) Validate() error {

	if c.CPF.GetValue() == "" {
		return errors.New("CPF is required")
	}

	if c.Name == "" {
		return errors.New("name is required")
	}

	if c.Email == "" || !isValidateEmail(c.Email) {
		return errors.New("e-mail is required")
	}

	return nil
}

// isValidateEmail valida o formato do string enviado no padrão email
func isValidateEmail(email string) bool {
	v := validator.EmailValidator{}
	return v.IsValid(email)

}
