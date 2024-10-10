package model

import "time"

// Checkout belongs to an order. It is the final step of the order process
type Checkout struct {
	ID                   string    `gorm:"type:char(36);primaryKey"`
	OrderID              string    `gorm:"not null;index:idx_order_id"`
	Order                Order     `gorm:"foreignKey:OrderID"`
	Gateway              string    `gorm:"not null;type:varchar(255)"`
	GatewayID            string    `gorm:"not null;type:varchar(255)"`
	GatewayTransactionID string    `gorm:"type:varchar(255)"`
	CustomerCPF          string    `gorm:"type:char(11);not null;index:idx_customer_cpf"`
	Total                float64   `gorm:"not null"`
	CreatedAt            time.Time `gorm:"not null"`
}
