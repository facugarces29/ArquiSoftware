package clients

import (
	"errors"
	"log"
	modelOrder "proyecto/ArquiSoftware/model/order"
	modelProduct "proyecto/ArquiSoftware/model/product"

	"github.com/jinzhu/gorm"
)

func InsertOrderDetail(orderDetail modelOrder.OrderDetail) (modelOrder.OrderDetail, error) {

	err := Db.Create(&orderDetail).Error

	if err != nil {
		log.Println(err)
		return orderDetail, err
	}

	var product modelProduct.Product
	err2 := Db.Where("product_id = ?", orderDetail.ProductID).First(&product).Error
	if err2 != nil {
		log.Println(err2)
		return orderDetail, err2
	}

	err3 := Db.Where("product_id = ?", product.ProductID).Update("stock", gorm.Expr("stock - ?", orderDetail.Quantity)).Error
	if err3 != nil {
		log.Println(err3)
		return orderDetail, err3
	}

	return orderDetail, nil
}

func GetOrderDetailsByOrderId(id int) (modelOrder.OrderDetails, error) {
	var orderDetails modelOrder.OrderDetails

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

func DeleteOrderDetailsByOrderID(id int) error {
	var orderDetails modelOrder.OrderDetails

	err := Db.Where("order_id = ?", id).Delete(&orderDetails).Error

	if err != nil {
		return err
	}
	if len(orderDetails) != 0 {
		err = errors.New("cant delete all orderDetails")
		return err
	}
	return nil
}

func GetOrderDetails() (modelOrder.OrderDetails, error) {
	var orderDetails modelOrder.OrderDetails

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
