package services

import (
	"log"

	orderDetailCliente "github.com/facugarces29/ArquiSoftware/clients/order"
	dto "github.com/facugarces29/ArquiSoftware/dto/order"
	model "github.com/facugarces29/ArquiSoftware/model/order"
)

type orderDetailService struct{}

type orderDetailServiceInterface interface {
	InsertOrderDetail(orderDetailDto dto.OrderDetailDto) (dto.OrderDetailDto, error)
}

var (
	OrderDetailService orderDetailServiceInterface
)

func init() {
	OrderDetailService = &orderDetailService{}
}

func (s *orderDetailService) InsertOrderDetail(orderDetailDto dto.OrderDetailDto) (dto.OrderDetailDto, error) {
	var orderDetail model.OrderDetail

	orderDetail.OrderID = orderDetailDto.OrderId
	orderDetail.ProductID = orderDetailDto.ProductId
	orderDetail.Price = orderDetailDto.Price
	orderDetail.Quantity = orderDetailDto.Quantity

	orderDetail, err := orderDetailCliente.InsertOrderDetail(orderDetail)

	if err != nil {
		log.Println(err)
		return orderDetailDto, err
	}

	return orderDetailDto, nil
}
