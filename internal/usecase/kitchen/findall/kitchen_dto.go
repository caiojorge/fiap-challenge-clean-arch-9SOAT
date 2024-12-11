package usecase

import (
	"time"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
)

type KitchenInputDTO struct {
	ID          string    `json:"id"`
	OrderID     string    `json:"order_id"`
	ItemID      string    `json:"item_order_id"`
	ProductName string    `json:"product_name"`
	Responsible string    `json:"responsible"`
	CreatedAt   time.Time `json:"created_at"`
}

func (dto *KitchenInputDTO) ToEntity() *entity.Kitchen {
	return &entity.Kitchen{
		ID:          dto.ID,
		OrderID:     dto.OrderID,
		ItemID:      dto.ItemID,
		ProductName: dto.ProductName,
		Responsible: dto.Responsible,
		CreatedAt:   dto.CreatedAt,
	}
}

type KitchenFindAllAOutputDTO struct {
	ID          string    `json:"id"`
	OrderID     string    `json:"order_id"`
	ItemID      string    `json:"item_order_id"`
	ProductName string    `json:"product_name"`
	Responsible string    `json:"responsible"`
	CreatedAt   time.Time `json:"created_at"`
}

func (dto *KitchenFindAllAOutputDTO) FromEntity(customer entity.Kitchen) {
	dto.ID = customer.ID
	dto.OrderID = customer.OrderID
	dto.ItemID = customer.ItemID
	dto.ProductName = customer.ProductName
	dto.Responsible = customer.Responsible
	dto.CreatedAt = customer.CreatedAt

}
