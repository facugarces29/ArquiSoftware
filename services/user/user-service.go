package services

import (
	userCliente "github.com/facugarces29/ArquiSoftware/clients/user"
	dto "github.com/facugarces29/ArquiSoftware/dto/user"
	model "github.com/facugarces29/ArquiSoftware/model/user"
)

type userService struct{}

type userServiceInterface interface {
	GetUserById(id int) dto.UserDto
	GetUsers() dto.UsersDto
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (s *userService) GetUserById(id int) dto.UserDto {

	var user model.User = userCliente.GetUserById(id)
	var userDto dto.UserDto

	if user.UserID == 0 {
		return userDto
	}
	userDto.Name = user.Name
	userDto.LastName = user.LastName
	userDto.UserName = user.UserName
	userDto.Id = user.UserID
	userDto.Email = user.Email
	userDto.Password = user.Pwd
	return userDto
}

func (s *userService) GetUsers() dto.UsersDto {
	var users model.Users = userCliente.GetUsers()
	var usersDto dto.UsersDto

	for _, user := range users {
		var userDto dto.UserDto
		userDto.Name = user.Name
		userDto.LastName = user.LastName
		userDto.UserName = user.UserName
		userDto.Id = user.UserID
		userDto.Email = user.Email
		userDto.Password = user.Pwd
		usersDto = append(usersDto, userDto)
	}

	return usersDto
}
