package clients

import (
	"log"

	model "github.com/facugarces29/ArquiSoftware/model/order"
)

func InsertOrderDetail(orderDetail model.OrderDetail) (model.OrderDetail, error) {

	err := Db.Create(&orderDetail).Error

	if err != nil {
		log.Println(err)
		return orderDetail, err
	}

	return orderDetail, nil
}
