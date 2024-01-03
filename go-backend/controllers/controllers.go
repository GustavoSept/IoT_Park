package controllers

import (
	"fmt"
	"net/http"

	"go-backend/database"
	"go-backend/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllUsers(c *gin.Context) {
	var usersList []models.User
	err := database.DB.Select(&usersList, "SELECT * FROM users")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, usersList)
}

func GetAllParkingLots(c *gin.Context) {
	var pLotsWithOwners []struct {
		ID   uuid.UUID `db:"id" json:"id"`
		CEP  string    `db:"cep" json:"cep"`
		Name string    `db:"name" json:"name"`
	}

	query := `
    SELECT pl.id, pl.cep, u.first_name || ' ' || u.last_name AS name
    FROM parking_lots pl
    JOIN users u ON pl.owner_id = u.id
    `

	err := database.DB.Select(&pLotsWithOwners, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pLotsWithOwners)
}

func CreateUser(c *gin.Context) {
	var newUser models.User

	// Attempt to parse JSON request body into newUser model
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Use newUser to insert into db
	_, err := database.DB.NamedExec(
		`INSERT INTO users (first_name, last_name, office_level)
		VALUES (:first_name, :last_name, :office_level)`, &newUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newUser)
}

func CreateParkingLot(c *gin.Context) {
	var newParkingLot models.ParkingLot
	var owner struct {
		FirstName string `form:"owner_first_name"`
		LastName  string `form:"owner_last_name"`
	}

	if err := c.Bind(&newParkingLot); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Couldn't bind parkinglot": err.Error()})
		return
	}

	if err := c.Bind(&owner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Couldn't bind owner": err.Error()})
		return
	}

	// Check if user exists and is an owner
	var user models.User
	err := database.DB.Get(&user, `SELECT * FROM users WHERE first_name = $1 AND last_name = $2 LIMIT 1`, owner.FirstName, owner.LastName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	if user.Office_Level != "dono" {
		c.JSON(http.StatusForbidden, gin.H{"error": "User is not an owner"})
		return
	}

	// Generate UUID for parking lot
	newParkingLot.ID = uuid.New()

	fmt.Printf("\nNew Parking Lot:\n")
	fmt.Printf("ID:%v, AddrStreet:%v, AddrNumber:%v, CEP:%v\n",
		newParkingLot.ID, newParkingLot.AddrStreet, newParkingLot.AddrNumber, newParkingLot.CEP)

	fmt.Printf("\nUser (Owner):\n")
	fmt.Printf("ID:%v, First Name:%v, Last Name:%v\n",
		user.ID, user.First_Name, user.Last_Name)

	// Create parking lot
	newParkingLot.OwnerID = user.ID
	_, err = database.DB.NamedExec(
		`INSERT INTO parking_lots (id, addr_street, addr_number, cep, owner_id) `+
			`VALUES (:id, :addr_street, :addr_number, :cep, :owner_id)`, &newParkingLot)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Couldn't insert into parking_lots": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newParkingLot)
}

func UpdateUser(c *gin.Context) {
	// Logic to update an User
	c.JSON(http.StatusOK, gin.H{"message": "PUT request called"})
}

func DeleteUser(c *gin.Context) {
	// Logic to delete an User
	c.JSON(http.StatusOK, gin.H{"message": "DELETE request called"})
}
