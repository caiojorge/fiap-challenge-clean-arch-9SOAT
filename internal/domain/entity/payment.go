package entity

import (
	"errors"

	"github.com/caiojorge/fiap-challenge-ddd/internal/shared"
)

type Item struct {
	SKU         string  `json:"sku_number"`
	Category    string  `json:"category"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	UnitPrice   float64 `json:"unit_price"`
	Quantity    int     `json:"quantity"`
	UnitMeasure string  `json:"unit_measure"`
	TotalAmount float64 `json:"total_amount"`
}

// Estrutura para a requisição de criação de pedido
// Payment é um agregado de pagamento.
type Payment struct {
	ID                string  `json:"id"`
	CheckoutID        string  `json:"checkout_id"`        // ID do checkout
	ExternalReference string  `json:"external_reference"` // ID da ordem
	Title             string  `json:"title"`
	Description       string  `json:"description"`
	NotificationURL   string  `json:"notification_url"`
	TotalAmount       float64 `json:"total_amount"`
	Items             []Item  `json:"items"`
	SponsorID         int     `json:"sponsor_id"`
	CashOutAmount     float64 `json:"cash_out_amount"`
}

func NewPayment(checkout Checkout, order Order, productList []*Product, notificationURL string, sponsorID int) (*Payment, error) {

	var items []Item
	orderItem := order.GetOrderItem()
	UnitMeasure := "un"

	for _, product := range productList {
		items = append(items, Item{
			SKU:         product.ID,
			Category:    product.Category,
			Title:       product.Name,
			Description: product.Description,
			UnitPrice:   product.Price,
			Quantity:    orderItem.Quantity,
			UnitMeasure: UnitMeasure,
			TotalAmount: product.Price * float64(orderItem.Quantity),
		})
	}

	payment := &Payment{
		ID:                shared.NewIDGenerator(),
		CheckoutID:        checkout.ID,
		ExternalReference: order.ID,
		Title:             "Pedido de compra",
		Description:       "Pedido de compra",
		NotificationURL:   notificationURL,
		TotalAmount:       order.Total,
		Items:             items,
		SponsorID:         sponsorID,
		CashOutAmount:     order.Total,
	}

	err := payment.Validate()
	if err != nil {
		return nil, err
	}

	return payment, nil

}

func (p *Payment) Validate() error {

	if p.ExternalReference == "" {
		return errors.New("external reference is required")
	}

	if p.Title == "" {
		return errors.New("title is required")
	}

	if p.Description == "" {
		return errors.New("description is required")
	}

	if p.NotificationURL == "" {
		return errors.New("notification url is required")
	}

	if p.TotalAmount == 0 {
		return errors.New("total amount is required")
	}

	if len(p.Items) == 0 {
		return errors.New("items is required")
	}

	if p.SponsorID == 0 {
		return errors.New("sponsor id is required")
	}

	if p.CashOutAmount == 0 {
		return errors.New("cash out amount is required")
	}

	return nil
}
