package usecase

import (
	"context"

	repository "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository"
	customerrors "github.com/caiojorge/fiap-challenge-ddd/internal/shared/error"
)

type CheckoutConfirmationInputDTO struct {
	OrderID string `json:"order_id" binding:"required"`
	Status  string `json:"status" binding:"required"`
}

type CheckoutConfirmationOutputDTO struct {
	CheckoutID string `json:"checkout_id" binding:"required"`
	OrderID    string `json:"order_id" binding:"required"`
	Status     string `json:"status" binding:"required"`
}

type ICheckoutConfirmationUseCase interface {
	ConfirmPayment(ctx context.Context, checkout *CheckoutConfirmationInputDTO) (*CheckoutConfirmationOutputDTO, error)
}

type CheckoutConfirmationUseCase struct {
	orderRepository    repository.OrderRepository
	checkoutRepository repository.CheckoutRepository
	tm                 repository.TransactionManager
}

func NewCheckoutConfirmation(orderRepository repository.OrderRepository,
	checkoutRepository repository.CheckoutRepository, tm repository.TransactionManager) *CheckoutConfirmationUseCase {
	return &CheckoutConfirmationUseCase{
		orderRepository:    orderRepository,
		checkoutRepository: checkoutRepository,
		tm:                 tm,
	}
}

func (cr *CheckoutConfirmationUseCase) ConfirmPayment(ctx context.Context, input *CheckoutConfirmationInputDTO) (*CheckoutConfirmationOutputDTO, error) {

	// 1. Verificar se a ordem existe
	order, err := cr.orderRepository.Find(ctx, input.OrderID)
	if err != nil {
		return nil, err
	}

	if order == nil {
		return nil, customerrors.ErrOrderNotFound
	}

	// 2. Verificar se o checkout existe, com base no id da ordem
	checkout, err := cr.checkoutRepository.FindbyOrderID(ctx, input.OrderID)
	if err != nil {
		return nil, err
	}

	if checkout == nil {
		return nil, customerrors.ErrCheckoutNotFound
	}

	// 3. Mudar status da ordem
	order.ConfirmPayment()

	// 4. Mudar status do checkout
	checkout.ConfirmPayment()

	err = cr.tm.RunInTransaction(ctx, func(ctx context.Context) error {
		// 5. Salvar novo status da ordem
		err := cr.orderRepository.Update(ctx, order)
		if err != nil {
			return err
		}

		// 6. Salvar novo status do checkout
		err = cr.checkoutRepository.Update(ctx, checkout)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	output := &CheckoutConfirmationOutputDTO{
		CheckoutID: checkout.ID,
		OrderID:    order.ID,
		Status:     order.Status,
	}

	return output, nil
}
