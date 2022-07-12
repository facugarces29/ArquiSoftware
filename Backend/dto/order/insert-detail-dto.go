package dto

type InsertDetailDto struct {
	OrderId   uint    `json:"order_id"`
	ProductId uint    `json:"product_id"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
}

type InsertDetailsDto []InsertDetailDto
