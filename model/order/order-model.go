package models

import (
	user "github.com/facugarces29/ArquiSoftware/model/user"

	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	UserID       uint
	User         user.User
	OrderDetails OrderDetails
}

type Orders []Order
