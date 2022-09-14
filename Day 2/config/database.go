package config

import (
	"fmt"
	"latihan/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB

func migration() {
	Instance.AutoMigrate(&models.User{}, &models.Book{})
}

func ConnectDataBase() *gorm.DB {
	username := os.Getenv("DBUSERNAME")
	password := os.Getenv("DBPASS")
	host := fmt.Sprintf("tcp(%v:%v)", os.Getenv("DBHOST"), os.Getenv("DBPORT"))
	database := os.Getenv("DBNAME")

	dsn := fmt.Sprintf("%v:%v@%v/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	Instance = db
	if err != nil {
		panic(err.Error())
	}
	migration()
	return db
}
