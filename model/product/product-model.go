package model

import (
	//order "Proyecto/ArquiSoftware/model/order"

	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `gorm:"type:varchar(255);not null;unique"`
	Description string  `gorm:"type:varchar(255);not null"`
	Price       float64 `gorm:"type:decimal(10,2);not null"`
	Stock       int     `gorm:"type:integer(255);not null"`
	Image       string  `gorm:"type:blob;not null"`
	//order.OrderDetails `gorm:"foreignkey:product_id"`
	//ProductCategories  `gorm:"foreignkey:product_id"`
}

type Products []Product
