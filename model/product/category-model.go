package model

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:varchar(255);not null"`
	//ProductCategories `gorm:"foreignkey:category_id"`
}

type Categories []Category
