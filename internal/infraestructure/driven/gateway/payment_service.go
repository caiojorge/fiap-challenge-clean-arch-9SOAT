package service

import (
	"context"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
	"github.com/caiojorge/fiap-challenge-ddd/internal/shared"
)

// FakePaymentService provides methods for payment operations.
type FakePaymentService struct {
}

func NewFakePaymentService() *FakePaymentService {
	return &FakePaymentService{}
}

// CreateCheckout creates a new checkout. This method should be implemented by the payment gateway.
func (p *FakePaymentService) CreateTransaction(ctx context.Context, checkout *entity.Checkout) (*string, error) {
	transactionID := shared.NewIDGenerator() // Fake transaction ID
	return &transactionID, nil
}

// CancelTransaction cancels a transaction. This method should be implemented by the payment gateway.
func (p *FakePaymentService) CancelTransaction(ctx context.Context, id string) error {
	return nil
}
