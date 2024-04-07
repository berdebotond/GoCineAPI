package routes

import (
	"log"

	controller "interview_test_KenTech/pkg/controllers"

	middleware "interview_test_KenTech/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	// routes for user operations
	log.Println("User Routes")
	incomingRoutes.POST("users/signup", controller.Signup())
	incomingRoutes.POST("users/login", controller.Login())

}
func AuthRoutes(incomingRoutes *gin.Engine) {
	// routes for user operations
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/usersdata", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())

}

// FilmRoutes sets up the routes for film management
func FilmRoutes(incomingRoutes *gin.Engine) {
	log.Println("Film Routes")
	// Make sure to use the UserAuthenticate middleware to ensure only authenticated users can manage films
	incomingRoutes.Use(middleware.Authenticate())
	// Define routes for film operations
	incomingRoutes.POST("/films", controller.CreateFilm())            // Add a new film
	incomingRoutes.GET("/films", controller.GetFilms())               // List all films
	incomingRoutes.GET("/films/:film_id", controller.GetFilm())       // Get a single film by ID
	incomingRoutes.PUT("/films/:film_id", controller.UpdateFilm())    // Update a film
	incomingRoutes.DELETE("/films/:film_id", controller.DeleteFilm()) // Delete a film
}
