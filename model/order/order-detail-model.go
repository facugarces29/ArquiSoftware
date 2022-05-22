package model

import "github.com/jinzhu/gorm"

type OrderDetail struct {
	gorm.Model
	OrderID   uint    `gorm:"column:order_id"`
	ProductID uint    `gorm:"column:product_id"`
	Price     float64 `gorm:"type:decimal(10,2);not null"`
	Quantity  int     `gorm:"type:integer(255);not null"`
}

type OrderDetails []OrderDetail
