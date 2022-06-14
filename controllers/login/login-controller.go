package controllers

import (
	"net/http"

	service "github.com/facugarces29/ArquiSoftware/services/login"

	dto "github.com/facugarces29/ArquiSoftware/dto/login"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Login(c *gin.Context) {
	var loginDto dto.LoginDto
	err1 := c.BindJSON(&loginDto)

	// Error Parsing json param
	if err1 != nil {
		log.Error(err1.Error())
		c.JSON(http.StatusBadRequest, err1.Error())
		return
	}

	userDto, err2 := service.LoginService.Login(loginDto)
	// Error del Insert

	if err2 != nil {
		c.JSON(http.StatusBadRequest, userDto)
		return
	}

	c.JSON(http.StatusCreated, userDto)
}
