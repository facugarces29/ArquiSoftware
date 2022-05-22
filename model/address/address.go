package model

import "github.com/jinzhu/gorm"

type Address struct {
	gorm.Model
	State       string `gorm:"type:varchar(255);not null"`
	City        string `gorm:"type:varchar(255);not null"`
	Zip         int    `gorm:"type:integer(11);not null"`
	Addressline string `gorm:"type:integer(11);not null"`
}

type Adresses []Address
