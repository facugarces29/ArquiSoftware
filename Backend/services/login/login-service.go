package services

import (
	"errors"

	userCliente "proyecto/ArquiSoftware/clients/user"
	loginDto "proyecto/ArquiSoftware/dto/login"
	userDto "proyecto/ArquiSoftware/dto/user"

	log "github.com/sirupsen/logrus"
)

type loginService struct{}

type loginServiceInterface interface {
	Login(loginDto.LoginDto) (userDto.UserDto, error)
}

var (
	LoginService loginServiceInterface
)

func init() {
	LoginService = &loginService{}
}

func (s *loginService) Login(login loginDto.LoginDto) (userDto.UserDto, error) {

	username := login.Username
	pass := login.Password
	log.Debug("pass", pass)
	user, err := userCliente.GetUserByUsername(username)
	var userDto userDto.UserDto

	if err != nil {
		log.Println(err)
		return userDto, err
	}

	if user.Pwd != pass {
		err = errors.New("username or password incorrect")
		log.Println(err)
		return userDto, err
	}

	userDto.Name = user.Name
	userDto.LastName = user.LastName
	userDto.UserName = user.UserName
	userDto.Id = user.UserID
	userDto.Email = user.Email
	userDto.Password = user.Pwd

	return userDto, nil

}
