package controllers

import (
	service "Proyecto/ArquiSoftware/services/login"
	"net/http"

	dto "Proyecto/ArquiSoftware/dto/login"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Login(c *gin.Context) {
	var loginDto dto.LoginDto
	err := c.BindJSON(&loginDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userDto := service.LoginService.Login(loginDto)
	// Error del Insert

	c.JSON(http.StatusCreated, userDto)
}
