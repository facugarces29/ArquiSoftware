package services

import (
	categoryCliente "Proyecto/ArquiSoftware/clients/product"
	dto "Proyecto/ArquiSoftware/dto/product"
	model "Proyecto/ArquiSoftware/model/product"
)

type categoryService struct{}

type categoryServiceInterface interface {
	GetCategoryById(id int) dto.CategoryDto
	GetCategories() dto.CategoriesDto
}

var (
	CategoryService categoryServiceInterface
)

func init() {
	CategoryService = &categoryService{}
}

func (s *categoryService) GetCategoryById(id int) dto.CategoryDto {

	var category model.Category = categoryCliente.GetCategoryById(id)
	var categoryDto dto.CategoryDto

	if category.CategoryID == 0 {
		return categoryDto
	}

	categoryDto.Id = category.CategoryID
	categoryDto.Name = category.Name

	return categoryDto
}

func (s *categoryService) GetCategories() dto.CategoriesDto {
	var categories model.Categories = categoryCliente.GetCategories()
	var categoriesDto dto.CategoriesDto

	for _, category := range categories {
		var categoryDto dto.CategoryDto
		categoryDto.Id = category.CategoryID
		categoryDto.Name = category.Name

		categoriesDto = append(categoriesDto, categoryDto)
	}

	return categoriesDto
}
