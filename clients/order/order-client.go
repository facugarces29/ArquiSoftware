package client

import (
	model "Proyecto/ArquiSoftware/model/order"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetOrderById(id int) model.Order {
	var order model.Order

	Db.Where("id = ?", id).First(&order)
	log.Debug("Order: ", order)

	return order
}

func GetOrders() model.Orders {
	var orders model.Orders

	Db.Find(&orders)

	log.Debug("Orders: ", orders)

	return orders
}
