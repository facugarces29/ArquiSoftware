package controllers

import (
	"net/http"

	service "proyecto/ArquiSoftware/services/login"

	dto "proyecto/ArquiSoftware/dto/login"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Login(c *gin.Context) {
	var loginReqDto dto.LoginRequestDto
	var loginRespDto dto.LoginResponseDto
	err1 := c.BindJSON(&loginReqDto)

	// Error Parsing json param
	if err1 != nil {
		log.Error(err1.Error())
		c.JSON(http.StatusBadRequest, err1.Error())
		return
	}

	loginRespDto, err2 := service.LoginService.Login(loginReqDto)
	// Error del Insert

	if err2 != nil {
		c.JSON(http.StatusBadRequest, loginRespDto)
		return
	}

	c.JSON(http.StatusCreated, loginRespDto)
}
