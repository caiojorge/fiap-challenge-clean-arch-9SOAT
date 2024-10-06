package entity

import (
	"time"

	"github.com/caiojorge/fiap-challenge-ddd/internal/shared"
)

type Kitchen struct {
	ID          string
	OrderID     string
	ItemID      string
	ProductName string
	Responsible string
	CreatedAt   time.Time
}

func NewKitchen(orderID, itemOrderID, productName, category string) *Kitchen {
	location, _ := time.LoadLocation("America/Sao_Paulo")
	var responsible string

	if category == "bebida" || category == "refrigerante" {
		responsible = "bar"
	} else {
		responsible = "kitchen"
	}

	return &Kitchen{
		ID:          shared.NewIDGenerator(),
		OrderID:     orderID,
		ItemID:      itemOrderID,
		ProductName: productName,
		Responsible: responsible,
		CreatedAt:   time.Now().In(location),
	}
}
