package database

import (
	"auth/src/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Represents the database connection
var Conn *gorm.DB

var err error

var model_list = []interface{}{
	&models.User{},
}

// ConnectDB initializes the database connection using GORM
func ConnectDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	Conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}
}

// Migrate creates the database tables based on the models
func Migrate() {
	err := Conn.AutoMigrate(model_list...)
	if err != nil {
		log.Fatal("Failed to migrate database")
	}
}
