package main

import (
	"github.com/gulnurrrrrrrrrrrrrrrrr/Online-Book-Store-API-backend/config"
	"github.com/gulnurrrrrrrrrrrrrrrrr/Online-Book-Store-API-backend/routes"
	"log"
)

func main() {
	config.ConnectDatabase()

	r := routes.SetupRouter()

	log.Println("Online Book Store API is running on http://localhost:8080")
	log.Fatal(r.Run(":8080"))
}
