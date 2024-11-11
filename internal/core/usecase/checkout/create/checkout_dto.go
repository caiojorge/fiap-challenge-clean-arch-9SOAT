package usecase

import (
	"time"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

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

type CheckoutInputDTO struct {
	OrderID     string `json:"order_id"`
	Gateway     string `json:"gateway"`
	GatewayID   string `json:"gateway_id"`
	CustomerCPF string `json:"customer_cpf"`
}

func (dto *CheckoutInputDTO) ToEntity() *entity.Checkout {
	return &entity.Checkout{
		//ID:                   dto.ID,
		OrderID:   dto.OrderID,
		Gateway:   dto.Gateway,
		GatewayID: dto.GatewayID,
		//GatewayTransactionID: dto.GatewayTransactionID,
		CustomerCPF: dto.CustomerCPF,
		//Total:                dto.Total,
		//CreatedAt:            dto.CreatedAt,
	}
}

type CheckoutOutputDTO struct {
	ID                   string `json:"id"`
	GatewayTransactionID string `json:"gateway_transaction_id"`
}

func (dto *CheckoutOutputDTO) FromEntity(product entity.Checkout) {
	dto.ID = product.ID
	dto.GatewayTransactionID = product.GatewayTransactionID
}
