package database

import (
	"go-bank-dki/models"
	"log"

	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{}, &models.Stock{}).Error
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
