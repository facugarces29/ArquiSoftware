package clients

import (
	"errors"
	"log"

	orderModel "github.com/facugarces29/ArquiSoftware/model/order"
	userModel "github.com/facugarces29/ArquiSoftware/model/user"

	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

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

//realizar un get order by id de order

/*func GetOrdersByUserId(id int) model.Order {
	var order model.Order

	return order




** db.Model(&User{}).Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&result{})
// SELECT users.name, emails.email FROM `users` left join emails on emails.user_id = users.id

}*/
