package clients

import (
	"errors"
	"log"

	model "github.com/facugarces29/ArquiSoftware/model/product"

	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func GetProductsBySearchParam(param string) (model.Products, error) {
	var products model.Products

	param = "%" + param + "%"

	res1 := Db.Where("description LIKE ?", param).Find(&products)

	err := res1.Error
	if err != nil {
		log.Println(err)
		return products, err
	}

	res2 := Db.Where("name LIKE ?", param).Find(&products)

	err = res2.Error

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
