package controllers

import (
	"log"
	"net/http"
	"strconv"

	dto "proyecto/ArquiSoftware/dto/order"
	service "proyecto/ArquiSoftware/services/order"

	"github.com/gin-gonic/gin"
)

func InsertOrderDetail(c *gin.Context) {
	var orderDetailDto dto.OrderDetailDto
	var orderDetailResponseDto dto.OrderDetailDto

	err := c.BindJSON(&orderDetailDto)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, orderDetailDto)
		return
	}

	orderDetailResponseDto, err2 := service.OrderDetailService.InsertOrderDetail(orderDetailDto)

	if err2 != nil {
		log.Println(err2)
		c.JSON(http.StatusBadRequest, orderDetailResponseDto)
		return
	}

	c.JSON(http.StatusCreated, orderDetailResponseDto)
}

func GetOrderDetailsByOrderId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("orderId"))

	orderDetailsDto, err := service.OrderDetailService.GetOrderDetailsByOrderId(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, orderDetailsDto)
		return
	}

	c.JSON(http.StatusOK, orderDetailsDto)
}

func GetOrderDetails(c *gin.Context) {
	orderDetailsDto, err := service.OrderDetailService.GetOrderDetails()

	if err != nil {
		c.JSON(http.StatusBadRequest, orderDetailsDto)
		return
	}

	c.JSON(http.StatusOK, orderDetailsDto)
}
