package usecase

import (
	"time"

	shared "github.com/caiojorge/fiap-challenge-ddd/internal/core/usecase/order/shared"
)

type OrderCreateInputDTO struct {
	Items       []*OrderItemCreateInputDTO `json:"items"`
	CustomerCPF string                     `json:"cpf"`
}

type OrderItemCreateInputDTO struct {
	ProductID string `json:"productid"`
	Quantity  int    `json:"quantity"`
}

type OrderCreateOutputDTO struct {
	ID          string                 `json:"id"`
	Items       []*shared.OrderItemDTO `json:"items"`
	Total       float64                `json:"total"`
	Status      string                 `json:"status"`
	CustomerCPF string                 `json:"customercpf"`
	CreatedAt   time.Time              `json:"created_at"`
}
