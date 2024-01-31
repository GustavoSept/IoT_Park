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
		AllowOrigins: []string{
			"http://0.0.0.0:5000", "http://localhost:5000", "http://flask-frontend:5000",
			"http://0.0.0.0:5000/*", "http://localhost:5000/*", "http://flask-frontend:5000/*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "X-Requested-With", "Accept", "HX-Request", "HX-Current-URL", "HX-Target"},
		AllowCredentials: true,
	}))

	// Define routes
	authorized := router.Group("/api")
	authorized.Use(middleware.AuthRequired)
	{
		authorized.GET("/get_all_users", controllers.GetAllUsers)
		authorized.GET("/get_all_pLots", controllers.GetAllParkingLots)
		authorized.POST("/logout", controllers.LogUserOut)

		onlyOwners := router.Group("/")
		onlyOwners.Use(middleware.OnlyOwners)
		{
			onlyOwners.GET("/get_all_usersAuth", controllers.GetAllUsersAuth)
			onlyOwners.POST("/create-user", middleware.GetParkingLotContext, controllers.CreateUser)
			onlyOwners.POST("/create-parking-lot", controllers.CreateParkingLot)
			onlyOwners.DELETE("/delete-user", controllers.DeleteUser)
		}

	}

	router.POST("/api/login", controllers.LoginUser)

	// Run the server
	router.Run(":8080")
}
