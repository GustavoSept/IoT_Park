package helpers

import (
	"crypto/rsa"
	"errors"
)

const (
	privKeyPath = "keys/app.rsa"
	pubKeyPath  = "keys/app.rsa.pub"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

func InitJWT() error {
	return errors.New("IMPLEMENT!")
}

func CreateNewTokens(uuid string, role string) (authTokenString, refreshTokenString, csrfSecret string, err error) {
	return
}

func CheckAndRefreshTokens(
	oldAuthTokenString string, oldRefreshTokenString string, oldCsrfSecret string) (
	newAuthTokenString, newRefreshTokenString, newCsrfSecret string, err error) {
	return

}

func createAuthTokenString(uuid string, role string, csrfSecret string) (authTokenString string, err error) {
	return
}
func createRefreshTokenString(uuid string, role string, csrfString string) (refreshTokenString string, err error) {
	return
}
func updateRefreshTokenExp(oldRefreshTokenString string) (newRefreshTokenString string, err error) {
	return
}

func updateAuthTokenString(refreshTokenString string, oldAuthTokenString string) (newAuthTokenString, csrfSecret string, err error) {
	return
}
func RevokeRefreshToken(refreshTokenString string) error {
	return errors.New("IMPLEMENT")
}
func updateRefreshTokenCsrf(oldRefreshTokenString string, newCsrfString string) (newRefreshTokenString string, err error) {
	return
}
func GrabUUID(authTokenString string) (string, error) {
	return "", errors.New("IMPLEMENT")
}
