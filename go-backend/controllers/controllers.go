package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handlers for the endpoints
func GetItems(c *gin.Context) {
	// Logic to retrieve items
	// Respond with JSON
	c.JSON(http.StatusOK, gin.H{"message": "GET request called"})
}

func CreateItem(c *gin.Context) {
	// Logic to create a new item
	c.JSON(http.StatusOK, gin.H{"message": "POST request called"})
}

func UpdateItem(c *gin.Context) {
	// Logic to update an item
	c.JSON(http.StatusOK, gin.H{"message": "PUT request called"})
}

func DeleteItem(c *gin.Context) {
	// Logic to delete an item
	c.JSON(http.StatusOK, gin.H{"message": "DELETE request called"})
}
