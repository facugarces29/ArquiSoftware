package controllers

import (
	"log"
	"net/http"
	"strconv"

	dto "proyecto/ArquiSoftware/dto/order"
	service "proyecto/ArquiSoftware/services/order"

	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) {
	ordersDto, err := service.OrderService.GetOrders()

	if err != nil {
		c.JSON(http.StatusBadRequest, ordersDto)
		return
	}

	c.JSON(http.StatusOK, ordersDto)
}

func GetOrdersByUserId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("userId"))

	ordersDto, err := service.OrderService.GetOrdersByUserId(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, ordersDto)
		return
	}

	c.JSON(http.StatusOK, ordersDto)
}

func InsertOrder(c *gin.Context) {
	var orderDto dto.OrderDto
	var insertOrderDto dto.InsertOrderDto

	err := c.BindJSON(&orderDto)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, orderDto)
		return
	}

	insertOrderDto.UserId = orderDto.UserId

	orderDto, err2 := service.OrderService.InsertOrder(insertOrderDto)

	if err2 != nil {
		log.Println(err2)
		c.JSON(http.StatusBadRequest, orderDto)
		return
	}

	c.JSON(http.StatusCreated, orderDto)
}
