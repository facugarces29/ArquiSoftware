package clients

import (
	model "github.com/facugarces29/ArquiSoftware/model/user"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetUserById(id int) model.User {
	var user model.User

	Db.Where("user_id = ?", id).First(&user)
	log.Debug("User: ", user)

	return user
}

func GetUserByUsername(username string) (model.User, error) {
	var user model.User

	err := Db.Where("user_name = ?", username).First(&user).Error

	if err != nil {
		log.Println(err)
		return user, err
	}
	log.Debug("User: ", user)

	return user, nil
}

func GetUsers() model.Users {
	var users model.Users

	Db.Find(&users)

	log.Debug("Users: ", users)

	return users
}
