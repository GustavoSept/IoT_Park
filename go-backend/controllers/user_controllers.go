package controllers

import (
	"encoding/hex"
	"fmt"
	"net/http"

	"go-backend/database"
	"go-backend/helpers"
	"go-backend/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	SALT_SIZE = 16
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

func GetAllUsersAuth(c *gin.Context) {
	var usersAuthList []models.User_Auth
	err := database.DB.Select(&usersAuthList, "SELECT * FROM users_authentication")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, usersAuthList)
}

func CreateUser(c *gin.Context) {
	var newUser models.User
	var newUserAuth models.User_Auth

	var salt []byte
	var hashedPass string
	var err error

	// Attempt to parse request body into models
	if err := c.Bind(&newUser); err != nil {
		return
	}
	if err := c.Bind(&newUserAuth); err != nil {
		return
	}

	rawPassword := c.PostForm("raw_password")

	// Validate newUser
	if err := models.Validate.Struct(newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.Validate.Struct(newUserAuth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ------------------------ Check parkingLotID
	// If user is not an owner, and there's no parkingLotID, we can't create the user
	// Owner users can be created without being immediately associated with a parkingLotID
	plID, pl_exists := c.Get("parkingLotID")
	if !pl_exists && newUser.Office_Level != "dono" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't create non-owner user without a parking lot"})
		return
	}

	// Start new transaction
	tx, err := database.DB.Beginx()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Generate a UUID
	newUser.ID = uuid.New()
	newUserAuth.UserID = newUser.ID

	// Generate Salt
	salt, err = helpers.GenerateSalt(SALT_SIZE)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Hash Password
	hashedPass, err = helpers.HashPassword(rawPassword, salt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Insert newUser into users table
	_, err = tx.NamedExec(
		`INSERT INTO users (id, first_name, last_name, office_level)
						VALUES (:id, :first_name, :last_name, :office_level)`, &newUser)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convert salt to hexadecimal
	hexSalt := make([]byte, hex.EncodedLen(len(salt)))
	hex.Encode(hexSalt, salt)

	newUserAuth.Salt = hexSalt
	newUserAuth.PasswordHash = hashedPass

	// Insert newUserAuth into users_authentication table
	_, err = tx.NamedExec(`
						INSERT INTO users_authentication (user_id, email, password_hash, salt)
						VALUES (:user_id, :email, :password_hash, :salt)`, &newUserAuth)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Insert into parking_lot_employees if necessary
	if pl_exists && newUser.Office_Level != "dono" {
		_, err = tx.Exec(`
						INSERT INTO parking_lot_employees (user_id, parking_lot_id)
						VALUES ($1, $2)`, newUser.ID, plID)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Send success response
	c.JSON(http.StatusOK, fmt.Sprintf("Bazinga! A conta de %s foi criada com sucesso!", newUser.First_Name))
}

func UpdateUser(c *gin.Context) {
	// Logic to update an User
	c.JSON(http.StatusOK, gin.H{"message": "PUT request called"})
}

func DeleteUser(c *gin.Context) {
	// Logic to delete an User
	c.JSON(http.StatusOK, gin.H{"message": "DELETE request called"})
}
