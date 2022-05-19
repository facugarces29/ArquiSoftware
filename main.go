package main

import (
	"Proyecto/ArquiSoftware/app"
	"Proyecto/ArquiSoftware/database"
)

func main() {
	database.StartDbEngine()
	app.StartRoute()
}
