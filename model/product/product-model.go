package model

import (
	orderdetail "Proyecto/ArquiSoftware/model/order"
)

type Product struct {
	ProductID    uint    `gorm:"primary_key"`
	CategoryID   uint
	Name         string  `gorm:"type:varchar(255);not null;unique"`
	Description  string  `gorm:"type:varchar(255);not null"`
	Price        float64 `gorm:"type:decimal(10,2);not null"`
	Stock        int     `gorm:"type:integer(255);not null"`
	Image        string  `gorm:"type:blob;not null"`
	OrderDetails orderdetail.OrderDetails
	
	Category     Category
}

type Products []Product
