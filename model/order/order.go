package model

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	
}

type Orders []Order
