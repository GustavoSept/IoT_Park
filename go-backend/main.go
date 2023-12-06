package main

import (
	"go-backend/database"
	"go-backend/routes"
)

func main() {
	database.ConnectToDatabase()
	routes.HandleRequest()
}
