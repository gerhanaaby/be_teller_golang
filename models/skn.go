package models

import "gorm.io/gorm"

type Skn struct {
	gorm.Model

	ID                        uint   `json:"id"`
	SknId                     string `json:"SknId" gorm:"type:varchar(7)"`
	creditAccountNo           string `json:"creditAccountNo" gorm:"type:varchar(50)"`
	amount                    string `json:"amount" gorm:"type:varchar(250)"`
	beneficiaryResidentStatus string `json:"beneficiaryResidentStatus" gorm:"type:varchar(5)"`
	clearingCode              string `json:"clearingCode" gorm:"type:varchar(200)"`
	remark                    string `json:"remark" gorm:"type:varchar(250)"`
	transactionDate           string `json:"transactionDate" gorm:"type:varchar(20)"`
	transactionTime           string `json:"transactionTime" gorm:"type:varchar(20)"`
	clearingTransactionCode   string `json:"clearingTransactionCode" gorm:"type:varchar(10)"`
	referenceId               string `json:"referenceId" gorm:"type:varchar(250)"`
	paymentDetails1           string `json:"paymentDetails1" gorm:"type:varchar(250)"`
	senderName                string `json:"senderName" gorm:"type:varchar(250)"`
	paymentDetails2           string `json:"paymentDetails2" gorm:"type:varchar(250)"`
	paymentDetails3           string `json:"paymentDetails3" gorm:"type:varchar(250)"`
	debitAccountNo            string `json:"debitAccountNo" gorm:"type:varchar(200)"`
	beneficiaryNationStatus   string `json:"beneficiaryNationStatus" gorm:"type:varchar(5)"`
	beneficiaryType           string `json:"beneficiaryType" gorm:"type:varchar(50)"`
	beneficiaryName           string `json:"beneficiaryName" gorm:"type:varchar(100)"`
	chargeAmount              string `json:"chargeAmount" gorm:"type:varchar(250)"`
	currency                  string `json:"currency" gorm:"type:varchar(50)"`
}

func (*Skn) TableName() string {
	return "skn"
}
