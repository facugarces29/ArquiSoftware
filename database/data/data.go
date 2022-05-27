package database

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	addressModel "Proyecto/ArquiSoftware/model/address"
	userModel "Proyecto/ArquiSoftware/model/user"
)

func InsertData(db *gorm.DB) {
	// Insert data
	log.Info("Inserting data...")

	//Creating users
	err := db.First(&userModel.User{}).Error

	if err != nil {
		db.Create(&userModel.User{Name: "lautaro", LastName: "Saenz", UserName: "lautarose", Email: "abcdefg@gmail.com", Pwd: "hola123"})
		db.Create(&userModel.User{Name: "Joaco", LastName: "Reyero", UserName: "jaocoreyero", Email: "12345@gmail.com", Pwd: "hola123"})
		db.Create(&userModel.User{Name: "Facundo", LastName: "Garces", UserName: "Facuelcapo", Email: "asasas@gmail.com", Pwd: "hola123"})
		db.Create(&userModel.User{Name: "Hernan", LastName: "Lachampionliga", UserName: "hernanchampion", Email: "hernan@gmail.com", Pwd: "hola123"})
		db.Create(&userModel.User{Name: "Saul", LastName: "Hudson", UserName: "slash", Email: "slashGNR@gmail.com", Pwd: "hola123"})
	}

	//Creating products

	//Creating addresses

	err = db.First(&addressModel.Address{}).Error

	if err != nil {
		db.Create(&addressModel.Address{UserID: 1, State: "Cordoba", City: "Cordoba", Zip: 5000, Addressline: "Eugenio Garzon 467"})
		db.Create(&addressModel.Address{UserID: 2, State: "Cordoba", City: "Cordoba", Zip: 5000, Addressline: "Ituzaingo 200"})
		db.Create(&addressModel.Address{UserID: 3, State: "Cordoba", City: "Cordoba", Zip: 5000, Addressline: "General Paz 344"})
		db.Create(&addressModel.Address{UserID: 4, State: "Cordoba", City: "Cordoba", Zip: 5000, Addressline: "Velez Sarsfield 1500"})
		db.Create(&addressModel.Address{UserID: 5, State: "Cordoba", City: "Cordoba", Zip: 5000, Addressline: "Belgrano 110"})
	}

	//db.Create(&productModel.Product{Name: "PC", Description: "pc gamer intel i7", Price: 100000, Stock: 5})

	//manage errors...

	log.Info("Data inserted")
}
