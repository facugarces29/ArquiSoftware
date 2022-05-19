package app

import (
	userController "Proyecto/ArquiSoftware/controllers/user"
)

func MapUrls() {

	// Users Mapping
	router.GET("/user/:id", userController.GetUserById)
	router.GET("/user", userController.GetUsers)

}
