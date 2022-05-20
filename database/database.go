package database

import (
	userClient "Proyecto/ArquiSoftware/clients/user"
	userModel "Proyecto/ArquiSoftware/model/user"

	productClient "Proyecto/ArquiSoftware/clients/product"
	productModel "Proyecto/ArquiSoftware/model/product"

	data "Proyecto/ArquiSoftware/database/data"

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

	data.InsertData(db)
}
