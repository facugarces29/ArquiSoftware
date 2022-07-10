package dto

type OrderProductDto struct {
	Id       uint    `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Image    string  `json:"image"`
}

type OrderProductsDto []OrderProductDto
