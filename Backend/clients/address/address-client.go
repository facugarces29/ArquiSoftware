package clients

import (
	addressModel "proyecto/ArquiSoftware/model/address"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

type AddressClientInterface interface {
	GetAddressByUserId(id int) addressModel.Addresses
	GetAddressById(id int) addressModel.Address
}

func GetAddressByUserId(id int) (addressModel.Addresses, error) {
	var addresses addressModel.Addresses

	err := Db.Where("user_id = ?", id).Find(addresses).Error

	if err != nil {
		log.Println(err)
		return addresses, err
	}

	log.Debug("Address: ", addresses)

	return addresses, nil
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
