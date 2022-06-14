package controllers

import (
	service "github.com/facugarces29/ArquiSoftware/services/product"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetProductById(c *gin.Context) {
	log.Debug("User id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	productDto, err := service.ProductService.GetProductById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, productDto)
		return
	}
	
	c.JSON(http.StatusOK, productDto)
}