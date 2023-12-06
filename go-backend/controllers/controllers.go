package controllers

import (
	"net/http"

	"go-backend/database"
	"go-backend/models"

	"github.com/gin-gonic/gin"
)

// Handlers for the endpoints
func GetUsers(c *gin.Context) {
	var usersList []models.User
	err := database.DB.Select(&usersList, "SELECT * FROM users")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, usersList)
}

func CreateUser(c *gin.Context) {
	// Logic to create a new User
	c.JSON(http.StatusOK, gin.H{"message": "POST request called"})
}

func UpdateUser(c *gin.Context) {
	// Logic to update an User
	c.JSON(http.StatusOK, gin.H{"message": "PUT request called"})
}

func DeleteUser(c *gin.Context) {
	// Logic to delete an User
	c.JSON(http.StatusOK, gin.H{"message": "DELETE request called"})
}
