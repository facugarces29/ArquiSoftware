package clients

import (
	"log"

	model "github.com/facugarces29/ArquiSoftware/model/order"

	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func InsertOrder(order model.Order) (model.Order, error) {
	err := Db.Create(&order).Error

	if err != nil {
		log.Println(err)
		return order, err
	}

	return order, nil
}

/*func GetOrdersByUserId(id int) model.Order {
	var order model.Order

	return order
}*/
