package clients

import (
	addressModel "proyecto/ArquiSoftware/model/address"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

type AddressClientInterface interface {
	GetAddressByUserId(id int) addressModel.Address
	GetAddressById(id int) addressModel.Address
}

func GetAddressByUserId(id int) (addressModel.Address, error) {
	var address addressModel.Address

	err := Db.Where("user_id = ?", id).First(&address).Error

	if err != nil {
		log.Println(err)
		return address, err
	}

	log.Debug("Address: ", address)

	return address, nil
}

func GetAddressById(id int) (addressModel.Address, error) {
	var address addressModel.Address

	err := Db.Where("id = ?", id).First(&address).Error

	if err != nil {
		log.Println(err)
		return address, err
	}

	log.Debug("Address: ", address)

	return address, nil
}
