package dto

type OrderByIdResponseDto struct {
	Id            uint             `json:"id"`
	UserId        uint             `json:"user_id"`
	Amount        float64          `json:"amount"`
	Date          string           `json:"date"`
	Image         string           `json:"image"`
	OrderProducts OrderProductsDto `json:"order_products"`
}

type OrdersByIdResponseDto []OrderByIdResponseDto
