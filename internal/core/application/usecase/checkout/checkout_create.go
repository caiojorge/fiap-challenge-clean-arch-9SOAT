package usecase

import (
	"context"
	"errors"

	portsrepository "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/repository"
	portsservice "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/service"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
	productRepository "github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/repository"
)

// CheckoutCreateUseCase é a implementação da interface CheckoutCreateUseCase.
// Nesse momento, não iremos implementar a integração com o gateway de pagamento.
type CheckoutCreateUseCase struct {
	orderRepository    portsrepository.OrderRepository
	checkoutRepository portsrepository.CheckoutRepository
	gatewayService     portsservice.GatewayTransactionService
	kitchenRepository  portsrepository.KitchenRepository
	productRepository  productRepository.ProductRepository
}

func NewCheckoutCreate(orderRepository portsrepository.OrderRepository,
	checkoutRepository portsrepository.CheckoutRepository,
	gatewayService portsservice.GatewayTransactionService,
	kitchenRepository portsrepository.KitchenRepository,
	productRepository productRepository.ProductRepository) *CheckoutCreateUseCase {
	return &CheckoutCreateUseCase{
		orderRepository:    orderRepository,
		checkoutRepository: checkoutRepository,
		gatewayService:     gatewayService,
		productRepository:  productRepository,
		kitchenRepository:  kitchenRepository,
	}
}

// CreateCheckout registra o checkout de um pedido.
func (cr *CheckoutCreateUseCase) CreateCheckout(ctx context.Context, checkout *entity.Checkout) (*string, error) {

	// Checkout - o cliente não pode fazer checkout duas vezes
	ch, _ := cr.checkoutRepository.FindbyOrderID(ctx, checkout.ID)
	if ch != nil {
		return nil, errors.New("you can not checkout twice")
	}

	// Order
	order, err := cr.orderRepository.Find(ctx, checkout.OrderID)
	if err != nil {
		return nil, err
	}

	if order == nil {
		return nil, errors.New("order not found")
	}

	if order.IsPaid() {
		return nil, errors.New("order already paid")
	}

	// Gateway
	transactionID, err := cr.gatewayService.CreateTransaction(ctx, checkout)
	if err != nil {
		return nil, err
	}

	if transactionID == nil {
		return nil, errors.New("failed to create transaction on gateway")
	}

	// Checkout
	order.Pay()
	err = cr.orderRepository.Update(ctx, order)
	if err != nil {
		cr.gatewayService.CancelTransaction(ctx, *transactionID) // não tem rollback no gateway
		//rollbackOrder(ctx, order, cr)
		return nil, err
	}

	cupon := 0.0 // apenas um exemplo de como poderiamos aplicar um cupom de desconto no pagamento
	err = checkout.ConfirmTransaction(*transactionID, order.Total-cupon)
	if err != nil {
		cr.gatewayService.CancelTransaction(ctx, *transactionID)
		//rollbackOrder(ctx, order, cr)
		return nil, err
	}

	err = cr.checkoutRepository.Create(ctx, checkout)
	if err != nil {
		cr.gatewayService.CancelTransaction(ctx, *transactionID)
		//rollbackOrder(ctx, order, cr)
		return nil, err
	}

	// Kitchen
	for _, item := range order.Items {
		product, _ := cr.productRepository.Find(ctx, item.ProductID)
		kt := entity.NewKitchen(order.ID, item.ID, product.Name, product.Category)
		err = cr.kitchenRepository.Create(ctx, kt)
		if err != nil {
			cr.gatewayService.CancelTransaction(ctx, *transactionID)
			return nil, err
		}
	}

	return transactionID, nil
}

// func rollbackOrder(ctx context.Context, order *entity.Order, cr *CheckoutCreateUseCase) {
// 	order.Status = valueobject.OrderStatusConfirmed
// 	cr.orderRepository.Update(ctx, order)
// }
