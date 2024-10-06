package dto

import "time"

type CreateCheckoutDTO struct {
	OrderID     string `json:"order_id"`
	Gateway     string `json:"gateway"`
	GatewayID   string `json:"gateway_id"`
	CustomerCPF string `json:"customer_cpf"`
}

type CheckoutDTO struct {
	ID                   string    `json:"id"`
	OrderID              string    `json:"order_id"`
	Gateway              string    `json:"gateway"`
	GatewayID            string    `json:"gateway_id"`
	GatewayTransactionID string    `json:"gateway_transaction_id"`
	CustomerCPF          string    `json:"customer_cpf"`
	Total                float64   `json:"total"`
	CreatedAt            time.Time `json:"created_at"`
}
