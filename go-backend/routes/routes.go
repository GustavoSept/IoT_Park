package routes

import (
	"go-backend/controllers"
	"go-backend/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://0.0.0.0:5000", "http://localhost:5000", "http://flask-frontend:5000"},
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "X-Requested-With", "Accept", "HX-Request", "HX-Current-URL", "HX-Target"},
		AllowCredentials: true,
	}))

	// Define routes
	router.GET("/api/get_all_users", controllers.GetAllUsers)
	router.GET("/api/get_all_usersAuth", controllers.GetAllUsersAuth)
	router.GET("/api/get_all_pLots", controllers.GetAllParkingLots)
	router.POST("/api/create-user", middleware.GetParkingLotContext, controllers.CreateUser)
	router.POST("/api/create-parking-lot", controllers.CreateParkingLot)
	// router.PUT("/users/:id", controllers.UpdateUser)
	// router.DELETE("/users/:id", controllers.DeleteUser)

	// Run the server
	router.Run(":8080")
}
