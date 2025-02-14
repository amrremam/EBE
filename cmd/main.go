package main

import (
	"log"
	"github.com/amrremam/EBE.git/cmd/api"
	"github.com/amrremam/EBE.git/config"
	"github.com/amrremam/EBE.git/models"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Connect to database
	if err := config.ConnectDatabase(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected successfully")

	// Auto migrate models
	config.DB.AutoMigrate(&models.User{}, &models.Task{})
	log.Println("Database migrated successfully")

	// Start server

	router := api.Routes()

	log.Println("Server is running")
	router.Run(":8080")
}

