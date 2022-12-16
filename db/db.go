package db

import (
	"fmt"
	"golang-echo/config"
	"golang-echo/entities"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Connect() {
	appConfig := config.GetConfig()
	connection_string := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", appConfig.DB_USERNAME,
		appConfig.DB_PASSWORD,
		appConfig.DB_HOST,
		appConfig.DB_PORT,
		appConfig.DB_NAME)

	db, err = gorm.Open(mysql.Open(connection_string), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to database")
	}
	log.Println("Connecting to database")
	Migrate()
}

func Migrate() {
	db.AutoMigrate(&entities.Product{})
	log.Println("Database migration completed")
}

func DbManager() *gorm.DB {
	return db
}
