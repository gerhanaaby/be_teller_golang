package models

import "gorm.io/gorm"

/**
 * @author [Fajar Dwi Nur Racmadi]
 * @email [fajar.d.rachmadi@banksinarmas.com]
 * @create date 2023-02-16
 * @modify date 2023-02-20
 * @desc [Model api]
 */
type Apis struct {
	gorm.Model

	// ID     			 int64     `json:"userid" gorm:"type:serial;primary_key"`
	Name             string    `gorm:"type:varchar(255);uniqueIndex;not null"`
	Url   			 string    `gorm:"type:varchar(255);not null"`
	Key   			 string    `gorm:"type:varchar(255);not null"`
	Value 			 string    `gorm:"type:varchar(255);not null"`
}

var ApiList []Apis
var ApiMap map[string]Apis

func (*Apis) TableName() string {
	return "apis"
}
