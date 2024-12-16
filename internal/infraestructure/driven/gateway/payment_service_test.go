package service

import (
	"context"
	"testing"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestPaymentService(t *testing.T) {

	checkout, err := entity.NewCheckout("order123", "gatewayteste", "gatewaytoken1234567890", 100)
	assert.Nil(t, err)
	assert.NotNil(t, checkout)

	product, err := entity.NewProduct("prod123", "product", "product category", 100)
	assert.Nil(t, err)
	assert.NotNil(t, product)

	orderItem, err := entity.NewOrderItem(product.ID, 1, 100)
	assert.Nil(t, err)
	assert.NotNil(t, orderItem)
	assert.NotEmpty(t, orderItem.ProductID)

	order, err := entity.NewOrder("order123", []*entity.OrderItem{orderItem})
	assert.Nil(t, err)
	assert.NotNil(t, order)
	//order.CalculateTotal()

	gateway := NewFakePaymentService()
	payment, err := gateway.ConfirmPayment(context.Background(), checkout, order, []*entity.Product{product}, "http://localhost:8080/checkout/notification", 1)
	// payment, err := entity.NewPayment(
	// 	*checkout,
	// 	*order,
	// 	[]*entity.Product{product},
	// 	"http://localhost:8080/checkout/notification",
	// 	1,
	// )
	assert.Nil(t, err)
	assert.NotNil(t, payment)
	assert.NotNil(t, payment.ID)
	assert.NotNil(t, payment.ExternalReference)
	assert.Equal(t, payment.CheckoutID, checkout.ID)
	assert.Equal(t, payment.ExternalReference, order.ID)
	assert.Equal(t, len(payment.Items), len(order.Items))

}
