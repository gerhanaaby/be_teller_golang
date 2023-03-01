/**
 * @author [Fajar Dwi Nur Racmadi]
 * @email [fajar.d.rachmadi@banksinarmas.com]
 * @create date 2023-02-14
 * @modify date 2023-02-27
 * @desc [User Controller (Login)]
 */
package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"teller/db"
	"teller/models"
	"teller/services"

	"github.com/gin-gonic/gin"
)

type AuthStatus struct {
	Status   string `json:"status"`
	Message  string `json:"message"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func UserLoginController(c *gin.Context) {
	request := models.SignInInput{}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, AuthStatus{
			Status:  "Fail",
			Message: "unable to login due error",
		})
		return
	}

	if request.Username == `` || request.Password == `` {
		c.AbortWithError(http.StatusBadRequest, errors.New("empy username or password"))
		c.JSON(http.StatusBadRequest, AuthStatus{
			Status:  "Fail",
			Message: "empty username or password",
		})
		return
	}

	LoginToken, err := services.Login(request)
	if err != nil{
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, AuthStatus{
			Status:  "Fail",
			Message: "Wrong Username or Password",
		})
		return
	}

	if LoginToken == `` {
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, AuthStatus{
			Status:  "Fail",
			Message: "unable to login due error getting token" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, AuthStatus{
		Status:   "Success",
		Message:  "Login berhasil.",
		Username: request.Username,
		Token:    LoginToken,
	})
}

func CreateUser(c *gin.Context) {
	var err error

	request := models.User{}
	if err = c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, AuthStatus{
			Status:  "Fail",
			Message: "error, "+err.Error(),
		})
		return
	}

	if len(request.Username) < 8 || len(request.Username) > 15 {
		c.JSON(http.StatusBadRequest, AuthStatus{
			Status:  "Fail",
			Message: "error, username min 8 and max 15 character",
		})
		return 
	}

	if len(request.Password) < 8 || len(request.Password) > 15 {
		c.JSON(http.StatusBadRequest, AuthStatus{
			Status:  "Fail",
			Message: "error, password min 8 and max 15 character",
		})
		return 
	}

	if len(request.Nik) > 6 || len(request.Nik) < 6 {
		c.JSON(http.StatusBadRequest, AuthStatus{
			Status:  "Fail",
			Message: "error, NIK must exact 6 character",
		})
		return
	}

	// request.Password, err = services.HashPassword(request.Password)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, AuthStatus{
	// 		Status:  "Fail",
	// 		Message: "error, "+err.Error(),
	// 	})
	// 	return 
	// }

	request.Password = services.GetMD5Hash(request.Password)

	err = db.GetDB().Create(&request).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, AuthStatus{
			Status:  "Fail",
			Message: "error, "+err.Error(),
		})
		return
	}

	fmt.Println(request)

	c.JSON(http.StatusOK, AuthStatus{
		Status:  "Success",
		Message: "user successfully created",
	})

}
// func RefreshToken(c *gin.Context) {
// 	reqToken := c.Request.Header.Get("Authorization")
// 	if reqToken == ""{
// 		c.AbortWithError(http.StatusInternalServerError, errors.New("error, empty token"))
// 		c.JSON(http.StatusBadRequest, AuthStatus{
// 		Status: "Fail", 
// 		Message: "error, empty token"})
// 	}
	
// 	token := strings.Replace(reqToken, "Bearer ", "", 1)
		
// 	claims := &Claims{}
// 	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(inits.Cfg.TokenSecret), nil
// 	})
// 	if err != nil {
// 		if err == jwt.ErrSignatureInvalid {
// 			c.AbortWithError(http.StatusInternalServerError, err)
// 			c.JSON(http.StatusUnauthorized, AuthStatus{
// 			Status: "Fail", 
// 			Message: err.Error()})
// 		return
// 		}
// 		c.AbortWithError(http.StatusInternalServerError, err)
// 		c.JSON(http.StatusBadRequest, AuthStatus{
// 		Status: "Fail", 
// 		Message: err.Error()})
// 		return
// 	}
// 	if !tkn.Valid {
// 		c.AbortWithError(http.StatusInternalServerError, err)
// 		c.JSON(http.StatusUnauthorized, AuthStatus{
// 		Status: "Fail", 
// 		Message: err.Error()})
// 		return
// 	}
// 	// (END) The code until this point is the same as the first part of the `Welcome` route

// 	// We ensure that a new token is not issued until enough time has elapsed
// 	// In this case, a new token will only be issued if the old token is within
// 	// 30 seconds of expiry. Otherwise, return a bad request status
// 	if time.Until(claims.ExpiresAt.Time) > 30*time.Second {
// 		c.AbortWithError(http.StatusInternalServerError, errors.New("token expiry"))
// 		c.JSON(http.StatusBadRequest, AuthStatus{
// 		Status: "Fail", 
// 		Message: "token expiry"})
// 		return
// 	}

// 	// Now, create a new token for the current use, with a renewed expiration time
// 	expirationTime := time.Now().Add(5 * time.Minute)
// 	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
// 	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := newToken.SignedString([]byte(inits.Cfg.TokenSecret))
// 	if err != nil {
// 		c.AbortWithError(http.StatusInternalServerError, errors.New("error, can't generate new token"))
// 		c.JSON(http.StatusBadRequest, AuthStatus{
// 		Status: "Fail", 
// 		Message: "can't generate new token"})
// 		return
// 	}

// 	// Set the new token as the users `token` cookie
// 	c.JSON(http.StatusOK, AuthStatus{
// 		Status: "Success", 
// 		Message: "Verified",
// 		Token: tokenString,
// 	})

// 	// http.SetCookie(w, &http.Cookie{
// 	// 	Name:    "token",
// 	// 	Value:   tokenString,
// 	// 	Expires: expirationTime,
// 	// })
// }
