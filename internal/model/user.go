package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    	string `json:"username"`
	Email    	string `json:"email" gorm:"unique;not null"`
	Password 	string `json:"-" gorm:"not null"`
	Role	 	string `json:"role" gorm:"default:STAFF;check:role IN ('ADMIN', 'STAFF')"`
}