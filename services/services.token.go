package services

import (
	"errors"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)


type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func CheckToken(reqToken string) (bool, error) {
	var isValid bool
	var err error

	if reqToken == "" {
		return false, errors.New("error, empty token")
	}

	token := strings.Replace(reqToken, "Bearer ", "", 1)

	if isValid, err = ValidateToken(token); err != nil {
		return false, err
	}

	if !isValid {
		return false, errors.New("unable to precess due to invalid or expired login")
	}

	return true, nil
}

func ValidateToken(reqToken string) (bool, error) {
	if reqToken == ""{
		return false, errors.New("empty token")
	}
	
	token := strings.Replace(reqToken, "Bearer ", "", 1)

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
		return false, errors.New("invalid token")
		}
		return false, errors.New("invalid token")
	}
	if !tkn.Valid {
		return false, errors.New("invalid token")
	}

	return true, nil
}
