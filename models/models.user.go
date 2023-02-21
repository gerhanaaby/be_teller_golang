package models

import (
	"time"

	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"gorm.io/gorm"
)

/**
 * @author [Fajar Dwi Nur Racmadi]
 * @email [fajar.d.rachmadi@banksinarmas.com]
 * @create date 2023-02-14
 * @modify date 2023-02-20
 * @desc [Model user]
 */
type User struct {
	gorm.Model

	ID     			 int64     `json:"userid" gorm:"type:serial;primary_key"`
	Name             string    `gorm:"type:varchar(255);not null"`
	Username 		 string    `gorm:"uniqueIndex;not null"`
	Email            string    `gorm:"uniqueIndex;not null"`
	Password         string    `gorm:"not null"`
	Nik				 string    `gorm:"type:varchar(255);not null"`
	BranchCode		 string    `gorm:"type:varchar(255);not null"`
	Roles            string    `gorm:"type:varchar(255);not null"`
	Division         string    `gorm:"type:varchar(255);not null"`
	Photo            string    `gorm:"not null"`
	VerificationCode string	   `gorm:"not null"`
	Verified         bool      `gorm:"not null"`
	CreatedAt        time.Time `gorm:"not null"`
	UpdatedAt        time.Time `gorm:"not null"`
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

