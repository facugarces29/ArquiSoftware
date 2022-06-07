package clients
import (
	model "github.com/facugarces29/ArquiSoftware/model/order"

	log "github.com/sirupsen/logrus"
)

func GetOrderDetailById(id int) model.OrderDetail {
	var orderDetail model.OrderDetail

	Db.Where("order_detail_id = ?", id).First(&orderDetail)
	log.Debug("Order Detail: ", orderDetail)

	return orderDetail
}

func GetOrderDetails() model.OrderDetails {
	var orderDetails model.OrderDetails

	Db.Find(&orderDetails)

	log.Debug("Order details: ", orderDetails)

	return orderDetails
}
