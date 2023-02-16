package services

import (
	"errors"
	"os"
	"teller/db"
	"teller/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func Login(user models.SignInInput) (string, error) {
	result := models.User{}
	username := user.Username
	password := user.Password

	err := db.GetDB().Where("username = ? AND password = ? AND verified = true", username, password).Find(&result).Error
	if err != nil {
		return ``, err
	}

	if result.Name == `` {
		return ``, errors.New("user")

	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(12 * time.Hour)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		// If there is an error in creating the JWT return an internal server error

		return ``, errors.New("token")
	}

	return tokenString, nil
}
