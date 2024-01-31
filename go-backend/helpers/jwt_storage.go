package helpers

import (
	"encoding/base64"
	"go-backend/database"
	"log"
)

func CreateAndStoreRefreshToken() (jti string, err error) {
	salt, err := GenerateSalt(32)
	if err != nil {
		return "", err
	}
	jti = base64.URLEncoding.EncodeToString(salt)

	// Insert jti into the database
	_, err = database.DB.Exec("INSERT INTO jwt_auth (jti) VALUES ($1)", jti)
	if err != nil {
		return "", err
	}

	return jti, nil
}

func DeleteRefreshToken(jti string) error {
	_, err := database.DB.Exec("DELETE FROM jwt_auth WHERE jti = $1", jti)
	return err
}

func isValidRefreshToken(jti string) bool {
	var exists bool
	err := database.DB.Get(&exists, "SELECT EXISTS(SELECT 1 FROM jwt_auth WHERE jti = $1)", jti)
	if err != nil {
		log.Println("Error checking for jti:", err)
		return false
	}
	return exists
}
