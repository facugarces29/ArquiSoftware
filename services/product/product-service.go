package services

import (
	productCliente "Proyecto/ArquiSoftware/clients/product"
	dto "Proyecto/ArquiSoftware/dto/product"
	model "Proyecto/ArquiSoftware/model/product"
)

type productService struct{}

type productServiceInterface interface {
	GetProductById(id int) dto.ProductDto
	GetProducts() dto.ProductsDto
}

var (
	ProductService productServiceInterface
)

func init() {
	ProductService = &productService{}
}

func (s *productService) GetProductById(id int) dto.ProductDto {

	var product model.Product = productCliente.GetProductById(id)
	var productDto dto.ProductDto

	if product.ProductID == 0 {
		return productDto
	}

	productDto.Id = product.ProductID
	productDto.Name = product.Name
	productDto.Description = product.Description
	productDto.Price = product.Price
	productDto.Stock = product.Stock
	productDto.Image = product.Image

	return productDto
}

func (s *productService) GetProducts() dto.ProductsDto {
	var products model.Products = productCliente.GetProducts()
	var productsDto dto.ProductsDto

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

	return productsDto
}
