package clients

import (
	addressModel "proyecto/ArquiSoftware/model/address"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetAddressById(id int) (addressModel.Address, error) {
	var address addressModel.Address

	err := Db.Where("user_id = ?", id).First(&address).Error

	if err != nil {
		log.Println(err)
		return address, err
	}

	log.Debug("Address: ", address)

	return address, nil
}

func GetAddresses() (addressModel.Addresses, error) {
	var addresses addressModel.Addresses

	err := Db.Find(&addresses).Error

	if err != nil {
		log.Println(err)
		return addresses, err
	}

	log.Debug("Addresses: ", addresses)

	return addresses, nil
}
