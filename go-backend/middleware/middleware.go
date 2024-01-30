package middleware

import (
	"fmt"
	"go-backend/database"
	"go-backend/helpers"
	"go-backend/models"
	"log"
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
	/*
		For unauthorized access, it's the frontend's job to redirect the user to either the login page,
		or to /api/tokenrefresh endpoint to refresh their AuthToken based on RefreshToken (expiration) validity
	*/

	// Read cookies
	authCookie, err := c.Cookie("AuthToken")
	if err == http.ErrNoCookie {
		log.Println("Unauthorized attempt! No auth cookie")
		helpers.NullifyTokenCookies(c)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	} else if err != nil {
		log.Panic(fmt.Sprintf("panic: %+v", err))
		helpers.NullifyTokenCookies(c)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	refreshCookie, err := c.Cookie("RefreshToken")
	if err == http.ErrNoCookie {
		log.Println("Unauthorized attempt! No refresh cookie")
		helpers.NullifyTokenCookies(c)
		c.Redirect(http.StatusFound, "/login")
		return
	} else if err != nil {
		log.Panic(fmt.Sprintf("panic: %+v", err))
		helpers.NullifyTokenCookies(c)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Grab the CSRF token
	requestCsrfToken := grabCsrfFromReq(c)

	// Check the JWTs for validity
	authTokenString, refreshTokenString, csrfSecret, err := helpers.CheckAndRefreshTokens(authCookie, refreshCookie, requestCsrfToken)
	if err != nil {
		if err.Error() == "Unauthorized" {
			log.Println("Unauthorized attempt! JWT's not valid!")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		} else {
			log.Panic(fmt.Sprintf("panic: %+v", err))
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}
	log.Println("Successfully recreated jwts")

	// Set headers and cookies
	helpers.SetAuthAndRefreshCookies(c, authTokenString, refreshTokenString)
	c.Writer.Header().Set("X-CSRF-Token", csrfSecret)

	c.Next()
	return
}

func grabCsrfFromReq(c *gin.Context) string {
	csrfToken := c.PostForm("X-CSRF-Token")

	if csrfToken == "" {
		csrfToken = c.GetHeader("X-CSRF-Token")
	}

	return csrfToken
}
