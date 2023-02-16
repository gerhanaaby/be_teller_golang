package models

import "gorm.io/gorm"

type Skn struct {
	gorm.Model

	ID                        uint   `json:"id"`
	CreditAccountNo           string `json:"creditAccountNo" gorm:"type:varchar(50)"`
	Amount                    string `json:"amount" gorm:"type:varchar(250)"`
	BeneficiaryResidentStatus string `json:"beneficiaryResidentStatus" gorm:"type:varchar(5)"`
	ClearingCode              string `json:"clearingCode" gorm:"type:varchar(200)"`
	Remark                    string `json:"remark" gorm:"type:varchar(250)"`
	TransactionDate           string `json:"transactionDate" gorm:"type:varchar(20)"`
	TransactionTime           string `json:"transactionTime" gorm:"type:varchar(20)"`
	ClearingTransactionCode   string `json:"clearingTransactionCode" gorm:"type:varchar(10)"`
	ReferenceId               string `json:"referenceId" gorm:"type:varchar(250)"`
	PaymentDetails1           string `json:"paymentDetails1" gorm:"type:varchar(250)"`
	SenderName                string `json:"senderName" gorm:"type:varchar(250)"`
	PaymentDetails2           string `json:"paymentDetails2" gorm:"type:varchar(250)"`
	PaymentDetails3           string `json:"paymentDetails3" gorm:"type:varchar(250)"`
	DebitAccountNo            string `json:"debitAccountNo" gorm:"type:varchar(200)"`
	BeneficiaryNationStatus   string `json:"beneficiaryNationStatus" gorm:"type:varchar(5)"`
	BeneficiaryType           string `json:"beneficiaryType" gorm:"type:varchar(50)"`
	BeneficiaryName           string `json:"beneficiaryName" gorm:"type:varchar(100)"`
	ChargeAmount              string `json:"chargeAmount" gorm:"type:varchar(250)"`
	Currency                  string `json:"currency" gorm:"type:varchar(50)"`
}

func (*Skn) TableName() string {
	return "skn"
}
