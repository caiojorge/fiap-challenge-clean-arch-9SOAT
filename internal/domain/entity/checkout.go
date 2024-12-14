package entity

import (
	"errors"
	"time"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/valueobject"
	"github.com/caiojorge/fiap-challenge-ddd/internal/shared"
)

type Checkout struct {
	ID         string
	OrderID    string
	Gateway    valueobject.Gateway // TODO pensar em um value object para Gateway
	Total      float64
	CheckedOut time.Time
}

func NewCheckout(orderID string, gatewayName string, gatewayToken string, total float64) (*Checkout, error) {

	return &Checkout{
		ID:      shared.NewIDGenerator(),
		OrderID: orderID,
		Gateway: valueobject.NewGateway(gatewayName, gatewayToken),
		Total:   total,
	}, nil
}

func (c *Checkout) ConfirmTransaction(transactionID string, total float64) error {
	location, err := shared.GetBRLocationDefault()
	if err != nil {
		return err
	}

	c.ID = shared.NewIDGenerator()
	c.CheckedOut = time.Now().In(location)
	c.Gateway.GatewayTransactionID = transactionID
	c.Total = total

	err = c.Validate()
	if err != nil {
		return err
	}

	return nil
}

func (c *Checkout) Validate() error {
	if c.OrderID == "" {
		return errors.New("orderID is required")
	}

	// para essa versão, não é necessário validar dados do gateway
	// if c.Gateway.GatewayName == "" {
	// 	return errors.New("gateway is required")
	// }

	// if c.Gateway.GatewayToken == "" {
	// 	return errors.New("gateway token is required")
	// }

	return nil
}
