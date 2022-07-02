package dto

type InsertOrderDto struct {
	UserId uint `json:"user_id"`
}

type InsertOrdersDto []InsertOrderDto
