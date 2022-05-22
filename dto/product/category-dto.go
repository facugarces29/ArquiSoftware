package dto

type CategoryDto struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CategoriesDto []CategoryDto
