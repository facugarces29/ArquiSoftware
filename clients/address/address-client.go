package clients

import (
	model "proyecto/ArquiSoftware/model/address"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetAddressById(id int) (model.Address, error) {
	var address model.Address

	err := Db.Where("id = ?", id).First(&address).Error

	if err != nil {
		log.Println(err)
		return address, err
	}

	log.Debug("Address: ", address)

	return address, nil
}

func GetAddresses() (model.Addresses, error) {
	var addresses model.Addresses

	err := Db.Find(&addresses).Error

	if err != nil {
		log.Println(err)
		return addresses, err
	}

	log.Debug("Addresses: ", addresses)

	return addresses, nil
}
