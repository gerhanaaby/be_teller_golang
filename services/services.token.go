/**
 * @author [Fajar Dwi Nur Racmadi]
 * @email [fajar.d.rachmadi@banksinarmas.com]
 * @create date 2023-02-14
 * @modify date 2023-02-20
 * @desc [Cek Keabsahan token pada jwt session]
 */

package services

import (
	"errors"
	"fmt"
	"strings"
	"teller/inits"
	"teller/models"

	"github.com/golang-jwt/jwt/v4"
)

 type Claims struct {
	User	models.User `josn:"User"`
	jwt.RegisteredClaims
}

func CheckToken(reqToken string) (claims jwt.MapClaims, isValid bool, err error) {

	if reqToken == "" {
		return claims, false, errors.New("error, empty token")
	}

	token, err := jwt.Parse(
		strings.Replace(reqToken, "Bearer ", "", 1),
		 func(token *jwt.Token) (interface{}, error) {
			return []byte(inits.Cfg.TokenSecret), nil
		})
	if err != nil {
		return nil, false, err
	}
	_,a := token.Claims.(jwt.MapClaims)
	if !a {
		fmt.Println()
	}

	if claims, ok  := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claims.VerifyExpiresAt(0, true)
		return claims, true, nil
	} else {
		return claims, false, errors.New("invalid JWT Token")
	}
}


// func ValidateToken(reqToken string, claims Claims) (models.User, bool, error) {

// 	token, err := jwt.ParseWithClaims(
// 		reqToken, 
// 		claims, 
// 		func(reqToken *jwt.Token) (interface{}, error) {
// 			return []byte(inits.Cfg.TokenSecret), nil
// 		})
// 	if err != nil {
// 		if err == jwt.ErrSignatureInvalid {
// 			return nil, false, errors.New("invalid token")
// 		}
// 		return false, errors.New("invalid token")
// 	}
// 	if !token.Valid {
// 		return false, errors.New("invalid token")
// 	}

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		return claims, true
// 	} else {
// 		log.Printf("Invalid JWT Token")
// 		return nil, false
// 	}

// 	return true, nil
// }
