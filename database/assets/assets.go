package database

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	userModel "Proyecto/ArquiSoftware/model/user"

	productModel "Proyecto/ArquiSoftware/model/product"
)

func InsertAssets(db *gorm.DB) {
	// Create assets
	log.Info("Creating Assets...")
	err := db.First(&userModel.User{}).Error

	if err != nil {
		//Creating users

		db.Create(&userModel.User{Name: "lautaro", LastName: "Saenz", UserName: "lautarose", Email: "abcdefg@gmail.com", Pwd: "hola123"})
		db.Create(&userModel.User{Name: "Joaco", LastName: "Reyero", UserName: "jaocoreyero", Email: "12345@gmail.com", Pwd: "hola123"})
		db.Create(&userModel.User{Name: "Facundo", LastName: "Garces", UserName: "Facuelcapo", Email: "asasas@gmail.com", Pwd: "hola123"})
		db.Create(&userModel.User{Name: "Hernan", LastName: "Lachampionliga", UserName: "hernanchampion", Email: "hernan@gmail.com", Pwd: "hola123"})
		db.Create(&userModel.User{Name: "Saul", LastName: "Hudson", UserName: "slash", Email: "slashGNR@gmail.com", Pwd: "hola123"})

		//Creating products

		db.Create(&productModel.Product{Name: "PC", Description: "pc gamer intel i7", Price: 100000, Stock: 5})

		//manage errors...

		log.Info("Assets Created")
	} else {
		log.Info("Assets Were Already Created")
	}
}
