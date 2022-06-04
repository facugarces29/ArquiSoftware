package models

import (
	user "Proyecto/ArquiSoftware/model/user"

	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	UserID       uint
	Amount       float64 `gorm:"type:decimal(10,2);not null"`
	User         user.User
	OrderDetails OrderDetails
}

type Orders []Order
