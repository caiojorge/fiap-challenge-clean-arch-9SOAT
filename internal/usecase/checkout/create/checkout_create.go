package usecase

import (
	"context"
	"errors"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	portsservice "github.com/caiojorge/fiap-challenge-ddd/internal/domain/gateway"
	repository "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository"
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
// Requisito: Checkout Pedido que deverá receber os produtos solicitados e retornar à identificação do pedido.
// o checkout recebe alguns parametros, sendo um deles, o id da ordem. Esse id, é usado para retornar a ordem com seus items, que por sua vez, tem os produtos.
// É dessa forma que penso em atender o requisito indicado acima
func (cr *CheckoutCreateUseCase) CreateCheckout(ctx context.Context, checkoutDTO *CheckoutInputDTO) (*CheckoutOutputDTO, error) {

	err := cr.validateDuplicatedCheckout(ctx, checkoutDTO)
	if err != nil {
		return nil, err
	}

	// Order- o pedido deve existir e não pode estar pago
	order, err := cr.validateAndReturnOrder(ctx, checkoutDTO)
	if err != nil {
		return nil, err
	}

	// via orderitems podemos pegar os produtos do pedido
	productList, err := cr.getProductList(ctx, order)
	if err != nil {
		return nil, errors.New("no products to send to checkout")
	}

	// converte dto para entidade
	checkout := checkoutDTO.ToEntity()

	// integração com o gateway de pagamento e confirmação do pagamento na ordem
	payment, err := cr.handlePayment(ctx, checkout, order, productList, checkoutDTO.NotificationURL, checkoutDTO.SponsorID, checkoutDTO.DiscontCoupon)
	if err != nil {
		return nil, err
	}

	// confirma a transação e grava o checkout
	err = cr.handleCheckout(ctx, checkout, order.Total, payment.ID)
	if err != nil {
		return nil, err
	}

	// Kitchen - cria um item na cozinha para cada produto do pedido
	// err = cr.handleKitchen(ctx, order, productList, *payment)
	// if err != nil {
	// 	return nil, err
	// }

	output := &CheckoutOutputDTO{
		ID:                   checkout.ID,
		GatewayTransactionID: payment.ID,
		OrderID:              order.ID,
	}

	return output, nil
}

// Checkout - o cliente não pode fazer checkout duas vezes
// se existir um checkout para o pedido, não pode fazer outro
func (cr *CheckoutCreateUseCase) validateDuplicatedCheckout(ctx context.Context, checkout *CheckoutInputDTO) error {
	ch, _ := cr.checkoutRepository.FindbyOrderID(ctx, checkout.OrderID)
	if ch != nil {
		return errors.New("you can not checkout twice")
	}

	return nil
}

// Order- o pedido deve existir e não pode estar pago
func (cr *CheckoutCreateUseCase) validateAndReturnOrder(ctx context.Context, checkout *CheckoutInputDTO) (*entity.Order, error) {
	// Order- o pedido deve existir e não pode estar pago
	order, err := cr.orderRepository.Find(ctx, checkout.OrderID)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, errors.New("order not found")
	}
	if order.IsPaymentApproved() {
		return nil, errors.New("order already paid")
	}

	return order, nil
}

// getProductList retorna a lista de produtos do pedido
// TODO: esse método pode ser movido para um repositório de produtos
func (cr *CheckoutCreateUseCase) getProductList(ctx context.Context, order *entity.Order) ([]*entity.Product, error) {
	var products []*entity.Product

	for _, item := range order.Items {
		product, err := cr.productRepository.Find(ctx, item.ProductID)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

// handlePayment integração com o gateway de pagamento e confirmação do pagamento na ordem
func (cr *CheckoutCreateUseCase) handlePayment(ctx context.Context, checkout *entity.Checkout, order *entity.Order, productList []*entity.Product, notificationURL string, sponsorID int, cupon float64) (*entity.Payment, error) {

	payment, err := cr.gatewayService.ConfirmPayment(ctx, checkout, order, productList, notificationURL, sponsorID)
	if err != nil {
		order.InformPaymentNotApproval()
		_ = cr.orderRepository.Update(ctx, order)
		return nil, err
	}

	if payment == nil {
		order.InformPaymentNotApproval()
		_ = cr.orderRepository.Update(ctx, order)
		return nil, errors.New("failed to create transaction on gateway")
	}

	// //order.ApprovePayment()
	// err = cr.orderRepository.Update(ctx, order)
	// if err != nil {
	// 	cr.gatewayService.CancelPayment(ctx, payment.ID) // não tem rollback no gateway
	// 	return nil, err
	// }

	return payment, nil
}

// handleCheckout confirma a transação e grava o checkout
func (cr *CheckoutCreateUseCase) handleCheckout(ctx context.Context, checkout *entity.Checkout, orderTotal float64, paymentID string) error {

	// confirma o pagamento de forma fake na própria entidade (muda status na entidade)
	// se houver algum cupom de desconto, ele será aplicado na ordem no handlePayment
	err := checkout.ConfirmTransaction(paymentID, orderTotal)
	if err != nil {
		cr.gatewayService.CancelPayment(ctx, paymentID)
		return err
	}

	// grava o checkout
	err = cr.checkoutRepository.Create(ctx, checkout)
	if err != nil {
		cr.gatewayService.CancelPayment(ctx, paymentID)
		return err
	}

	return nil
}

func (cr *CheckoutCreateUseCase) handleKitchen(ctx context.Context, order *entity.Order, productList []*entity.Product, payment entity.Payment) error {

	// Kitchen - cria um item na cozinha para cada produto do pedido
	for _, item := range productList {
		kt := entity.NewKitchen(order.ID, item.ID, item.Name, item.Category)
		err := cr.kitchenRepository.Create(ctx, kt)
		if err != nil {
			return err
		}
	}

	return nil
}
