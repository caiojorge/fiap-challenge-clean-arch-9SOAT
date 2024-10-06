package dto

import (
	"testing"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestOrderDTO_ToEntity(t *testing.T) {
	dto := OrderDTO{
		ID:          "order123",
		Items:       []*OrderItemDTO{{ID: "item1", ProductID: "product1", Quantity: 2, Price: 10.0, Status: "new"}},
		Total:       20.0,
		Status:      "pending",
		CustomerCPF: "12345678900",
	}

	order, err := dto.ToEntity()
	assert.NoError(t, err)
	assert.NotNil(t, order)
	assert.Equal(t, dto.ID, order.ID)
	assert.Equal(t, dto.Total, order.Total)
	assert.Equal(t, dto.Status, order.Status)
	assert.Equal(t, dto.CustomerCPF, order.CustomerCPF)
	assert.Len(t, order.Items, 1)
	assert.Equal(t, dto.Items[0].ID, order.Items[0].ID)
	assert.Equal(t, dto.Items[0].ProductID, order.Items[0].ProductID)
	assert.Equal(t, dto.Items[0].Quantity, order.Items[0].Quantity)
	assert.Equal(t, dto.Items[0].Price, order.Items[0].Price)
	assert.Equal(t, dto.Items[0].Status, order.Items[0].Status)
}

func TestOrderDTO_FromEntity(t *testing.T) {
	order := entity.Order{
		ID:          "order123",
		Items:       []*entity.OrderItem{{ID: "item1", ProductID: "product1", Quantity: 2, Price: 10.0, Status: "new"}},
		Total:       20.0,
		Status:      "pending",
		CustomerCPF: "12345678900",
	}

	dto := OrderDTO{}
	dto.FromEntity(order)

	assert.Equal(t, order.ID, dto.ID)
	assert.Equal(t, order.Total, dto.Total)
	assert.Equal(t, order.Status, dto.Status)
	assert.Equal(t, order.CustomerCPF, dto.CustomerCPF)
	assert.Len(t, dto.Items, 1)
	assert.Equal(t, order.Items[0].ID, dto.Items[0].ID)
	assert.Equal(t, order.Items[0].ProductID, dto.Items[0].ProductID)
	assert.Equal(t, order.Items[0].Quantity, dto.Items[0].Quantity)
	assert.Equal(t, order.Items[0].Price, dto.Items[0].Price)
	assert.Equal(t, order.Items[0].Status, dto.Items[0].Status)
}
