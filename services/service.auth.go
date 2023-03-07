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

	// err := db.GetDB().Where("username = ? AND verified = true", username).Find(&result).Error
	if err != nil {
		return ``, err
	}

	// isMatch := CheckPasswordHash(result.Password, password)
	// if !isMatch {
	// 	return ``, errors.New("error, not match encrypt")
	// }
		
	if result.Name == `` {
		return ``, errors.New("error, invalid user")

	}

	expirationTime := time.Now().Add(12 * time.Hour)

	claims := &Claims{
		User: result,
		// Username: result.Username,
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
