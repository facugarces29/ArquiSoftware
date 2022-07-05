package database

import (
	userClient "proyecto/ArquiSoftware/clients/user"
	userModel "proyecto/ArquiSoftware/model/user"

	productClient "proyecto/ArquiSoftware/clients/product"
	productModel "proyecto/ArquiSoftware/model/product"

	addressClient "proyecto/ArquiSoftware/clients/address"
	addressModel "proyecto/ArquiSoftware/model/address"

	orderClient "proyecto/ArquiSoftware/clients/order"
	orderModel "proyecto/ArquiSoftware/model/order"

	searchClient "proyecto/ArquiSoftware/clients/search"

	homeClient "proyecto/ArquiSoftware/clients/home"

	data "proyecto/ArquiSoftware/database/data"

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
	DBName := "proyect"      //variable de entorno para nombre de la base de datos
	DBUser := "proyect_user" //variable de entorno para el usuario de la base de datos
	//DBPass := ""
	DBPass := "secret" //variable de entorno para la pass de la base de datos
	DBHost := "localhost"
	// ------------------------

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	db.LogMode(true)

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// We need to add all CLients that we build
	userClient.Db = db
	productClient.Db = db
	addressClient.Db = db
	orderClient.Db = db
	searchClient.Db = db
	homeClient.Db = db

}

func StartDbEngine() {

	// We need to migrate all classes model.
	db.AutoMigrate(&userModel.User{}, &addressModel.Address{}, &orderModel.Order{}, &orderModel.OrderDetail{}, &productModel.Product{}, &productModel.Category{})

	log.Info("Finishing Migration Database Tables")
	data.InsertData(db)
}
