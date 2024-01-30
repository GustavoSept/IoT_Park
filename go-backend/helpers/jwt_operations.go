package helpers

import (
	"crypto/rsa"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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
func revokeRefreshToken(refreshTokenString string) error {
	return errors.New("IMPLEMENT")
}
func updateRefreshTokenCsrf(oldRefreshTokenString string, newCsrfString string) (newRefreshTokenString string, err error) {
	return
}
func GrabUUID(authTokenString string) (string, error) {
	return "", errors.New("IMPLEMENT")
}

func SetAuthAndRefreshCookies(c *gin.Context, authTokenString string, refreshTokenString string) {
	c.SetCookie("AuthToken", authTokenString, 3600, "/", "", false, true)
	c.SetCookie("RefreshToken", refreshTokenString, 3600, "/", "", false, true)
}

func NullifyTokenCookies(c *gin.Context) {
	// Invalidate Tokens locally
	c.SetCookie("AuthToken", "", -1, "/", "", false, true)
	c.SetCookie("RefreshToken", "", -1, "/", "", false, true)

	// If present, revoke the refresh token from the database
	refreshToken, err := c.Cookie("RefreshToken")
	if err == http.ErrNoCookie {
		// Do nothing, there is no refresh cookie present
		return
	} else if err != nil {
		log.Panicf("panic: %+v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	revokeRefreshToken(refreshToken)
}
