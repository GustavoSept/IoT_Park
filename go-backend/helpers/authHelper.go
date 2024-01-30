package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"go-backend/database"
	"go-backend/models"
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/argon2"
)

func GenerateSalt(size int) ([]byte, error) {
	salt := make([]byte, size)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

// returns hash, error
func HashPassword(password string, salt []byte) (string, error) {
	hash := argon2.IDKey(
		[]byte(password), // password
		salt,             // salt
		3,                // time, number of iterations
		64*1024,          // amount of memory to use
		4,                // threads
		32,               // Key Length
	)
	encodedHash := base64.RawStdEncoding.EncodeToString(hash)
	if encodedHash == "" {
		return "", errors.New("failed to encode hash")
	}

	return encodedHash, nil
}

func CheckPassword(inPassword string, storedHash string, salt []byte) (bool, error) {
	hash, err := HashPassword(inPassword, salt)
	if err != nil {
		return false, err // Error while hashing
	}

	// Compare the hash of the provided password with the stored hash
	match := hash == storedHash

	return match, nil
}

func MatchLoginDataWithUser(inPassword, inEmail string) (bool, error) {
	var valUser models.User_Auth

	if err := database.DB.Get(&valUser, `
            SELECT user_id, password_hash, salt
            FROM users_authentication
            WHERE email=$1`, inEmail); err != nil {
		// Log the error for internal tracking
		log.Printf("Error fetching user data: %v", err)
		return false, errors.New("internal server error")
	}

	if valUser.UserID == uuid.Nil {
		return false, errors.New("Login Inválido")
	}

	return CheckPassword(inPassword, valUser.PasswordHash, valUser.Salt)
}

func GenerateCSRFSecret() (string, error) {

}

func CreateToken(userId string, role string) (string, error) {

}
