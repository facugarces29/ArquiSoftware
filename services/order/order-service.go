package services

import (
	orderCliente "github.com/facugarces29/ArquiSoftware/clients/order"
	dto "github.com/facugarces29/ArquiSoftware/dto/order"
	model "github.com/facugarces29/ArquiSoftware/model/order"
)

type orderService struct{}

type orderServiceInterface interface {
	InsertOrder(dto.InsertOrderDto) (dto.OrderDto, error)
	GetOrdersByUserId(int) (dto.OrdersDto, error)
	GetOrders() (dto.OrdersDto, error)
}

var (
	OrderService orderServiceInterface
)

func init() {
	OrderService = &orderService{}
}

func (s *orderService) InsertOrder(insertOrderDto dto.InsertOrderDto) (dto.OrderDto, error) {

	var order model.Order
	var orderDto dto.OrderDto

	order.UserID = insertOrderDto.UserId

	order, err := orderCliente.InsertOrder(order)

	if err != nil {
		return orderDto, err
	}

	orderDto.Id = order.ID
	orderDto.UserId = order.UserID

	return orderDto, nil
}

func (s *orderService) GetOrders() (dto.OrdersDto, error) {
	orders, err := orderCliente.GetOrders()

	var ordersDto dto.OrdersDto

	if err != nil {
		return ordersDto, err
	}

	for _, order := range orders {
		var orderDto dto.OrderDto
		orderDto.Id = order.ID
		orderDto.UserId = order.UserID
		for _, orderDetail := range order.OrderDetails {
			var orderDetailDto dto.OrderDetailDto
			orderDetailDto.Id = orderDetail.OrderDetailID
			orderDetailDto.OrderId = orderDetail.OrderID
			orderDetailDto.Price = orderDetail.Price
			orderDetailDto.ProductId = orderDetail.ProductID
			orderDetailDto.Quantity = orderDetail.Quantity

			orderDto.OrderDetailsDto = append(orderDto.OrderDetailsDto, orderDetailDto)
		}
		ordersDto = append(ordersDto, orderDto)
	}

	return ordersDto, nil
}

func (s *orderService) GetOrdersByUserId(idUser int) (dto.OrdersDto, error) {
	orders, err := orderCliente.GetOrdersByUserId(idUser)

	var ordersDto dto.OrdersDto

	if err != nil {
		return ordersDto, err
	}

	for _, order := range orders {
		var orderDto dto.OrderDto
		orderDto.Id = order.ID
		orderDto.UserId = order.UserID
		for _, orderDetail := range order.OrderDetails {
			var orderDetailDto dto.OrderDetailDto
			orderDetailDto.Id = orderDetail.OrderDetailID
			orderDetailDto.OrderId = orderDetail.OrderID
			orderDetailDto.Price = orderDetail.Price
			orderDetailDto.ProductId = orderDetail.ProductID
			orderDetailDto.Quantity = orderDetail.Quantity

			orderDto.OrderDetailsDto = append(orderDto.OrderDetailsDto, orderDetailDto)
		}
		ordersDto = append(ordersDto, orderDto)
	}

	return ordersDto, nil
}
