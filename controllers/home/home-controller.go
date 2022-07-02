package controllers

import (
	"net/http"

	service "proyecto/ArquiSoftware/services/home"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetHomeProducts(c *gin.Context) {
	productsDto, err := service.HomeService.GetHomeProducts()

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, productsDto)
		return
	}

	c.JSON(http.StatusOK, productsDto)
}
