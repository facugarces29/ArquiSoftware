package clients

import (
	model "github.com/facugarces29/ArquiSoftware/model/order"

	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func GetOrdersByUserId(id int) model.Order {
	var order model.Order

	return order
}
