package clients

import (
	"errors"
	"log"

	model "proyecto/ArquiSoftware/model/order"
)

func InsertOrderDetail(orderDetail model.OrderDetail) (model.OrderDetail, error) {

	err := Db.Create(&orderDetail).Error

	if err != nil {
		log.Println(err)
		return orderDetail, err
	}
	var productId int

	err2 := Db.Where("product_id = ?", orderDetail.ProductID).First(productId).Error
	if err2 != nil {
		log.Println(err2)
		return orderDetail, err2
	}

	return orderDetail, nil
}

func GetOrderDetailsByOrderId(id int) (model.OrderDetails, error) {
	var orderDetails model.OrderDetails

	err := Db.Where("order_id = ?", id).Find(&orderDetails).Error

	if err != nil {
		log.Println(err)
		return orderDetails, err
	}

	if len(orderDetails) == 0 {
		err = errors.New("nothing found")
		return orderDetails, err
	}

	return orderDetails, nil
}

func GetOrderDetails() (model.OrderDetails, error) {
	var orderDetails model.OrderDetails

	err := Db.Find(&orderDetails).Error

	if err != nil {
		log.Println(err)
		return orderDetails, err
	}

	if len(orderDetails) == 0 {
		err = errors.New("nothing found")
		return orderDetails, err
	}

	return orderDetails, nil
}
