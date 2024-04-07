package middleware

import (
	"log"
	"net/http"

	helper "interview_test_KenTech/pkg/helper"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Authenticating")
		clientToken := c.Request.Header.Get("Authorization") // Use the standard Authorization header

		// Check if the token was provided
		if clientToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No Authorization header provided"})
			c.Abort()
			return
		}

		// Assuming the token comes with a 'Bearer ' prefix
		if len(clientToken) > 7 && clientToken[:7] == "Bearer " {
			clientToken = clientToken[7:]
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		claims, err := helper.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Set user details in the Gin context for use in subsequent handlers
		c.Set("username", claims.Username)
		c.Set("uid", claims.Uid)
		c.Set("user_type", claims.User_type)

		c.Next() // Proceed to the next handler if authentication is successful
	}
}
