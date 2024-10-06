package portsservice

import (
	"context"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

type GatewayTransactionService interface {
	CreateTransaction(ctx context.Context, checkout *entity.Checkout) (*string, error)
	CancelTransaction(ctx context.Context, id string) error
}
