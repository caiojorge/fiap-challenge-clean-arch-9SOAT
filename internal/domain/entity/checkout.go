package entity

import (
	"errors"
	"time"

	"github.com/caiojorge/fiap-challenge-ddd/internal/shared"
	"github.com/caiojorge/fiap-challenge-ddd/internal/shared/validator"
)

type Checkout struct {
	ID                   string
	OrderID              string
	Gateway              string // TODO pensar em um value object para Gateway
	GatewayID            string
	GatewayTransactionID string
	CustomerCPF          string
	Total                float64
	CreatedAt            time.Time
}

func NewCheckout(orderID, gateway, gatewayID, customerCPF string, total float64) (*Checkout, error) {

	location, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		return nil, err
	}

	return &Checkout{
		ID:          shared.NewIDGenerator(),
		OrderID:     orderID,
		Gateway:     gateway,
		GatewayID:   gatewayID,
		CustomerCPF: customerCPF,
		Total:       total,
		CreatedAt:   time.Now().In(location),
	}, nil
}

func (c *Checkout) ConfirmTransaction(transactionID string, total float64) error {
	location, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		return err
	}

	c.ID = shared.NewIDGenerator()
	c.CreatedAt = time.Now().In(location)
	c.GatewayTransactionID = transactionID
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

	if c.Gateway == "" {
		return errors.New("gateway is required")
	}

	if c.GatewayID == "" {
		return errors.New("gatewayID is required")
	}

	if c.CustomerCPF == "" {
		return errors.New("customerCPF is required")
	} else {
		validator := validator.CPFValidator{}
		err := validator.Validate(c.CustomerCPF)
		if err != nil {
			return err
		}
	}

	// if c.Total == 0 {
	// 	return errors.New("total is required")
	// }

	return nil
}
