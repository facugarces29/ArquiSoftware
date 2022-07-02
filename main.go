package main

import (
	"proyecto/ArquiSoftware/app"
	"proyecto/ArquiSoftware/database"
)

func main() {
	database.StartDbEngine()
	app.StartRoute()
}
