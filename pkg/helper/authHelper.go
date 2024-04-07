package helper

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// CheckUserType checks if the user has the required role to access a resource.
func CheckUserType(c *gin.Context, role string) error {
	userType := c.GetString("user_type")

	if userType != role {
		return errors.New("you are not authorised to access this resource")
	}
	return nil
}

// MatchUserTypeToUid checks if a "USER" has the same ID as the userId parameter or if the userType matches a specific role.
func MatchUserTypeToUid(c *gin.Context, userId string) error {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")

	if userType == "USER" && uid != userId {
		return errors.New("you are not authorised to access this resource")
	}
	// Assuming you want to validate that userType has a specific role; if not, this can be removed.
	if _, ok := validRoles[userType]; !ok {
		return errors.New("invalid user type")
	}
	return nil
}

// Assuming there's a list of valid roles; this helps ensure roles are checked against defined roles.
var validRoles = map[string]bool{
	"USER":  true,
	"ADMIN": true,
	// Add other valid roles as needed.
}
