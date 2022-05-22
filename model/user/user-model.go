package model

import (
	address "Proyecto/ArquiSoftware/model/address"
	order "Proyecto/ArquiSoftware/model/order"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name             string `gorm:"type:varchar(255);not null"`
	LastName         string `gorm:"type:varchar(255);not null"`
	UserName         string `gorm:"type:varchar(255);not null;unique"`
	Email            string `gorm:"type:varchar(255);not null"`
	Pwd              string `gorm:"type:varchar(255);not null"`
	address.Adresses `gorm:"foreignkey:user_id"`
	order.Orders     `gorm:"foreignkey:user_id"`
}

type Users []User

/*`id` int(11) NOT NULL,
  `name` varchar(255) COLLATE utf8_bin NOT NULL,
  `lastName` varchar(255) COLLATE utf8_bin NOT NULL,
  `username`	varchar(255) COLLATE utf8_bin NOT NULL unique,
  `pwd` varchar(255) COLLATE utf8_bin NOT NULL,
  `mail` varchar(255) COLLATE utf8_bin NOT NULL*/
