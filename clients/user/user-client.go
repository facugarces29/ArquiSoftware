package client

import (
	model "Proyecto/ArquiSoftware/model/user"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetUserById(id int) model.User {
	var user model.User

	Db.Where("id = ?", id).First(&user)
	log.Debug("User: ", user)

	return user
}

func GetUsers() model.Users {
	var users model.Users

	Db.Find(&users)

	log.Debug("Users: ", users)

	return users
}
