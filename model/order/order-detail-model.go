package model

import "github.com/jinzhu/gorm"

type OrderDetail struct {
	gorm.Model
	//OrderID   uint
	//ProductID uint
	Price    float64 `gorm:"type:decimal(10,2);not null"`
	Quantity int     `gorm:"type:integer(255);not null"`
}

type OrderDetails []OrderDetail
