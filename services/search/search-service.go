package services

import (
	productCliente "github.com/facugarces29/ArquiSoftware/clients/search"
	dto "github.com/facugarces29/ArquiSoftware/dto/product"
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

	products, err := productCliente.GetProductsBySearchParam(param)

	var productsDto dto.ProductsDto

	if err != nil {
		return productsDto, err
	}

	for _, product := range products {
		var productDto dto.ProductDto
		productDto.Id = product.ProductID
		productDto.Name = product.Name
		productDto.Description = product.Description
		productDto.Price = product.Price
		productDto.Stock = product.Stock
		productDto.Image = product.Image

		productsDto = append(productsDto, productDto)
	}

	return productsDto, err
}
