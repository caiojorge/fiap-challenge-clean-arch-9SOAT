package portsservice

import (
	"context"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
)

type GatewayTransactionService interface {
	CreateTransaction(ctx context.Context, checkout *entity.Checkout, order *entity.Order, productList []*entity.Product, notificationURL string, sponsorID int) (*entity.Payment, error)
	CancelTransaction(ctx context.Context, id string) error
}
