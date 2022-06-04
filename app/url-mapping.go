package app

import (
	loginController "Proyecto/ArquiSoftware/controllers/login"
	productController "Proyecto/ArquiSoftware/controllers/product"
	searchController "Proyecto/ArquiSoftware/controllers/search"
	userController "Proyecto/ArquiSoftware/controllers/user"
)

func MapUrls() {
	// Users Mapping
	router.GET("/user/:id", userController.GetUserById)
	router.GET("/user", userController.GetUsers)

	// Login Mapping
	router.POST("/login", loginController.Login)

	// Products Mapping
	router.GET("/product/:id", productController.GetProductById)
	router.GET("/product", productController.GetProducts)

	// Categories Mapping
	router.GET("/category/:id", productController.GetCategoryById)
	router.GET("/category", productController.GetCategories)

	// Search Mapping
	router.GET("/search/:param", searchController.GetProductsBySearchParam)
}
