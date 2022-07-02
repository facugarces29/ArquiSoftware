package services

import (
	productCliente "proyecto/ArquiSoftware/clients/product"
	searchCliente "proyecto/ArquiSoftware/clients/search"
	dto "proyecto/ArquiSoftware/dto/product"
)

type searchService struct{}

type searchServiceInterface interface {
	GetProductsBySearchParam(string) (dto.ProductsDto, error)
}

var (
	SearchService searchServiceInterface
)

func init() {
	SearchService = &searchService{}
}

func (s *searchService) GetProductsBySearchParam(param string) (dto.ProductsDto, error) {

	products, err := searchCliente.GetProductsBySearchParam(param)

	var productsDto dto.ProductsDto

	if err != nil {
		return productsDto, err
	}

	for _, product := range products {
		var productDto dto.ProductDto

		category, err := productCliente.GetCategoryById(int(product.CategoryID))

		if err != nil {
			return productsDto, err
		}

		productDto.Id = product.ProductID
		productDto.Name = product.Name
		productDto.Description = product.Description
		productDto.Price = product.Price
		productDto.Stock = product.Stock
		productDto.Image = product.Image
		productDto.Category = category.Name

		productsDto = append(productsDto, productDto)
	}

	return productsDto, err
}
