package usecase

import (
	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/valueobject"
)

type CheckoutInputDTO struct {
	OrderID      string `json:"order_id"`      // o ID do pedido que será pago
	GatewayName  string `json:"gateway_name"`  // nome do gateway (nesse caso, mercado livre fake)
	GatewayToken string `json:"gateway_token"` // id ou token para uso no gateway
}

func (dto *CheckoutInputDTO) ToEntity() *entity.Checkout {
	return &entity.Checkout{
		//ID:                   dto.ID,
		OrderID: dto.OrderID,
		Gateway: valueobject.NewGateway(dto.GatewayName, dto.GatewayToken),
	}
}

type CheckoutOutputDTO struct {
	ID                   string `json:"id"`                     // ID do checkout
	GatewayTransactionID string `json:"gateway_transaction_id"` // ID de transação gerado pelo gateway
	OrderID              string `json:"order_id"`               // ID do pedido
}

func (dto *CheckoutOutputDTO) FromEntity(entity entity.Checkout) {
	dto.ID = entity.ID
	dto.GatewayTransactionID = entity.Gateway.GatewayTransactionID
	dto.OrderID = entity.OrderID
}
