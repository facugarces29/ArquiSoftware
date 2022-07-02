package clients

import (
	"errors"
	"log"

	model "proyecto/ArquiSoftware/model/product"

	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func GetProductsBySearchParam(param string) (model.Products, error) {
	var products model.Products

	param = "%" + param + "%"

	err := Db.Find(&products, "name LIKE ? OR description LIKE ?", param, param).Error

	if err != nil {
		log.Println(err)
		return products, err
	}

	if len(products) == 0 {
		err := errors.New("nothing found")
		return products, err
	}

	return products, nil
}
