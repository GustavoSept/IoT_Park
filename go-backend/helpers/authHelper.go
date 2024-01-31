package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"go-backend/database"
	"go-backend/models"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

func MatchCredentialsWithUser(inPassword, inEmail string) (uuid.UUID, string, error) {
	// struct is a SuperSet of models.User_Auth + models.User.Office_Level
	// I've done it this way to request data from the DB only once
	var valUser struct {
		UserID       uuid.UUID `db:"user_id"`
		Email        string    `db:"email"`
		PasswordHash string    `db:"password_hash"`
		Salt         []byte    `db:"salt"`
		Office_Level string    `db:"office_level"`
	}

	log.Printf("\ninEmail: %v\ninPassword: %v\n", inEmail, inPassword)

	if err := database.DB.Get(&valUser, `
		SELECT ua.user_id, ua.password_hash, ua.salt, u.office_level
		FROM users_authentication ua
		JOIN users u ON ua.user_id = u.id
		WHERE ua.email=$1`, inEmail); err != nil {

		log.Printf("Error fetching user data: %v", err)
		return uuid.Nil, "", errors.New("internal server error")
	}

	if valUser.UserID == uuid.Nil {
		log.Println("Error fetching user: empty UUID (query returned empty)")
		return uuid.Nil, "", errors.New("Login Inválido")
	}

	if ok, err := CheckPassword(inPassword, valUser.PasswordHash, valUser.Salt); err != nil {
		return uuid.Nil, "", err
	} else if ok {
		return valUser.UserID, valUser.Office_Level, nil
	} else {
		log.Println("Error in CheckPassword(): passwords don't match")
		return uuid.Nil, "", errors.New("Login Inválido")
	}
}

func GenerateCSRFSecret() (string, error) {
	s, err := GenerateSalt(32)
	return base64.URLEncoding.EncodeToString(s), err
}

func CreateToken(userId string, role string) (string, error) {
	claims := models.TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			// Set standard claims like expiry, issuer, subject etc.
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(models.AuthTokenValidTime)),
			// ... other standard claims
		},
		Role: role,
		Csrf: "some_csrf_token", // Generate a CSRF token as per your requirement
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
