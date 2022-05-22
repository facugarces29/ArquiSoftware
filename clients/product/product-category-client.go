package client

import (
	model "Proyecto/ArquiSoftware/model/product"

	log "github.com/sirupsen/logrus"
)

func GetProductCategoryById(id int) model.ProductCategory {
	var productCategory model.ProductCategory

	Db.Where("id = ?", id).First(&productCategory)
	log.Debug("ProductCategory: ", productCategory)

	return productCategory
}

func GetProductCategories() model.ProductCategories {
	var productCategories model.ProductCategories

	Db.Find(&productCategories)

	log.Debug("ProductCategories: ", productCategories)

	return productCategories
}
