/**
 * @author [Fajar Dwi Nur Racmadi]
 * @email [fajar.d.rachmadi@banksinarmas.com]
 * @create date 2023-02-14
 * @modify date 2023-02-20
 * @desc [Model user]
 */
package models

import (
	"time"

	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID     			 int64     `json:"userid" gorm:"type:serial;primary_key"`
	Name             string    `json:"name" gorm:"type:varchar(255);not null"`
	Username 		 string    `json:"username" gorm:"uniqueIndex;not null"`
	Email            string    `json:"email" gorm:"uniqueIndex;not null"`
	Password         string    `json:"password" gorm:"not null"`
	Nik				 string    `json:"nik" gorm:"type:varchar(255); uniqueIndex;not null"`
	BranchCode		 string    `json:"branch_code" gorm:"type:varchar(255);not null"`
	Roles            string    `json:"role" gorm:"type:varchar(255);not null"`
	Division         string    `json:"div" gorm:"type:varchar(255);not null"`
	Photo            string    `json:"photo" gorm:"not null"`
	VerificationCode string	   `json:"verify_code" gorm:"not null"`
	Verified         bool      `json:"verified" gorm:"not null"`
	CreatedAt        time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt        time.Time `json:"updated_at" gorm:"not null"`
}

type SignUpInput struct {
	Name            string `json:"name" binding:"required"`
	Username		string `json:"username" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
	Photo           string `json:"photo" binding:"required"`
}

type SignInInput struct {
	Username    string `json:"username"  binding:"required"`
	Password 	string `json:"password"  binding:"required"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty"`
	Photo     string    `json:"photo,omitempty"`
	Provider  string    `json:"provider"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

