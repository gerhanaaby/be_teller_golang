package utils

import (
	"fmt"
	"teller/models"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(userLogin models.SignInInput, secretJWTKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["user"] = "username"

	// now := time.Now().UTC()
	// claims := token.Claims.(jwt.MapClaims)

	// claims["sub"] = payload
	// claims["exp"] = now.Add(ttl).Unix()
	// claims["iat"] = now.Unix()
	// claims["nbf"] = now.Unix()

	tokenString, err := token.SignedString([]byte(secretJWTKey))

	if err != nil {
		return "", fmt.Errorf("generating JWT Token failed: %w", err)
	}

	return tokenString, nil
}
