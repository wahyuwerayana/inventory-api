package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	SKU         string    `json:"sku" gorm:"unique;not null"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description"`
	Price       float64   `json:"price" gorm:"not null"`
	Stock       int       `json:"stock" gorm:"default:0"`
	CategoryID  uint      `json:"category_id" gorm:"not null"`
	SupplierID  uint      `json:"supplier_id" gorm:"not null"`

	Category Category `json:"category" gorm:"foreignKey:CategoryID"`
	Supplier Supplier `json:"supplier" gorm:"foreignKey:SupplierID"`
}