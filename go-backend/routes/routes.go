package routes

import (
	"go-backend/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	router := gin.Default()

	// Define routes
	router.GET("/users", controllers.GetUsers)
	// router.POST("/users", controllers.CreateUser)
	// router.PUT("/users/:id", controllers.UpdateUser)
	// router.DELETE("/users/:id", controllers.DeleteUser)

	// Run the server
	router.Run(":8080")
}
