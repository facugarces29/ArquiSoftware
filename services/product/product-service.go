package services

import (
	productCliente "Proyecto/ArquiSoftware/clients/product"
	dto "Proyecto/ArquiSoftware/dto/product"
)

type productService struct{}

type productServiceInterface interface {
	GetProductById(id int) (dto.ProductDto, error)
	GetProducts() (dto.ProductsDto, error)
}

var (
	ProductService productServiceInterface
)

func init() {
	ProductService = &productService{}
}

func (s *productService) GetProductById(id int) (dto.ProductDto, error) {

	product, err := productCliente.GetProductById(id)
	var productDto dto.ProductDto

	if err != nil {
		return productDto, err
	}

	productDto.Id = product.ProductID
	productDto.Name = product.Name
	productDto.Description = product.Description
	productDto.Price = product.Price
	productDto.Stock = product.Stock
	productDto.Image = product.Image

	return productDto, err
}

func (s *productService) GetProducts() (dto.ProductsDto, error) {

	products, err := productCliente.GetProducts()
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
