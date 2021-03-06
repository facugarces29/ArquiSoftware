package clients

import (
	model "proyecto/ArquiSoftware/model/product"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetProductById(id int) (model.Product, error) {
	var product model.Product

	err := Db.Where("product_id = ?", id).First(&product).Error
	log.Debug("Product: ", product)

	if err != nil {
		log.Println(err)
		return product, err
	}

	return product, err
}

func GetProducts() (model.Products, error) {
	var products model.Products

	err := Db.Find(&products).Error

	log.Debug("Products: ", products)

	if err != nil {
		log.Println(err)
		return products, err
	}

	return products, err
}

func DecreaseProductById(id int, quantity int) {
	
	err := Db.Where("product_id = ?", id).Update("stock", gorm.Expr("stock - ?", quantity))

	if err != nil {
		log.Println(err)
		return
	}
	return
}
