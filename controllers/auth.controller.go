package controllers

// import (
// 	"fmt"
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// )

// // import (
// // 	"net/http"
// // 	"strings"
// // 	"teller/inits"
// // 	"teller/models"
// // 	"teller/utils"
// // 	"time"

// // 	"github.com/gin-gonic/gin"
// // 	"github.com/thanhpk/randstr"

// // 	// "github.com/wpcodevo/golang-gorm-postgres/initializers"

// // 	// "github.com/wpcodevo/golang-gorm-postgres/models"

// // 	// "github.com/wpcodevo/golang-gorm-postgres/utils"
// // 	"gorm.io/gorm"
// // )

// /**
//  * @author [Fajar Dwi Nur Racmadi]
//  * @email [fajar.d.rachmadi@banksinarmas.com]
//  * @create date 2023-02-14
//  * @modify date 2023-02-20
//  * @desc [Create Account, Verify Account]
//  */
// // type AuthController struct {
// // 	DB *gorm.DB
// // }

// // func NewAuthController(DB *gorm.DB) AuthController {
// // 	return AuthController{DB}
// // }

// // // [...] SignUp User
// // func (ac *AuthController) SignUpUser(ctx *gin.Context) {
// // 	var payload *models.SignUpInput

// // 	if err := ctx.ShouldBindJSON(&payload); err != nil {
// // 		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
// // 		return
// // 	}

// // 	if payload.Password != payload.PasswordConfirm {
// // 		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Passwords do not match"})
// // 		return
// // 	}

// // 	hashedPassword, err := utils.HashPassword(payload.Password)
// // 	if err != nil {
// // 		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
// // 		return
// // 	}

// // 	now := time.Now()
// // 	newUser := models.User{
// // 		Name:      payload.Name,
// // 		Email:     strings.ToLower(payload.Email),
// // 		Password:  hashedPassword,
// // 		Roles:      "user",
// // 		Division:  "local",
// // 		Verified:  false,
// // 		Photo:     payload.Photo,
// // 		CreatedAt: now,
// // 		UpdatedAt: now,
// // 	}

// // 	result := ac.DB.Create(&newUser)

// // 	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
// // 		ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "User with that email already exists"})
// // 		return
// // 	} else if result.Error != nil {
// // 		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": "Something bad happened"})
// // 		return
// // 	}

// // 	// config, _ := inits.LoadConfig(".")

// // 	// Generate Verification Code
// // 	code := randstr.String(20)

// // 	verification_code := utils.Encode(code)

// // 	// Update User in Database
// // 	newUser.VerificationCode = verification_code
// // 	ac.DB.Save(newUser)

// // 	var firstName = newUser.Name

// // 	if strings.Contains(firstName, " ") {
// // 		firstName = strings.Split(firstName, " ")[1]
// // 	}

// // 	// ? Send Email
// // 	emailData := utils.EmailData{
// // 		URL:       config.ClientOrigin + "/verifyemail/" + code,
// // 		FirstName: firstName,
// // 		Subject:   "Your account verification code",
// // 	}

// // 	utils.SendEmail(&newUser, &emailData)

// // 	message := "We sent an email with a verification code to " + newUser.Email
// // 	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "message": message})
// // }

// // [...] SignOut User
// func (ac *AuthController) LogoutUser(ctx *gin.Context) {

// 	if GlobalCorsEnabled {
// 		enableCors(&w)
// 	}

// 	tkn, failedAuth := CheckJwtAuth(w, r)
// 	if failedAuth {
// 		return
// 	}

// 	//remove old token from list
// 	whiteListTokens = remove(whiteListTokens, tkn.Raw)

// 	w.Write([]byte(fmt.Sprintf("LoggedOut")))

// 	ctx.SetCookie("token", "", -1, "/", "localhost", false, true)
// 	ctx.JSON(http.StatusOK, gin.H{"status": "success"})

// time.Now().UnixMicro()

// }