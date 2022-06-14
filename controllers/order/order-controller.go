package controllers

import (
	"net/http"
	"strconv"

	service "github.com/facugarces29/ArquiSoftware/services/order"

	"github.com/gin-gonic/gin"
)

func GetOrdersByUserId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	ordersDto, err := service.OrderService.GetOrdersByUserId(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, ordersDto)
		return
	}

	c.JSON(http.StatusOK, ordersDto)
}
