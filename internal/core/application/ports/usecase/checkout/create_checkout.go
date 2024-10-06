package portsusecase

import (
	"context"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

// CreateCheckoutUseCase is the interface that wraps the CreateCheckout method.
type CreateCheckoutUseCase interface {
	CreateCheckout(ctx context.Context, checkout *entity.Checkout) (*string, error)
}
