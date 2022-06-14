package clients

import (
	model "github.com/facugarces29/ArquiSoftware/model/product"

	log "github.com/sirupsen/logrus"
)

func GetCategoryById(id int) (model.Category, error) {
	var category model.Category

	err := Db.Where("category_id = ?", id).First(&category).Error

	if err != nil {
		log.Println(err)
		return category, err
	}

	log.Debug("Category: ", category)

	return category, nil
}

func GetCategories() (model.Categories, error) {
	var categories model.Categories

	err := Db.Find(&categories).Error

	if err != nil {
		log.Println(err)
		return categories, err
	}

	log.Debug("Products: ", categories)

	return categories, nil
}
