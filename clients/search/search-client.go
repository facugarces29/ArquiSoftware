package clients

import (
	model "Proyecto/ArquiSoftware/model/product"

	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func GetProductsBySearchParam(param string) (model.Products, error) {
	var products model.Products

	param = "%" + param + "%"

	Db.Where("description LIKE ?", param).Find(&products)

	Db.Where("name LIKE ?", param).Find(&products)

	//como saber si el modelo esta vacio?? ---> error

	return products, nil
}
