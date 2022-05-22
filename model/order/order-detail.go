package model

import "github.com/jinzhu/gorm"

type OrderDetail struct {
	gorm.Model
}

type OrderDetails []OrderDetail
