package model

import (
	"github.com/jinzhu/gorm"
)

type ProductCategory struct {
	gorm.Model
	//CategoryID uint `gorm:"column:category_id"`
	//ProductID  uint `gorm:"column:product_id"`
}

type ProductCategories []ProductCategory
