package database

import(
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"main.go/models"
	"log"
)

var DB *gorm.DB

func Connect(){
	
	dsn := "root:Neesha@98@tcp(127.0.0.1:3306)/shopdb?charset=utf8mb4&parseTime=True&loc=Local"

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Invalid database url")
	}
	DB = connection

	connection.AutoMigrate(&models.Employee{})



}