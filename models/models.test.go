package models

import "gorm.io/gorm"

type MyTest struct {
	gorm.Model

	Notix 			string		`json:"notix" gorm:"type:varchar(255);not null;primary_key"`
	// Name            string		`json:"name" gorm:"type:varchar(255);not null"`
	CompanyID 		string		`json:"CompanyID" gorm:"column:CompanyID;type:varchar(255);not null"`
  	Company   		Company		`gorm:"references:CompanyID"` // use Company.CompanyID as references
}

type Company struct {
  CompanyID   string `gorm:"uniqueIndex;not null"`
  Name        string
}

