package database

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB

func StartDataBaseConnect() {
	dsn := os.Getenv("DSN")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Fail to connect database: ", err)
	}

	database = db

	RunMigrations(db)
}

func GetDatabase() *gorm.DB {
	return database
}
