package models

import (
	user "proyecto/ArquiSoftware/model/user"
)

type Address struct {
	AddressID   uint   `gorm:"primary_key"`
	State       string `gorm:"type:varchar(255);not null"`
	City        string `gorm:"type:varchar(255);not null"`
	Zip         int    `gorm:"type:integer(11);not null"`
	Addressline string `gorm:"type:varchar(255);not null"`
	UserID      uint   `gorm:"type:integer(11);not null; unique"`
	User        user.User
}

type Addresses []Address
