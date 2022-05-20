package database

import (
	userClient "Proyecto/ArquiSoftware/clients/user"
	userModel "Proyecto/ArquiSoftware/model/user"

	productClient "Proyecto/ArquiSoftware/clients/product"
	productModel "Proyecto/ArquiSoftware/model/product"

	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// DB Connections Paramters
	DBName := os.Getenv("PROY_DB_NAME") //variable de entorno para nombre de la base de datos
	DBUser := os.Getenv("PROY_DB_USER") //variable de entorno para el usuario de la base de datos
	//DBPass := ""
	DBPass := os.Getenv("PROY_DB_PASS") //variable de entorno para la pass de la base de datos
	DBHost := "localhost"
	// ------------------------

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// We need to add all CLients that we build
	userClient.Db = db
	productClient.Db = db

}

func StartDbEngine() {

	// We need to migrate all classes model.
	db.AutoMigrate(&userModel.User{})

	db.AutoMigrate((&productModel.Product{}))
	log.Info("Finishing Migration Database Tables")

	// Create users
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

		/*Name        string  `gorm:"type:varchar(255);not null;unique"`
		Description string  `gorm:"type:varchar(255);not null"`
		Price       float64 `gorm:"type:decimal(10,2);not null"`
		Stock       int     `gorm:"type:integer(255);not null"`
		Image       string  `gorm:"type:blob;not null"`*/

		//manage errors...

		log.Info("Assets Created")
	} else {
		log.Info("Assets Were Already Created")
	}
}
