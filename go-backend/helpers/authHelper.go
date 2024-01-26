package helpers

import (
	"crypto/rand"
	"encoding/base64"

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

// returns password, salt, error
func HashPassword(password string, salt []byte) (string, error) {
	hash := argon2.IDKey(
		[]byte(password), // password
		salt,             // salt
		3,                // time, number of iterations
		64*1024,          // amount of memory to use
		4,                // threads
		32,               // Key Length
	)
	return base64.RawStdEncoding.EncodeToString(hash), nil
}

func CheckPassword(providedPassword string, storedHash string, salt []byte) (bool, error) {
	hash, err := HashPassword(providedPassword, salt)
	if err != nil {
		return false, err // Error while hashing
	}

	// Compare the hash of the provided password with the stored hash
	match := hash == storedHash

	return match, nil
}
