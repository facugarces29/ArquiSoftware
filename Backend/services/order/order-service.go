package services

import (
	"log"
	orderCliente "proyecto/ArquiSoftware/clients/order"
	productCliente "proyecto/ArquiSoftware/clients/product"
	dto "proyecto/ArquiSoftware/dto/order"
	model "proyecto/ArquiSoftware/model/order"
)

type orderService struct{}

type orderServiceInterface interface {
	InsertOrder(dto.InsertOrderDto) (dto.OrderDto, error)
	GetOrdersByUserId(int) (dto.OrdersByIdResponseDto, error)
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
		orderDto.Date = order.CreatedAt.Format("2006-January-02")

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

func (s *orderService) GetOrdersByUserId(idUser int) (dto.OrdersByIdResponseDto, error) {
	orders, err := orderCliente.GetOrdersByUserId(idUser)

	var ordersResponseDto dto.OrdersByIdResponseDto

	if err != nil {
		return ordersResponseDto, err
	}

	for _, order := range orders {
		var orderResponseDto dto.OrderByIdResponseDto
		orderResponseDto.Id = order.ID
		orderResponseDto.UserId = order.UserID
		var amount float64
		orderResponseDto.Date = order.CreatedAt.Format("2006-January-02")

		orderDetails, _ := orderCliente.GetOrderDetailsByOrderId(int(orderResponseDto.Id))
		var orderProductsDto dto.OrderProductsDto

		for _, orderDetail := range orderDetails {
			var orderProductDto dto.OrderProductDto

			orderProduct, err := productCliente.GetProductById(int(orderDetail.ProductID))

			if err != nil {
				log.Println(err)
				return ordersResponseDto, err
			}

			orderProductDto.Id = orderDetail.ProductID
			orderProductDto.Name = orderProduct.Name
			orderProductDto.Price = orderDetail.Price
			orderProductDto.Quantity = orderDetail.Quantity
			orderProductDto.Image = orderProduct.Image

			orderProductsDto = append(orderProductsDto, orderProductDto)

			amount = amount + (orderProductDto.Price * float64(orderProductDto.Quantity))
		}

		orderResponseDto.OrderProducts = orderProductsDto
		orderResponseDto.Amount = amount
		orderResponseDto.Image = orderProductsDto[0].Image

		ordersResponseDto = append(ordersResponseDto, orderResponseDto)
	}

	return ordersResponseDto, nil
}
