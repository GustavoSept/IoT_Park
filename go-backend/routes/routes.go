package routes

import (
	"go-backend/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	router := gin.Default()

	// Define routes
	router.GET("/items", controllers.GetItems)
	router.POST("/items", controllers.CreateItem)
	router.PUT("/items/:id", controllers.UpdateItem)
	router.DELETE("/items/:id", controllers.DeleteItem)

	// Run the server
	router.Run(":8080")
}
