package dto

type OrderDto struct {
	Id           uint            `json:"id"`
	UserId       uint            `json:"user_id"`
	Amount       float64         `json:"amount"`
	Date         string          `json:"date"`
	OrderDetails OrderDetailsDto `json:"order_details"`
}

type OrdersDto []OrderDto
