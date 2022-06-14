package clients

import (
	"errors"
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

func GetOrdersByUserId(idUser int) (model.Orders, error) {
	var orders model.Orders

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

/*func GetOrdersByUserId(id int) model.Order {
	var order model.Order

	return order




** db.Model(&User{}).Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&result{})
// SELECT users.name, emails.email FROM `users` left join emails on emails.user_id = users.id

}*/
