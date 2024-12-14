package usecase

import (
	"context"
	"testing"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	mocks "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository/mocks"
	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/valueobject"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateCheckout(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCheckoutRepository := mocks.NewMockCheckoutRepository(ctrl)
	assert.NotNil(t, mockCheckoutRepository)
	mockOrderRepository := mocks.NewMockOrderRepository(ctrl)
	assert.NotNil(t, mockOrderRepository)
	mockProductRepository := mocks.NewMockProductRepository(ctrl)
	assert.NotNil(t, mockProductRepository)
	mockGatewayService := NewMLFakePaymentService()
	assert.NotNil(t, mockGatewayService)
	mockKitchenRepository := mocks.NewMockKitchenRepository(ctrl)
	assert.NotNil(t, mockKitchenRepository)

	useCase := NewCheckoutCreate(
		mockOrderRepository,
		mockCheckoutRepository,
		mockGatewayService,
		mockKitchenRepository,
		mockProductRepository,
	)
	assert.NotNil(t, useCase)

	ctx := context.Background()
	assert.NotNil(t, ctx)

	// Define input DTO
	checkoutInput := &CheckoutInputDTO{
		OrderID: "order123",
	}

	// Define entities for the mocks to return
	order := &entity.Order{
		ID:     "order123",
		Status: valueobject.OrderItemStatusConfirmed,
		Items: []*entity.OrderItem{
			{ProductID: "prod123", Quantity: 1, Status: valueobject.OrderItemStatusConfirmed, Price: 100.0},
		},
	}

	order.CalculateTotal()

	product := &entity.Product{
		ID:    "prod123",
		Name:  "Test Product",
		Price: 100.0,
	}

	payment := &entity.Payment{
		ID: "payment123",
	}

	// Set up mock expectations for a successful checkout
	mockCheckoutRepository.EXPECT().
		FindbyOrderID(ctx, "order123").
		Return(nil, nil) // No duplicate checkout found

	mockOrderRepository.EXPECT().
		Find(ctx, "order123").
		Return(order, nil) // Order found and not paid

	mockProductRepository.EXPECT().
		Find(ctx, "prod123").
		Return(product, nil) // Product found

	// mockGatewayService.EXPECT().
	// 	CreateTransaction(ctx, gomock.Any(), order, gomock.Any(), "http://localhost:8080/checkout", 1).
	// 	Return(payment, nil) // Payment successful

	mockCheckoutRepository.EXPECT().
		Create(ctx, gomock.Any()).
		Return(nil) // Checkout creation successful

	mockKitchenRepository.EXPECT().
		Create(ctx, gomock.Any()).
		Return(nil) // Kitchen entry creation successful

	// Execute the test
	result, err := useCase.CreateCheckout(ctx, checkoutInput)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, payment.ID, result.GatewayTransactionID)

}

type MLFakePaymentService struct {
}

func NewMLFakePaymentService() *MLFakePaymentService {
	return &MLFakePaymentService{}
}

// CreateCheckout creates a new checkout. This method should be implemented by the payment gateway.
func (p *MLFakePaymentService) CreateTransaction(ctx context.Context, checkout *entity.Checkout, order *entity.Order, productList []*entity.Product, notificationURL string, sponsorID int) (*entity.Payment, error) {
	payment, err := entity.NewPayment(*checkout, *order, productList, notificationURL, sponsorID)
	if err != nil {
		return nil, err
	}

	// TODO: connectar no server de pagamento
	// enviar dados de pagamento para o gateway
	// tratar a resposta do gateway

	return payment, nil
}

// CancelTransaction cancels a transaction. This method should be implemented by the payment gateway.
func (p *MLFakePaymentService) CancelTransaction(ctx context.Context, id string) error {
	return nil
}
