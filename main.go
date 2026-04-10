package main

import (
	"log"

	"github.com/gulnurrrrrrrrrrrrrrrrr/Online-Book-Store-API-backend/config"
	"github.com/gulnurrrrrrrrrrrrrrrrr/Online-Book-Store-API-backend/routes"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment")
	}
	config.ConnectDatabase()

	r := routes.SetupRouter()

	log.Println("Online Book Store API is running on http://localhost:8080")
	log.Fatal(r.Run(":8080"))
}
