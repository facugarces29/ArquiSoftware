package dto

type UpdateOrderDto struct {
	Id              uint    `json:"id"`
	Amount          float64 `json:"amount"`
}
