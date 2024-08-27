package database

import (
	"go-bank-dki/models"
	"log"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func SeedUsers(db *gorm.DB) {
	var count int64
	db.Model(&models.User{}).Count(&count)

	if count == 0 {
		passwordHash, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Failed to hash password: %v", err)
		}

		user := models.User{
			Username: "admin",
			Password: string(passwordHash),
		}

		if err := db.Create(&user).Error; err != nil {
			log.Fatalf("Failed to seed users: %v", err)
		} else {
			log.Println("Seeded default user: admin")
		}
	} else {
		log.Println("Users table already seeded")
	}
}
