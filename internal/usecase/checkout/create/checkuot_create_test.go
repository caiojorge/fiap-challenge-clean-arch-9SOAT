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

	// Criar product
	product := &entity.Product{
		ID:          "1",
		Name:        "Product 1",
		Description: "Product 1 description",
		Price:       10.0,
		Category:    "Category 1",
	}

	// Criar customer
	customer := &entity.Customer{
		CPF: valueobject.CPF{
			Value: "123.456.789-09",
		},
		Name:  "John Doe",
		Email: "john@email.com",
	}

	// Criar order
	order := &entity.Order{
		ID:          "1",
		CustomerCPF: customer.GetCPF().Value,
		Status:      "confirmed",
		Items: []*entity.OrderItem{
			{
				ProductID: product.ID,
				Quantity:  1,
				Price:     product.Price,
				Status:    "confirmed",
			},
		},
	}

	order.CalculateTotal()
	assert.Equal(t, 10.0, order.Total)

	// Criar checkout
	checkout, err := entity.NewCheckout(order.ID, "mercado livre", "123456789", order.Total)
	assert.Nil(t, err)
	assert.NotNil(t, checkout)

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

	checkoutInput := &

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
