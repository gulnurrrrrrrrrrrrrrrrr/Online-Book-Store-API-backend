package config

import (
	"fmt"
	"github.com/gulnurrrrrrrrrrrrrrrrr/Online-Book-Store-API-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=7983 dbname=bookstore port=5432 sslmode=disable"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	err = DB.AutoMigrate(&models.Book{}, &models.Author{}, &models.Category{})
	if err != nil {
		log.Fatal("Migration error:", err)
	}

	fmt.Println("✅ Connected to PostgreSQL with GORM")
}
