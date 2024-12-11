package usecase

import (
	"context"
)

// CreateCheckoutUseCase is the interface that wraps the CreateCheckout method.
type CreateCheckoutUseCase interface {
	CreateCheckout(ctx context.Context, checkout *CheckoutInputDTO) (*CheckoutOutputDTO, error)
}
