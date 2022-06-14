package controllers

import (
	"net/http"
	"strconv"

	service "github.com/facugarces29/ArquiSoftware/services/user"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetUserById(c *gin.Context) {
	log.Debug("User id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var userDto = service.UserService.GetUserById(id)

	c.JSON(http.StatusOK, userDto)
}

func GetUserByUsername(c *gin.Context) {
	log.Debug("User username to load: " + c.Param("username"))

	username := c.Param("username")
	userDto, err := service.UserService.GetUserByUsername(username)

	if err != nil {
		c.JSON(http.StatusBadRequest, userDto)
	}

	c.JSON(http.StatusOK, userDto)
}

func GetUsers(c *gin.Context) {
	var usersDto = service.UserService.GetUsers()

	c.JSON(http.StatusOK, usersDto)
}
