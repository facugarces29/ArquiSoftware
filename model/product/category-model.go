package model

type Category struct {
	CategoryID uint   `gorm:"primary_key"`
	Name       string `gorm:"type:varchar(255);not null"`
}

type Categories []Category
