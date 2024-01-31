package helpers

import (
	"crypto/rsa"
	"errors"
	"go-backend/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const (
	PRIV_KEY_PATH = "keys/app.rsa"
	PUB_KEY_PATH  = "keys/app.rsa.pub"
)

var (
	VERIFY_KEY *rsa.PublicKey
	SIGN_KEY   *rsa.PrivateKey
)

func InitJWT() error {
	signBytes, err := os.ReadFile(PRIV_KEY_PATH)
	if err != nil {
		return err
	}

	SIGN_KEY, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return err
	}

	verifyBytes, err := os.ReadFile(PUB_KEY_PATH)
	if err != nil {
		return err
	}

	VERIFY_KEY, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return err
	}

	return nil
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
