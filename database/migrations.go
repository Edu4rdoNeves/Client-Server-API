package database

import (
	"github.com/Edu4rdoNeves/Client-Server-API/entity"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&entity.ExchangeRate{})
}
