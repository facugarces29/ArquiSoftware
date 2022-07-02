package services

import (
	categoryCliente "proyecto/ArquiSoftware/clients/product"
	dto "proyecto/ArquiSoftware/dto/product"
)

type categoryService struct{}

type categoryServiceInterface interface {
	GetCategoryById(id int) (dto.CategoryDto, error)
	GetCategories() (dto.CategoriesDto, error)
}

var (
	CategoryService categoryServiceInterface
)

func init() {
	CategoryService = &categoryService{}
}

func (s *categoryService) GetCategoryById(id int) (dto.CategoryDto, error) {

	category, err := categoryCliente.GetCategoryById(id)

	var categoryDto dto.CategoryDto

	if err != nil {
		return categoryDto, err
	}

	if category.CategoryID == 0 {
		return categoryDto, nil
	}

	categoryDto.Id = category.CategoryID
	categoryDto.Name = category.Name

	return categoryDto, nil
}

func (s *categoryService) GetCategories() (dto.CategoriesDto, error) {
	categories, err := categoryCliente.GetCategories()
	var categoriesDto dto.CategoriesDto

	if err != nil {
		return categoriesDto, err
	}

	for _, category := range categories {
		var categoryDto dto.CategoryDto
		categoryDto.Id = category.CategoryID
		categoryDto.Name = category.Name

		categoriesDto = append(categoriesDto, categoryDto)
	}

	return categoriesDto, nil
}
