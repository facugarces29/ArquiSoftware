package client

import (
	"Proyecto/ArquiSoftware/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetUserById(id int) model.ModelUser {
	var user model.ModelUser

	Db.Where("id = ?", id).First(&user)
	log.Debug("User: ", user)

	return user
}

func GetUsers() model.ModelUsers {
	var users model.ModelUsers

	Db.Find(&users)

	log.Debug("Users: ", users)

	return users
}
