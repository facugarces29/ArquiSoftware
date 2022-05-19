package services

import (
	userCliente "Proyecto/ArquiSoftware/clients"
	"Proyecto/ArquiSoftware/dto"
	"Proyecto/ArquiSoftware/model"
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

	var user model.ModelUser = userCliente.GetUserById(id)
	var userDto dto.UserDto

	if user.Id == 0 {
		return userDto
	}
	userDto.Name = user.Name
	userDto.LastName = user.LastName
	userDto.UserName = user.UserName
	userDto.Id = user.Id
	return userDto
}

func (s *userService) GetUsers() dto.UsersDto {
	var users model.ModelUsers = userCliente.GetUsers()
	var usersDto dto.UsersDto

	for _, user := range users {
		var userDto dto.UserDto
		userDto.Name = user.Name
		userDto.LastName = user.LastName
		userDto.UserName = user.Name
		userDto.Id = user.Id

		usersDto = append(usersDto, userDto)
	}

	return usersDto
}
