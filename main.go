package main

import (
	database "auth/config"
	"auth/src/helpers"
	routes "auth/src/routers"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database connection
	database.ConnectDB()
	database.Migrate()

	// Set up routes
	router := routes.SetupRoutes()
	helpers.LogRoutes(router)

	// Initialize server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
