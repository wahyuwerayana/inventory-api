package model

import (
	"gorm.io/gorm"
)

type StockMovement struct {
	gorm.Model
	ProductID uint      `json:"product_id" gorm:"not null"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	Type      string    `json:"type" gorm:"not null;check:type IN ('IN', 'OUT')"` // "in" or "out"
	Quantity  int       `json:"quantity" gorm:"not null"`
	Notes     string    `json:"notes"`

	Product Product `json:"product" gorm:"foreignKey:ProductID"`
	User User `json:"user" gorm:"foreignKey:UserID"`
}