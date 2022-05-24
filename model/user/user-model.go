package model

import (
	address "Proyecto/ArquiSoftware/model/address"

	"github.com/jinzhu/gorm"
	//order "Proyecto/ArquiSoftware/model/order"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null"`
	LastName string `gorm:"type:varchar(255);not null"`
	UserName string `gorm:"type:varchar(255);not null;unique"`
	Email    string `gorm:"type:varchar(255);not null"`
	Pwd      string `gorm:"type:varchar(255);not null"`
	address.Addresses
	//order.Orders      `gorm:"foreignkey:user_id"`
}

type Users []User
