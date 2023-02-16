package models

import "gorm.io/gorm"

type GetDetail struct {
	gorm.Model

	ID            uint   `json:"id"`
	TransactionID string `json:"transactionID" gorm:"type:varchar(250)"`
	AccountNumber string `json:"accountNumber" gorm:"type:varchar(250)"`
}

func (*GetDetail) TableName() string {
	return "getdetail"
}
