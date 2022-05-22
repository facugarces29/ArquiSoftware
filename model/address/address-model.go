package model

import (
	order "Proyecto/ArquiSoftware/model/order"

	"github.com/jinzhu/gorm"
)

type Address struct {
	gorm.Model
	UserID       uint   `gorm:"column:user_id"`
	State        string `gorm:"type:varchar(255);not null"`
	City         string `gorm:"type:varchar(255);not null"`
	Zip          int    `gorm:"type:integer(11);not null"`
	Addressline  string `gorm:"type:integer(11);not null"`
	order.Orders `gorm:"foreignkey:address_id"`
}

type Addresses []Address
