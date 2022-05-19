package services

import (
	"Proyecto/ArquiSoftware/dto"
)

type userService struct{}

type userServiceInterface interface {
	GetUserById(id int) dto.UserDto
	GetUsers() dto.UsersDto
	InsertUser(userDto dto.UserDto) dto.UserDto
}


