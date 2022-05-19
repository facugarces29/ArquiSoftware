package model

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255);not null;unique"`
	Description string `gorm:"type:varchar(255);not null"`
	Price       string `gorm:"type:decimal(10,2);not null"`
	Stock       string `gorm:"type:integer(255);not null"`
	Image       string `gorm:"type:blob;not null"`
}

type Users []Product
