package services

import (
	orderCliente "proyecto/ArquiSoftware/clients/order"
	dto "proyecto/ArquiSoftware/dto/order"
	model "proyecto/ArquiSoftware/model/order"
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
		var amount float64
		orderDto.Id = order.ID
		orderDto.UserId = order.UserID

		orderDetails, _ := orderCliente.GetOrderDetailsByOrderId(int(orderDto.Id))
		var orderDetailsDto dto.OrderDetailsDto

		for _, orderDetail := range orderDetails {
			var orderDetailDto dto.OrderDetailDto
			orderDetailDto.Id = orderDetail.OrderDetailID
			orderDetailDto.OrderId = orderDetail.OrderID
			orderDetailDto.Price = orderDetail.Price
			orderDetailDto.ProductId = orderDetail.ProductID
			orderDetailDto.Quantity = orderDetail.Quantity
			orderDetailsDto = append(orderDetailsDto, orderDetailDto)

			amount = amount + (orderDetailDto.Price * float64(orderDetailDto.Quantity))
		}

		orderDto.OrderDetails = orderDetailsDto
		orderDto.Amount = amount

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

		orderDetails, _ := orderCliente.GetOrderDetailsByOrderId(int(orderDto.Id))
		var orderDetailsDto dto.OrderDetailsDto

		for _, orderDetail := range orderDetails {
			var orderDetailDto dto.OrderDetailDto
			orderDetailDto.Id = orderDetail.OrderDetailID
			orderDetailDto.OrderId = orderDetail.OrderID
			orderDetailDto.Price = orderDetail.Price
			orderDetailDto.ProductId = orderDetail.ProductID
			orderDetailDto.Quantity = orderDetail.Quantity
			orderDetailsDto = append(orderDetailsDto, orderDetailDto)
		}

		orderDto.OrderDetails = orderDetailsDto

		ordersDto = append(ordersDto, orderDto)
	}

	return ordersDto, nil
}
