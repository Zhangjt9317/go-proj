package main

import (
	"polling-site/routes"
	"polling-site/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	database.InitDB()

	// Create a new Gin router
	router := gin.Default()

	// Set up routes
	routes.SetupRoutes(router)

	// Start the server on port 8080
	router.Run(":8080")
}