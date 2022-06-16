package services

import (
	homeCliente "github.com/facugarces29/ArquiSoftware/clients/home"
	productCliente "github.com/facugarces29/ArquiSoftware/clients/product"
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
