package config

import (
	"fmt"
	"inventory-api/internal/model"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Println("Warning: File .env not found, using system environment variable")
	}
}

func ConnectDB(){
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	DB = database
	log.Println("Successfully connected to PostgreSQL database!")

	
	err = DB.AutoMigrate(&model.User{}, &model.Category{}, &model.Supplier{}, &model.Product{}, &model.StockMovement{})
	if err != nil {
		log.Fatal("Failed to auto migrate: ", err)
	}

	log.Println("Auto Migrate Complete")
}