package app

import (
	//ver
	addressController "proyecto/ArquiSoftware/controllers/address"
	homeController "proyecto/ArquiSoftware/controllers/home"
	loginController "proyecto/ArquiSoftware/controllers/login"
	orderController "proyecto/ArquiSoftware/controllers/order"
	productController "proyecto/ArquiSoftware/controllers/product"
	searchController "proyecto/ArquiSoftware/controllers/search"
	userController "proyecto/ArquiSoftware/controllers/user"
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
	router.GET("/", homeController.GetHomeProducts)

	// Order Mapping
	router.GET("/order", orderController.GetOrders)
	router.GET("/order/:userId", orderController.GetOrdersByUserId)
	router.POST("/order", orderController.InsertOrder)
	router.DELETE("/order/:orderId", orderController.DeleteOrder)

	// Order Detail Mapping
	router.POST("/order/detail", orderController.InsertOrderDetail)
	router.GET("/order/detail", orderController.GetOrderDetails)
	router.GET("/order/detail/:orderId", orderController.GetOrderDetailsByOrderId)

	// Address Mapping
	router.GET("/address/:id", addressController.GetAddressByUserId)
}
