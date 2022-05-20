package app

import (
	productController "Proyecto/ArquiSoftware/controllers/product"
	userController "Proyecto/ArquiSoftware/controllers/user"
)

func MapUrls() {

	// Users Mapping
	router.GET("/user/:id", userController.GetUserById)
	router.GET("/user", userController.GetUsers)

	// Products Mapping
	router.GET("/product/:id", productController.GetProductById)
	router.GET("/product", productController.GetProducts)

}
