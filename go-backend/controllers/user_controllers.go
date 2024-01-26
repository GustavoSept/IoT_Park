package controllers

import (
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

	newUserAuth.PasswordHash = hashedPass

	// Insert into DB, start new Tx

	_, err = database.DB.NamedExec(
		`INSERT INTO users (id, first_name, last_name, office_level)
		VALUES (:id, :first_name, :last_name, :office_level)`, &newUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newUser)
}

func UpdateUser(c *gin.Context) {
	// Logic to update an User
	c.JSON(http.StatusOK, gin.H{"message": "PUT request called"})
}

func DeleteUser(c *gin.Context) {
	// Logic to delete an User
	c.JSON(http.StatusOK, gin.H{"message": "DELETE request called"})
}
