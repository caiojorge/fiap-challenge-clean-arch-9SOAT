package model

import "time"

type Kitchen struct {
	ID          string    `gorm:"type:char(36);primaryKey"`
	OrderID     string    `gorm:"type:varchar(36);not null"` // FK
	ItemID      string    `gorm:"type:varchar(36);not null"` // FK
	ProductName string    `gorm:"type:varchar(255);not null"`
	Responsible string    `gorm:"type:varchar(36);not null"`
	CreatedAt   time.Time `gorm:"not null"`
}
