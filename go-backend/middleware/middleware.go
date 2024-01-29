package middleware

import (
	"go-backend/database"
	"go-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// If there's parkingLot info in the (form) request,
// get the parkingLot ID, store it in gin's Context.
// If there's no parkingLot ID, go to c.Next()
func GetParkingLotContext(c *gin.Context) {
	var pl_req models.ParkingLot

	if err := c.ShouldBind(&pl_req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Couldn't bind parkinglot": err.Error()})
		c.Next()
		return
	}

	// if any of these fields are missing, we can't search for a pl_req.ID
	if pl_req.PLot_Name == "" || pl_req.AddrStreet == "" || pl_req.AddrNumber == 0 {
		c.Next()
		return
	}

	// Get pl.req.ID from database
	query := `
			SELECT id
			FROM parking_lots
			WHERE pl_name = $1 AND addr_street = $2 AND addr_number = $3
		`

	var plID uuid.UUID
	err := database.DB.Get(&plID, query, pl_req.PLot_Name, pl_req.AddrStreet, pl_req.AddrNumber)
	if err != nil || plID == uuid.Nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't retrieve parking lot ID: " + err.Error()})
		c.Next()
		return
	}

	// Store the parking lot ID in Gin's context for subsequent handlers
	c.Set("parkingLotID", plID)

	c.Next()
}

func AuthRequired(c *gin.Context) {
	userAuthenticated := false // TODO: implement actual authentication check

	if !userAuthenticated {
		// If not authenticated, redirect to login page or return an error
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}

	c.Next()
}
