package usecase

import (
	"context"
	"errors"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
	portsservice "github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/gateway"
	repository "github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/repository"
)

// CheckoutCreateUseCase é a implementação da interface CheckoutCreateUseCase.
// Nesse momento, não iremos implementar a integração com o gateway de pagamento.
type CheckoutCreateUseCase struct {
	orderRepository    repository.OrderRepository
	checkoutRepository repository.CheckoutRepository
	gatewayService     portsservice.GatewayTransactionService
	kitchenRepository  repository.KitchenRepository
	productRepository  repository.ProductRepository
}

func NewCheckoutCreate(orderRepository repository.OrderRepository,
	checkoutRepository repository.CheckoutRepository,
	gatewayService portsservice.GatewayTransactionService,
	kitchenRepository repository.KitchenRepository,
	productRepository repository.ProductRepository) *CheckoutCreateUseCase {
	return &CheckoutCreateUseCase{
		orderRepository:    orderRepository,
		checkoutRepository: checkoutRepository,
		gatewayService:     gatewayService,
		productRepository:  productRepository,
		kitchenRepository:  kitchenRepository,
	}
}

// CreateCheckout registra o checkout de um pedido.
func (cr *CheckoutCreateUseCase) CreateCheckout(ctx context.Context, checkout *CheckoutInputDTO) (*CheckoutOutputDTO, error) {

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

	checkoutEntity := checkout.ToEntity()

	// Gateway
	transactionID, err := cr.gatewayService.CreateTransaction(ctx, checkoutEntity)
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
	err = checkoutEntity.ConfirmTransaction(*transactionID, order.Total-cupon)
	if err != nil {
		cr.gatewayService.CancelTransaction(ctx, *transactionID)
		//rollbackOrder(ctx, order, cr)
		return nil, err
	}

	err = cr.checkoutRepository.Create(ctx, checkoutEntity)
	if err != nil {
		cr.gatewayService.CancelTransaction(ctx, *transactionID)
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

	output := &CheckoutOutputDTO{
		ID:                   checkoutEntity.ID,
		GatewayTransactionID: *transactionID,
	}

	return output, nil
}
