package main

import (
	"github.com/facugarces29/ArquiSoftware/app"
	"github.com/facugarces29/ArquiSoftware/database"
)

func main() {
	database.StartDbEngine()
	app.StartRoute()
}
