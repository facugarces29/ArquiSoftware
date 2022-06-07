package database

import (
	userClient "github.com/facugarces29/ArquiSoftware/clients/user"
	userModel "github.com/facugarces29/ArquiSoftware/model/user"

	productClient "github.com/facugarces29/ArquiSoftware/clients/product"
	productModel "github.com/facugarces29/ArquiSoftware/model/product"

	addressClient "github.com/facugarces29/ArquiSoftware/clients/address"
	addressModel "github.com/facugarces29/ArquiSoftware/model/address"

	orderClient "github.com/facugarces29/ArquiSoftware/clients/order"
	orderModel "github.com/facugarces29/ArquiSoftware/model/order"

	searchClient "github.com/facugarces29/ArquiSoftware/clients/search"

	data "github.com/facugarces29/ArquiSoftware/database/data"

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

}

func StartDbEngine() {

	// We need to migrate all classes model.
	db.AutoMigrate(&userModel.User{}, &addressModel.Address{}, &orderModel.Order{}, &orderModel.OrderDetail{}, &productModel.Product{}, &productModel.Category{})

	log.Info("Finishing Migration Database Tables")
	data.InsertData(db)
}
