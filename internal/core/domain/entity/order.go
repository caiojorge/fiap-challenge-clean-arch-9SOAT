package entity

import (
	"errors"
	"log"
	"time"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/valueobject"
	"github.com/caiojorge/fiap-challenge-ddd/internal/shared"
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

func OrderInit(customerCPF string) *Order {

	location, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		log.Default().Println("failed to load location")
	}

	order := Order{
		ID:          shared.NewIDGenerator(),
		CustomerCPF: customerCPF,
		Status:      valueobject.OrderStatusConfirmed,
		CreatedAt:   time.Now().In(location),
	}

	return &order
}

func NewOrder(cpf string, items []*OrderItem) (*Order, error) {
	location, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		return nil, err
	}

	order := Order{
		ID:          shared.NewIDGenerator(),
		CustomerCPF: cpf,
		Items:       items,
		Status:      valueobject.OrderStatusConfirmed,
		CreatedAt:   time.Now().In(location),
	}

	err = order.Validate()
	if err != nil {
		return nil, err
	}

	return &order, nil
}

// ConfirmOrder confirma o pedido. Tem muita lógica de negócio aqui.
// Toda preparação necessária, validação de cpf, cálculo do total e validação dos itens.
// As regras aplicadas impactam apenas os dados da ordem / item.
func (o *Order) ConfirmOrder() error {

	location, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		return errors.New("failed to load location")
	}

	o.ID = shared.NewIDGenerator()
	o.Status = valueobject.OrderStatusConfirmed
	o.CreatedAt = time.Now().In(location)

	for _, item := range o.Items {
		item.ConfirmItem()
	}

	// Calcula o total do pedido se o item for confirmado
	o.CalculateTotal()

	// Valida o pedido
	err = o.Validate()
	if err != nil {
		return errors.New("failed to validate order")
	}

	return nil
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
			o.Total += item.Price
		}
	}
}

func (o *Order) IsPaid() bool {
	return o.Status == valueobject.OrderStatusPaid
}

func (o *Order) Pay() {
	o.Status = valueobject.OrderStatusPaid
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
