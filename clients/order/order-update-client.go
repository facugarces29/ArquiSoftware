package clients

import (
	"log"

	model "github.com/facugarces29/ArquiSoftware/model/order"
)

func UpdateOrder(order model.Order) (model.Order, error) {
	err := Db.Update(&order).Error

	if err != nil {
		log.Println(err)
		return order, err
	}

	return order, nil
}
