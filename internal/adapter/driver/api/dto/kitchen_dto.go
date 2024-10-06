package dto

import "time"

type KitchenDTO struct {
	ID          string    `json:"id"`
	OrderID     string    `json:"order_id"`
	ItemID      string    `json:"item_order_id"`
	ProductName string    `json:"product_name"`
	Responsible string    `json:"responsible"`
	CreatedAt   time.Time `json:"created_at"`
}
