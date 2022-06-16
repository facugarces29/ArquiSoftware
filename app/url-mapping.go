package app

import (
	//ver
	homeController "github.com/facugarces29/ArquiSoftware/controllers/home"
	loginController "github.com/facugarces29/ArquiSoftware/controllers/login"
	orderController "github.com/facugarces29/ArquiSoftware/controllers/order"
	productController "github.com/facugarces29/ArquiSoftware/controllers/product"
	searchController "github.com/facugarces29/ArquiSoftware/controllers/search"
	userController "github.com/facugarces29/ArquiSoftware/controllers/user"
)

// MapUrls maps the urls
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

	// Order Mapping
	router.GET("/order/:userId", orderController.GetOrdersByUserId)
	router.POST("/order", orderController.InsertOrder)

}
