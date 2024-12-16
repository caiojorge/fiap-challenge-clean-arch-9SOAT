package entity

import (
	"errors"
	"time"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/valueobject"
	"github.com/caiojorge/fiap-challenge-ddd/internal/shared"
	"github.com/caiojorge/fiap-challenge-ddd/internal/shared/formatter"
	"github.com/caiojorge/fiap-challenge-ddd/internal/shared/validator"
)

type Order struct {
	ID          string
	Items       []*OrderItem
	Total       float64
	Status      string
	CustomerCPF string
	CreatedAt   time.Time
}

// OrderInit cria um novo pedido. TODO usada apenas no order_test.
func OrderInit(customerCPF string) *Order {

	order := Order{
		ID:          shared.NewIDGenerator(),
		CustomerCPF: customerCPF,
		Status:      valueobject.OrderStatusConfirmed,
	}

	return &order
}

// NewOrder cria um novo pedido. TODO Não esta sendo usada.
func NewOrder(cpf string, items []*OrderItem) (*Order, error) {

	order := Order{
		ID:          shared.NewIDGenerator(),
		CustomerCPF: cpf,
		Items:       items,
		Status:      valueobject.OrderStatusConfirmed,
	}

	if len(order.Items) > 0 {
		order.CalculateTotal()
	}

	err := order.Validate()
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (o *Order) GetOrderItemByProductID(productID string) *OrderItem {

	for _, item := range o.Items {
		if item.ProductID == productID {
			return item
		}
	}

	return nil
}

// Confirm confirma o pedido. Tem muita lógica de negócio aqui.
// Toda preparação necessária, validação de cpf, cálculo do total e validação dos itens.
// As regras aplicadas impactam apenas os dados da ordem / item.
func (o *Order) Confirm() error {

	o.ID = shared.NewIDGenerator()
	o.Status = valueobject.OrderStatusConfirmed

	for _, item := range o.Items {
		item.Confirm()
	}

	// Calcula o total do pedido se o item for confirmado
	o.CalculateTotal()

	// Valida o pedido
	err := o.Validate()
	if err != nil {
		return errors.New("failed to validate order")
	}

	return nil
}

func (o *Order) IsCustomerInformed() bool {
	return o.CustomerCPF != ""
}

func (o *Order) RemoveMaksFromCPF() {
	if o.CustomerCPF != "" {
		o.CustomerCPF = formatter.RemoveMaskFromCPF(o.CustomerCPF)
	}
}

func (o *Order) GetID() string {
	return o.ID
}

func (o *Order) Validate() error {

	if o.CustomerCPF != "" && len(o.CustomerCPF) == 11 {
		cpfValidator := validator.CPFValidator{}

		err := cpfValidator.Validate(o.CustomerCPF)
		if err != nil {
			return err
		}
	}

	if len(o.Items) == 0 {
		return errors.New("invalid order items")
	}

	return nil
}

func (o *Order) AddItem(item *OrderItem) {
	o.Items = append(o.Items, item)
}

func (o *Order) RemoveItem(item *OrderItem) {
	for i, v := range o.Items {
		if v == item {
			o.Items = append(o.Items[:i], o.Items[i+1:]...)
		}
	}
}

func (o *Order) CalculateTotal() {

	for _, item := range o.Items {
		if item.Status == valueobject.OrderItemStatusConfirmed {
			o.Total += (item.Price * float64(item.Quantity))
		}
	}
}

func (o *Order) IsPaid() bool {
	return o.Status == valueobject.OrderStatusPaid
}

func (o *Order) ConfirmPayment() {
	o.Status = valueobject.OrderItemStatusConfirmed
}

func (o *Order) NotConfirmPayment() {
	o.Status = valueobject.OrderStatusNotConfirmed
}

func (o *Order) Prepare() {
	o.Status = valueobject.OrderStatusPreparing
}

func (o *Order) Deliver() {
	o.Status = valueobject.OrderStatusDelivered
}

func (o *Order) Cancel() {
	o.Status = valueobject.OrderStatusCanceled
}
