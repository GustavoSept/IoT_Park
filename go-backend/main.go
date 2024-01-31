package main

import (
	"go-backend/database"
	"go-backend/helpers"
	"go-backend/routes"
	"log"
)

func main() {
	jwtErr := helpers.InitJWT()
	if jwtErr != nil {
		log.Println("Error initializing the JWT's!")
		log.Fatal(jwtErr)
	}

	database.ConnectToDatabase()
	routes.HandleRequest()
}
