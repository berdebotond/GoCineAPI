package main

import (
	"interview_test_KenTech/config"
	routes "interview_test_KenTech/pkg/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the router
	router := gin.New()

	// Define routes for different modules
	routes.UserRoutes(router)
	routes.AuthRoutes(router)
	routes.FilmRoutes(router)

	// Start the server and log the port being used
	log.Printf("Server started on port %s", config.CFG.Port)

	// Starting the server with error handling
	err := router.Run(":" + config.CFG.Port)
	if err != nil {
		// Log fatal error: this type of error usually indicates an unrecoverable state
		// and typically warrants a program shutdown.
		log.Fatalf("Failed to start server: %v", err)
	}
}
