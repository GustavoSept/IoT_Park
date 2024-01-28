package controllers

import (
	"net/http"

	"go-backend/database"
	"go-backend/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllParkingLots(c *gin.Context) {
	var pLotsWithOwners []struct {
		ID        uuid.UUID `db:"id" json:"id"`
		PLot_Name string    `db:"pl_name" json:"pl_name"`
		CEP       string    `db:"cep" json:"cep"`
		OwnerName string    `db:"name" json:"name"`
	}

	query := `
    SELECT pl.id, pl.pl_name, pl.cep, u.first_name || ' ' || u.last_name AS name
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

func CreateParkingLot(c *gin.Context) {
	var newParkingLot models.ParkingLot
	var owner models.Owner_onlyName

	if err := c.ShouldBind(&newParkingLot); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBind(&owner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate Data

	if err := models.Validate.Struct(newParkingLot); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.Validate.Struct(owner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user exists and is an owner
	var user models.User
	err := database.DB.Get(&user, `SELECT * FROM users WHERE first_name = $1 AND last_name = $2 LIMIT 1`, owner.First_Name, owner.Last_Name)
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

	// Create parking lot
	newParkingLot.OwnerID = user.ID
	_, err = database.DB.NamedExec(
		`INSERT INTO parking_lots (id, pl_name, addr_street, addr_number, cep, owner_id) `+
			`VALUES (:id, :pl_name, :addr_street, :addr_number, :cep, :owner_id)`, &newParkingLot)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Estacionamento criado com sucesso!")
}
