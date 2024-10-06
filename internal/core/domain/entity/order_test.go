package entity

import (
	"testing"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/valueobject"
	"github.com/caiojorge/fiap-challenge-ddd/internal/shared"
	"github.com/caiojorge/fiap-challenge-ddd/internal/shared/validator"
	"github.com/stretchr/testify/assert"
)

// Test Order

func TestOrder(t *testing.T) {
	// Customer
	cpf, err := valueobject.NewCPF("19528476562")
	assert.Nil(t, err)
	customer, err := NewCustomer(*cpf, "Caio", "email@email.com")
	assert.Nil(t, err)
	assert.Equal(t, "Caio", customer.Name)

	// Product
	product, err := NewProduct("Lanche xpto", "Pão, carne e queijo", "lanche", 10)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	product2, err := NewProduct("Coca Cola", "Água com gás e xarope de coca", "refrigerante", 10)
	assert.Nil(t, err)
	assert.NotNil(t, product2)

	// Order
	order := OrderInit(customer.GetCPF().Value)
	order.ID = "1"
	assert.Equal(t, "1", order.ID)
	assert.Equal(t, "19528476562", order.CustomerCPF)

	orderItem, err := NewOrderItem(product.GetID(), 1, product.Price)
	assert.Nil(t, err)
	assert.NotNil(t, orderItem)
	orderItem2, err := NewOrderItem(product2.GetID(), 1, product2.Price)
	assert.Nil(t, err)
	assert.NotNil(t, orderItem2)

	order.AddItem(orderItem)
	order.AddItem(orderItem2)

	assert.Equal(t, 2, len(order.Items))

	order.CalculateTotal()
	assert.Equal(t, 20.00, order.Total)

	// validando o cpf q seria informado pelo cliente.
	v := validator.CPFValidator{}
	assert.True(t, v.IsValid(order.CustomerCPF))
	err = v.Validate(order.CustomerCPF)
	assert.Nil(t, err)

	err = order.Validate()
	assert.Nil(t, err)

}

func TestOrderWithNoItens(t *testing.T) {
	// Customer
	cpf, err := valueobject.NewCPF("19528476562")
	assert.Nil(t, err)
	customer, err := NewCustomer(*cpf, "Caio", "email@email.com")
	assert.Nil(t, err)
	assert.Equal(t, "Caio", customer.Name)

	// Product
	product, err := NewProduct("Lanche xpto", "Pão, carne e queijo", "lanche", 10)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	product2, err := NewProduct("Coca Cola", "Água com gás e xarope de coca", "refrigerante", 10)
	assert.Nil(t, err)
	assert.NotNil(t, product2)

	// Order
	order := OrderInit(customer.GetCPF().Value)
	order.ID = "1"
	assert.Equal(t, "1", order.ID)
	assert.Equal(t, "19528476562", order.CustomerCPF)

	// validando o cpf q seria informado pelo cliente.
	v := validator.CPFValidator{}
	assert.True(t, v.IsValid(order.CustomerCPF))
	err = v.Validate(order.CustomerCPF)
	assert.Nil(t, err)

	err = order.Validate()
	assert.NotNil(t, err)

}

func TestOrderWithNoCustomer(t *testing.T) {

	// Product
	product, err := NewProduct("Lanche xpto", "Pão, carne e queijo", "lanche", 10)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	product2, err := NewProduct("Coca Cola", "Água com gás e xarope de coca", "refrigerante", 10)
	assert.Nil(t, err)
	assert.NotNil(t, product2)

	// Order
	order := OrderInit("")
	order.ID = "1"
	assert.Equal(t, "1", order.ID)

	orderItem, err := NewOrderItem(product.GetID(), 1, product.Price)
	assert.Nil(t, err)
	assert.NotNil(t, orderItem)
	orderItem2, err := NewOrderItem(product2.GetID(), 1, product2.Price)
	assert.Nil(t, err)
	assert.NotNil(t, orderItem2)

	order.AddItem(orderItem)
	order.AddItem(orderItem2)

	assert.Equal(t, 2, len(order.Items))

	order.CalculateTotal()
	assert.Equal(t, 20.00, order.Total)

	// cpf não informado; tem q dar erro.
	v := validator.CPFValidator{}
	err = v.Validate(order.CustomerCPF)
	assert.NotNil(t, err)

	// não tem cliente mas tem itens, tem q dar certo
	err = order.Validate()
	assert.Nil(t, err)

}

func TestOrderWithNoRegistration(t *testing.T) {
	// Customer
	cpf, err := valueobject.NewCPF("19528476562")
	assert.Nil(t, err)
	customer, err := NewCustomerWithCPFOnly(cpf)
	assert.Nil(t, err)

	// Product
	product, err := NewProduct("Lanche xpto", "Pão, carne e queijo", "lanche", 10)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	product2, err := NewProduct("Coca Cola", "Água com gás e xarope de coca", "refrigerante", 10)
	assert.Nil(t, err)
	assert.NotNil(t, product2)

	// Order
	order := OrderInit(customer.GetCPF().Value)
	order.ID = "1"
	assert.Equal(t, "1", order.ID)
	assert.Equal(t, "19528476562", order.CustomerCPF)

	orderItem, err := NewOrderItem(product.GetID(), 1, product.Price)
	assert.Nil(t, err)
	assert.NotNil(t, orderItem)
	orderItem2, err := NewOrderItem(product2.GetID(), 1, product2.Price)
	assert.Nil(t, err)
	assert.NotNil(t, orderItem2)

	// cancelando o item2, então o valor do pedido deveria ser 10
	orderItem2.Cancel()

	order.AddItem(orderItem)
	order.AddItem(orderItem2)

	assert.Equal(t, 2, len(order.Items))

	order.CalculateTotal()
	assert.Equal(t, 10.00, order.Total)

	assert.Equal(t, valueobject.OrderStatusConfirmed, order.Status)

	// validando o cpf q seria informado pelo cliente.
	v := validator.CPFValidator{}
	assert.True(t, v.IsValid(order.CustomerCPF))
	err = v.Validate(order.CustomerCPF)
	assert.Nil(t, err)

	err = order.Validate()
	assert.Nil(t, err)

}

func TestConfirmedOrder(t *testing.T) {
	cpf := "75419654059"

	lanche := Product{
		ID:          shared.NewIDGenerator(),
		Name:        "Burger Kong",
		Description: "Pão, carne e queijo",
		Category:    "lanche",
		Price:       50.0,
	}

	refri := Product{
		ID:          shared.NewIDGenerator(),
		Name:        "Pepsicola",
		Description: "Peptococa",
		Category:    "refrigerante",
		Price:       10.0,
	}

	item := OrderItem{
		ProductID: lanche.ID,
		Quantity:  1,
		Price:     lanche.Price,
	}

	item2 := OrderItem{
		ProductID: refri.ID,
		Quantity:  1,
		Price:     lanche.Price,
	}

	order := Order{
		Items:       []*OrderItem{&item, &item2},
		CustomerCPF: cpf,
	}

	err := order.ConfirmOrder()
	assert.Nil(t, err)

}
