package main

import (
	"Mereb-V2/config"
	"Mereb-V2/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Warning: Could not load .env file")
	}

	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "5000"
	}

	// Create a new Gin router
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(config.CorsConfig)

	routes.PersonRoutes(router)

	log.Printf("Server starting on port %s", Port)
	router.Run(":" + Port)
}
