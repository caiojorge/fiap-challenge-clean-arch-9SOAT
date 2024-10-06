package dto

import (
	"time"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

type OrderDTO struct {
	ID          string          `json:"id"`
	Items       []*OrderItemDTO `json:"items"`
	Total       float64         `json:"total"`
	Status      string          `json:"status"`
	CustomerCPF string          `json:"customercpf"`
	CreatedAt   time.Time       `json:"created_at"`
}

func (dto *OrderDTO) ToEntity() (*entity.Order, error) {
	items := dto.fromOrderItemDTO()
	order := entity.Order{
		ID:          dto.ID,
		Items:       items,
		Total:       dto.Total,
		Status:      dto.Status,
		CustomerCPF: dto.CustomerCPF,
	}

	return &order, nil
}

func (dto *OrderDTO) FromEntity(order entity.Order) {
	dto.fromOrderItemEntity(order.Items)
	dto.ID = order.ID
	dto.Status = order.Status
	dto.CustomerCPF = order.CustomerCPF
	dto.Total = order.Total
	dto.CreatedAt = order.CreatedAt
}

func (dto *OrderDTO) fromOrderItemDTO() []*entity.OrderItem {
	items := []*entity.OrderItem{}
	for _, value := range dto.Items {
		entity := entity.OrderItem{
			ID:        value.ID,
			ProductID: value.ProductID,
			Quantity:  value.Quantity,
			Price:     value.Price,
			Status:    value.Status,
		}
		items = append(items, &entity)
	}

	return items

}

func (dto *OrderDTO) fromOrderItemEntity(items []*entity.OrderItem) {
	for _, value := range items {
		item := OrderItemDTO{
			ID:        value.ID,
			ProductID: value.ProductID,
			Quantity:  value.Quantity,
			Price:     value.Price,
			Status:    value.Status,
		}
		dto.Items = append(dto.Items, &item)
	}

}
