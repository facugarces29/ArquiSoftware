package services

import (
	"log"

	orderDetailCliente "proyecto/ArquiSoftware/clients/order"
	dto "proyecto/ArquiSoftware/dto/order"
	model "proyecto/ArquiSoftware/model/order"
)

type orderDetailService struct{}

type orderDetailServiceInterface interface {
	InsertOrderDetail(dto.OrderDetailDto) (dto.OrderDetailDto, error)
	GetOrderDetailsByOrderId(int) (dto.OrderDetailsDto, error)
	GetOrderDetails() (dto.OrderDetailsDto, error)
}

var (
	OrderDetailService orderDetailServiceInterface
)

func init() {
	OrderDetailService = &orderDetailService{}
}

func (s *orderDetailService) InsertOrderDetail(orderDetailDto dto.OrderDetailDto) (dto.OrderDetailDto, error) {
	var orderDetail model.OrderDetail
	var orderDetailResponseDto dto.OrderDetailDto

	orderDetail.OrderID = orderDetailDto.OrderId
	orderDetail.ProductID = orderDetailDto.ProductId
	orderDetail.Price = orderDetailDto.Price
	orderDetail.Quantity = orderDetailDto.Quantity

	orderDetail, err := orderDetailCliente.InsertOrderDetail(orderDetail)

	if err != nil {
		log.Println(err)
		return orderDetailDto, err
	}

	orderDetailResponseDto.Id = orderDetail.OrderDetailID
	orderDetailResponseDto.OrderId = orderDetail.OrderID
	orderDetailResponseDto.ProductId = orderDetail.ProductID
	orderDetailResponseDto.Price = orderDetail.Price
	orderDetailResponseDto.Quantity = orderDetail.Quantity

	return orderDetailResponseDto, nil
}

func (s *orderDetailService) GetOrderDetailsByOrderId(id int) (dto.OrderDetailsDto, error) {
	var orderDetails model.OrderDetails
	var orderDetailsDto dto.OrderDetailsDto

	orderDetails, err := orderDetailCliente.GetOrderDetailsByOrderId(id)

	if err != nil {
		return orderDetailsDto, err
	}

	for _, orderDetail := range orderDetails {
		var orderDetailDto dto.OrderDetailDto
		orderDetailDto.Id = orderDetail.OrderDetailID
		orderDetailDto.OrderId = orderDetail.OrderID
		orderDetailDto.Price = orderDetail.Price
		orderDetailDto.ProductId = orderDetail.ProductID
		orderDetailDto.Quantity = orderDetail.Quantity
		orderDetailsDto = append(orderDetailsDto, orderDetailDto)
	}

	return orderDetailsDto, nil
}

func (s *orderDetailService) GetOrderDetails() (dto.OrderDetailsDto, error) {
	var orderDetailsDto dto.OrderDetailsDto

	orderDetails, err := orderDetailCliente.GetOrderDetails()

	if err != nil {
		return orderDetailsDto, err
	}

	for _, orderDetail := range orderDetails {
		var orderDetailDto dto.OrderDetailDto
		orderDetailDto.Id = orderDetail.OrderDetailID
		orderDetailDto.OrderId = orderDetail.OrderID
		orderDetailDto.Price = orderDetail.Price
		orderDetailDto.ProductId = orderDetail.ProductID
		orderDetailDto.Quantity = orderDetail.Quantity
		orderDetailsDto = append(orderDetailsDto, orderDetailDto)
	}

	return orderDetailsDto, nil
}
