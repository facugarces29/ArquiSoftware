package services

import (
	homeCliente "github.com/facugarces29/ArquiSoftware/clients/home"
	dto "github.com/facugarces29/ArquiSoftware/dto/product"
)

type homeService struct{}

type homeServiceInterface interface {
	GetHomeProducts() (dto.ProductsDto, error)
}

var (
	HomeService homeServiceInterface
)

func init() {
	HomeService = &homeService{}
}

func (s *homeService) GetHomeProducts() (dto.ProductsDto, error) {

	products, err := homeCliente.GetHomeProducts()
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
