package controllers

import (
	service "Proyecto/ArquiSoftware/services/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetUserById(c *gin.Context) {
	log.Debug("User id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var userDto = service.UserService.GetUserById(id)

	c.JSON(http.StatusOK, userDto)
}

func GetUsers(c *gin.Context) {
	var usersDto = service.UserService.GetUsers()

	c.JSON(http.StatusOK, usersDto)
}
