package entity

import (
	"time"

	sharedgenerator "github.com/caiojorge/fiap-challenge-ddd/internal/shared/generator"
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
	var responsible string

	if category == "bebida" || category == "refrigerante" {
		responsible = "bar"
	} else {
		responsible = "kitchen"
	}

	return &Kitchen{
		ID:          sharedgenerator.NewIDGenerator(),
		OrderID:     orderID,
		ItemID:      itemOrderID,
		ProductName: productName,
		Responsible: responsible,
	}
}
