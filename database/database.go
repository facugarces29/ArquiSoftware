package database

import (
	userClient "Proyecto/ArquiSoftware/clients"
	model "Proyecto/ArquiSoftware/model/user"
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
}

func StartDbEngine() {

	// We need to migrate all classes model.
	db.AutoMigrate(&model.User{})

	// Create users
	err := db.First(&model.User{}).Error

	if err != nil {
		log.Info("Database already established")
	} else {
		db.Create(&model.User{Name: "lautaro", LastName: "Saenz", UserName: "lautarose", Email: "abcdefg@gmail.com", Pwd: "hola123"})
		db.Create(&model.User{Name: "Joaco", LastName: "Reyero", UserName: "jaocoreyero", Email: "12345@gmail.com", Pwd: "hola123"})
		db.Create(&model.User{Name: "facundo", LastName: "Garces", UserName: "Facuelcapo", Email: "asasas@gmail.com", Pwd: "hola123"})
		db.Create(&model.User{Name: "Hernan", LastName: "Lachampionliga", UserName: "hernanchampion", Email: "hernan@gmail.com", Pwd: "hola123"})
	}

	log.Info("Finishing Migration Database Tables")
}
