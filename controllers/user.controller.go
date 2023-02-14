package controllers

import (
	"errors"
	"net/http"
	"strings"
	"teller/db"
	"teller/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type AuthStatus struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Username string `json:"username"`
	Token string `json:"token"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var SecretKey = []byte("KMZWA87AWAA")

func UserLogin(c *gin.Context) {
	var request models.SignInInput

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var err error
	result := models.User{}

	username := request.Username
	password := request.Password

	err = db.GetDB().Where("username = ? AND password = ?", username, password).Find(&result).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		c.JSON(http.StatusOK, AuthStatus{
			Status: "Fail", 
			Message: "wrong Username or Password"})
		return
	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: request.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		c.AbortWithError(http.StatusInternalServerError, err)
		c.JSON(http.StatusOK, AuthStatus{
			Status: "Fail", 
			Message: err.Error()})
		return
	}

	// http.SetCookie(w, &http.Cookie{
	// 	Name:    "token",
	// 	Value:   tokenString,
	// 	Expires: expirationTime,
	// })

	c.JSON(http.StatusOK, AuthStatus{
			Status: "Success", 
			Message: "Login berhasil.",
			Username: request.Username,
			Token: tokenString,
		})
}

func ValidateToken(reqToken string) (bool, error) {
	if reqToken == ""{
		return false, errors.New("empty token")
	}
	
	token := strings.Replace(reqToken, "Bearer ", "", 1)

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
		return false, err
		}
		return false, err
	}
	if !tkn.Valid {
		return false, errors.New("invalid token")
	}

	return true, nil
}

func RefreshToken(c *gin.Context) {
	reqToken := c.Request.Header.Get("Authorization")
	if reqToken == ""{
		c.AbortWithError(http.StatusInternalServerError, errors.New("error, empty token"))
		c.JSON(http.StatusBadRequest, AuthStatus{
		Status: "Fail", 
		Message: "error, empty token"})
	}
	
	token := strings.Replace(reqToken, "Bearer ", "", 1)
		
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.AbortWithError(http.StatusInternalServerError, err)
			c.JSON(http.StatusUnauthorized, AuthStatus{
			Status: "Fail", 
			Message: err.Error()})
		return
		}
		c.AbortWithError(http.StatusInternalServerError, err)
		c.JSON(http.StatusBadRequest, AuthStatus{
		Status: "Fail", 
		Message: err.Error()})
		return
	}
	if !tkn.Valid {
		c.AbortWithError(http.StatusInternalServerError, err)
		c.JSON(http.StatusUnauthorized, AuthStatus{
		Status: "Fail", 
		Message: err.Error()})
		return
	}
	// (END) The code until this point is the same as the first part of the `Welcome` route

	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// 30 seconds of expiry. Otherwise, return a bad request status
	if time.Until(claims.ExpiresAt.Time) > 30*time.Second {
		c.AbortWithError(http.StatusInternalServerError, errors.New("token expiry"))
		c.JSON(http.StatusBadRequest, AuthStatus{
		Status: "Fail", 
		Message: "token expiry"})
		return
	}

	// Now, create a new token for the current use, with a renewed expiration time
	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := newToken.SignedString(SecretKey)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New("error, can't generate new token"))
		c.JSON(http.StatusBadRequest, AuthStatus{
		Status: "Fail", 
		Message: "can't generate new token"})
		return
	}

	// Set the new token as the users `token` cookie
	c.JSON(http.StatusOK, AuthStatus{
		Status: "Success", 
		Message: "Verified",
		Token: tokenString,
	})

	// http.SetCookie(w, &http.Cookie{
	// 	Name:    "token",
	// 	Value:   tokenString,
	// 	Expires: expirationTime,
	// })
}