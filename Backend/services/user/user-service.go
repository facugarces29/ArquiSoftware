package services

import (
	userCliente "proyecto/ArquiSoftware/clients/user"
	dto "proyecto/ArquiSoftware/dto/user"
)

type userService struct{}

type userServiceInterface interface {
	GetUserById(id int) (dto.UserDto, error)
	GetUserByUsername(string) (dto.UserDto, error)
	GetUsers() (dto.UsersDto, error)
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (s *userService) GetUserById(id int) (dto.UserDto, error) {

	user, err := userCliente.GetUserById(id)
	var userDto dto.UserDto

	if err != nil {
		return userDto, err
	}
	if user.UserID == 0 {
		return userDto, nil
	}
	userDto.Name = user.Name
	userDto.LastName = user.LastName
	userDto.UserName = user.UserName
	userDto.Id = user.UserID
	userDto.Email = user.Email
	userDto.Password = user.Pwd
	return userDto, nil
}

func (s *userService) GetUsers() (dto.UsersDto, error) {
	users, err := userCliente.GetUsers()
	var usersDto dto.UsersDto

	if err != nil {
		return usersDto, err
	}

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

	return usersDto, nil
}

func (s *userService) GetUserByUsername(username string) (dto.UserDto, error) {

	user, err := userCliente.GetUserByUsername(username)
	var userDto dto.UserDto

	if err != nil {
		return userDto, err
	}

	userDto.Name = user.Name
	userDto.LastName = user.LastName
	userDto.UserName = user.UserName
	userDto.Id = user.UserID
	userDto.Email = user.Email
	return userDto, nil
}
