package clients

import (
	"errors"
	"log"

	orderModel "proyecto/ArquiSoftware/model/order"
	userModel "proyecto/ArquiSoftware/model/user"

	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

// InsertOrder inserts an order in database
func InsertOrder(order orderModel.Order) (orderModel.Order, error) {

	var user userModel.User

	userId := order.UserID
	err := Db.Where("user_id = ?", userId).First(&user).Error

	if err != nil {
		log.Println(err)
		return order, err
	}

	if user.UserID == 0 {
		err = errors.New("user not found")
		log.Println(err)
		return order, err
	}

	err = Db.Create(&order).Error

	if err != nil {
		log.Println(err)
		return order, err
	}

	return order, nil
}

// UpdateOrder updates an order in database

// InsertOrder returns all orders in database
func GetOrders() (orderModel.Orders, error) {
	var orders orderModel.Orders

	err := Db.Find(&orders).Error

	if err != nil {
		log.Println(err)
		return orders, err
	}

	if len(orders) == 0 {
		err = errors.New("nothing found")
		log.Println(err)
		return orders, err
	}

	return orders, nil
}

// GetOrdersByUserId returns orders from user id
func GetOrdersByUserId(idUser int) (orderModel.Orders, error) {

	var orders orderModel.Orders

	err := Db.Where("user_id = ?", idUser).Find(&orders).Error

	if err != nil {
		log.Println(err)
		return orders, err
	}

	if len(orders) == 0 {
		err = errors.New("nothing found")
		return orders, err
	}

	return orders, nil

}

func DeleteOrderById(id int) error {
	var order orderModel.Order
	err := Db.Where("id = ?", id).Delete(&order).Error

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
