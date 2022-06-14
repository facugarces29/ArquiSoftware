package services

import (
	orderCliente "github.com/facugarces29/ArquiSoftware/clients/order"
	dto "github.com/facugarces29/ArquiSoftware/dto/order"
	model "github.com/facugarces29/ArquiSoftware/model/order"
)

type orderService struct{}

type orderServiceInterface interface {
	InsertOrder(orderDto dto.InsertOrderDto) (dto.OrderDto, error)
	UpdateOrder(orderDto dto.UpdateOrderDto) (dto.OrderDto, error)
}

var (
	OrderService orderServiceInterface
)

func init() {
	OrderService = &orderService{}
}

func (s *orderService) InsertOrder(orderDto dto.InsertOrderDto) (dto.OrderDto, error) {

	var order model.Order
	var returnOrderDto dto.OrderDto

	//user.Name = userDto.Name

	order.UserID = orderDto.UserId

	order, err := orderCliente.InsertOrder(order)

	if err != nil {
		return returnOrderDto, err
	}

	returnOrderDto.Id = order.ID

	return returnOrderDto, nil
}

func (s *orderService) UpdateOrder(orderDto dto.UpdateOrderDto) (dto.OrderDto, error) {

	var order model.Order
	var returnOrderDto dto.OrderDto

	//user.Name = userDto.Name

	order.ID = orderDto.Id
	order.Amount = orderDto.Amount


	order, err := orderCliente.UpdateOrder(order)

	if err != nil {
		return returnOrderDto, err
	}

	returnOrderDto.Id = order.ID
	returnOrderDto.Amount = order.Amount

	return returnOrderDto, nil
}
