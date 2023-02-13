package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model

	CustId  string `json:"custId" gorm:"type:varchar(10) not null" valid:"required~custId is required"`
	Name    string `json:"name" gorm:"type:varchar(50) not null" valid:"required~name is required"`
	Address string `json:"address" gorm:"type:varchar(50) not null" valid:"required~address is required"`
	City    string `json:"city" gorm:"type:varchar(50) not null" valid:"required~city is required"`
	State   string `json:"state" gorm:"type:varchar(50) not null" valid:"required~state is required"`
	Zip     int    `json:"zip" gorm:"type:integer not null" valid:"required~zip is required"`
}

type Customerjadi struct {
	ID      uint
	CustId  string `json:"custId" gorm:"type:varchar(10) not null"`
	Name    string `json:"name" gorm:"type:varchar(50) not null"`
	Address string `json:"address" gorm:"type:varchar(50) not null"`
	City    string `json:"city" gorm:"type:varchar(50) not null"`
	State   string `json:"state" gorm:"type:varchar(50) not null"`
	Zip     int    `json:"zip" gorm:"type:integer not null"`
}

func (*Customer) TableName() string {
	return "customer"
}

func (*Customerjadi) TableName() string {
	return "customer"
}

func (v *Customer) Validate() error {
	_, err := govalidator.ValidateStruct(v)
	if err != nil {
		return err
	}
	return nil
}
