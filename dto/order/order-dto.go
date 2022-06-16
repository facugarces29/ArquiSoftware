package dto

type OrderDto struct {
	Id              uint `json:"id"`
	UserId          uint `json:"user_id"`
	OrderDetailsDto `json:"order_details"`
}

type OrdersDto []OrderDto
