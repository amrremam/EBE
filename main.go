package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/amrremam/EBE.git/config"
	"github.com/amrremam/EBE.git/models"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Connect to the database
	config.ConnectDatabase()

	// Auto-migrate models
	config.DB.AutoMigrate(&models.User{}, &models.Task{})

	log.Println("Database migrated successfully!")

	// Start server
	r := gin.Default()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server is running on port", port)
	r.Run(":" + port)
}
