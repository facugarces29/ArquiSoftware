package controllers

import (
	"net/http"
	"strconv"

	service "proyecto/ArquiSoftware/services/address"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetAddressByUserId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	addressDto, err := service.AddressService.GetAddressByUserId(id)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, addressDto)
		return
	}

	c.JSON(http.StatusOK, addressDto)
}
