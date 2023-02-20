package services

import (
	"errors"
	"teller/db"
	"teller/inits"
	"teller/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

/**
 * @author [Fajar Dwi Nur Racmadi]
 * @email [fajar.d.rachmadi@banksinarmas.com]
 * @create date 2023-02-16
 * @modify date 2023-02-20
 * @desc [Authentikasi login dan generate jwt token]
 */
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

	expirationTime := time.Now().Add(12 * time.Hour)

	claims := &Claims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	tokenString, err := token.SignedString([]byte(inits.Cfg.TokenSecret))
	if err != nil {
		return ``, errors.New("token")
	}

	return tokenString, nil
}
