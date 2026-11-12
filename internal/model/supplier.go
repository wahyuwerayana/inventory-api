package model

import (
	"gorm.io/gorm"
)

type Supplier struct {
	gorm.Model
	Name      string    `json:"name" gorm:"not null"`
	Contact   string    `json:"contact"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
}