package routes

import (
	"go-backend/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	router := gin.Default()

	// Define routes
	router.GET("/api/get_all_users", controllers.GetAllUsers)
	router.GET("/api/parkinglot=:id", controllers.GetAllUsers)
	// router.POST("/users", controllers.CreateUser)
	// router.PUT("/users/:id", controllers.UpdateUser)
	// router.DELETE("/users/:id", controllers.DeleteUser)

	// Run the server
	router.Run(":8080")
}
