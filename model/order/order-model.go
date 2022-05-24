package model

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	UserID    uint    `gorm:"column:user_id"`
	AddressID uint    `gorm:"column:address_id"`
	Amount    float64 `gorm:"type:decimal(10,2);not null"`
	//OrderDetails `gorm:"foreignkey:order_id"`
}

type Orders []Order
