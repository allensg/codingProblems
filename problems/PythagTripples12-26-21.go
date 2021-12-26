package problems

import (
	"github.com/gin-gonic/gin"
)

// DeleteUserGame links a UserMedia object to a User
func PythagTripples(c *gin.Context) {
	// func (h *Handler) PythagTripples(c *gin.Context) {
	// session := sessions.Default(c)
	// userID := session.Get("userID")
	// userString := fmt.Sprintf("%v", userID)

	// c.HTML(http.StatusOK, "Hello, Docker! <3lololol", {})
	c.JSON(200, gin.H{"Success": "Game removed from library"})
	return
}
