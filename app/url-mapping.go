package app

import (
	homeController "github.com/facugarces29/ArquiSoftware/controllers/home"
	loginController "github.com/facugarces29/ArquiSoftware/controllers/login"
	productController "github.com/facugarces29/ArquiSoftware/controllers/product"
	searchController "github.com/facugarces29/ArquiSoftware/controllers/search"
	userController "github.com/facugarces29/ArquiSoftware/controllers/user"
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

	// Home Mapping
	router.GET("/home", homeController.GetHomeProducts)
}
