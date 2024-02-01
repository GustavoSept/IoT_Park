package helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	local_SALT_SIZE = 16 // salt size in bytes
)

func TestGenerateSalt(t *testing.T) {
	salt, err := GenerateSalt(local_SALT_SIZE)
	assert.NoError(t, err)
	assert.Len(t, salt, local_SALT_SIZE, "Salt should be of the correct length")

	salt2, err := GenerateSalt(local_SALT_SIZE)
	assert.NoError(t, err)
	assert.NotEqual(t, salt, salt2, "Subsequent salts should be different")
}

func TestHashPassword(t *testing.T) {
	password := "testPassword"
	salt, err := GenerateSalt(local_SALT_SIZE)
	hash1_A, err := HashPassword(password, salt)
	fmt.Printf("\nTestHashPassword's Hash: %v\n", hash1_A)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash1_A, "Hash should not be empty")

	salt2, err := GenerateSalt(local_SALT_SIZE)
	hash2, err := HashPassword(password, salt2)
	assert.NoError(t, err)
	assert.NotEqual(t, hash1_A, hash2, "Hashes of the same password should be different")

	hash1_B, err := HashPassword(password, salt)
	assert.NoError(t, err)
	assert.Equal(t, hash1_A, hash1_B, "Hashing the same password, with the same salt multiple times should be equal")

	password2 := "new2@password"
	hash3, err := HashPassword(password2, salt)
	assert.NoError(t, err)
	assert.NotEqual(t, hash1_A, hash3, "Hashing different passwords with the same salt should be different")
}

func TestCheckPassword(t *testing.T) {
	providedPassword := "testPassword"
	salt, err := GenerateSalt(local_SALT_SIZE)
	assert.NoError(t, err)

	// Hash the provided password with the salt
	storedHash, err := HashPassword(providedPassword, salt)
	assert.NoError(t, err)

	// Check the password using CheckPassword
	match, err := CheckPassword(providedPassword, storedHash, salt)
	assert.NoError(t, err)
	assert.True(t, match, "Password should match")

	// Test with incorrect password
	incorrectPassword := "incorrectPassword"
	match, err = CheckPassword(incorrectPassword, storedHash, salt)
	assert.NoError(t, err)
	assert.False(t, match, "Password should not match")
}
