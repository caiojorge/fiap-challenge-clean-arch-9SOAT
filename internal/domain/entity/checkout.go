package entity

import (
	"errors"
	"time"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/valueobject"
	sharedgenerator "github.com/caiojorge/fiap-challenge-ddd/internal/shared/generator"
)

type Checkout struct {
	ID        string
	OrderID   string
	Gateway   valueobject.Gateway // TODO pensar em um value object para Gateway
	Total     float64
	CreatedAt time.Time
}

func NewCheckout(orderID string, gatewayName string, gatewayToken string, total float64) (*Checkout, error) {

	return &Checkout{
		ID:      sharedgenerator.NewIDGenerator(),
		OrderID: orderID,
		Gateway: valueobject.NewGateway(gatewayName, gatewayToken),
		Total:   total,
	}, nil
}

func (c *Checkout) ConfirmTransaction(transactionID string, total float64) error {

	c.ID = sharedgenerator.NewIDGenerator()
	c.Gateway.GatewayTransactionID = transactionID
	c.Total = total

	err := c.Validate()
	if err != nil {
		return err
	}

	return nil
}

func (c *Checkout) Validate() error {
	if c.OrderID == "" {
		return errors.New("orderID is required")
	}

	return nil
}
