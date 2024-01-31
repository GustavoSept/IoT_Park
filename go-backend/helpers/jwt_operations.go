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
	csrfSecret, err = GenerateCSRFSecret()
	if err != nil {
		return
	}

	refreshTokenString, err = createRefreshTokenString(uuid, role, csrfSecret)

	authTokenString, err = createAuthTokenString(uuid, role, csrfSecret)
	if err != nil {
		return
	}

	return
}

func CheckAndRefreshTokens(
	oldAuthTokenString, oldRefreshTokenString, oldCsrfSecret string) (
	newAuthTokenString, newRefreshTokenString, newCsrfSecret string, err error) {

	if oldCsrfSecret == "" {
		log.Println("No CSRF token!")
		return "", "", "", errors.New("Unauthorized")
	}

	// Parsing the token with claims
	var authTokenClaims models.TokenClaims
	authToken, err := jwt.ParseWithClaims(oldAuthTokenString, &authTokenClaims, func(token *jwt.Token) (interface{}, error) {
		return VERIFY_KEY, nil
	})

	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenMalformed):
			log.Println("That's not even a token")
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			log.Println("Invalid signature")
		case errors.Is(err, jwt.ErrTokenExpired), errors.Is(err, jwt.ErrTokenNotValidYet):
			log.Println("Timing is everything")
			newAuthTokenString, newCsrfSecret, err = updateAuthTokenString(oldRefreshTokenString, oldAuthTokenString)
			if err != nil {
				return "", "", "", errors.New("Unauthorized")
			}

			newRefreshTokenString, err = updateRefreshTokenExp(oldRefreshTokenString)
			if err != nil {
				return "", "", "", errors.New("Unauthorized")
			}

			newRefreshTokenString, err = updateRefreshTokenCsrf(newRefreshTokenString, newCsrfSecret)
			return "", "", "", errors.New("Unauthorized")

		default:
			log.Println("Couldn't handle this token:", err)
		}
		return "", "", "", errors.New("Unauthorized")
	}

	if oldCsrfSecret != authTokenClaims.Csrf {
		log.Println("CSRF token doesn't match jwt!")
		return "", "", "", errors.New("Unauthorized")
	}

	if authToken.Valid {
		log.Println("Auth token is valid")
		newCsrfSecret = authTokenClaims.Csrf
		newRefreshTokenString, err = updateRefreshTokenExp(oldRefreshTokenString)
		newAuthTokenString = oldAuthTokenString
		return newAuthTokenString, newRefreshTokenString, newCsrfSecret, nil
	}

	log.Println("Auth token is not valid")
	return "", "", "", errors.New("Unauthorized")

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
