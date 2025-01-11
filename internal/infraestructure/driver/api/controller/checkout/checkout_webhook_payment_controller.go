package controller

import (
	"context"

	portsusecase "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/checkout/create"
)

type WebhookCheckoutController struct {
	ctx     context.Context
	usecase portsusecase.ICreateCheckoutUseCase
}
