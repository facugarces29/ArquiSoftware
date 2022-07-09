package clients

import (
	model "proyecto/ArquiSoftware/model/product"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetHomeProducts() (model.Products, error) {
	var products model.Products

	err := Db.Order("rand()").Limit(20).Find(&products).Error

	log.Debug("Products: ", products)

	if err != nil {
		log.Println(err)
		return products, err
	}

	return products, err
}
