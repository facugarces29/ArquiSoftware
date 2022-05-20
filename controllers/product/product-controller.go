package product

import (
	service "Proyecto/ArquiSoftware/services/product"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetProductById(c *gin.Context) {
	log.Debug("User id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var productDto = service.ProductService.GetProductById(id)

	c.JSON(http.StatusOK, productDto)
}

func GetProducts(c *gin.Context) {
	var productsDto = service.ProductService.GetProducts()

	c.JSON(http.StatusOK, productsDto)
}
