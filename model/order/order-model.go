package models

import (
	user "proyecto/ArquiSoftware/model/user"

	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	UserID       uint
	User         user.User
	OrderDetails OrderDetails
}

type Orders []Order
