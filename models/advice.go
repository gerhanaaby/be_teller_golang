package models

import "gorm.io/gorm"

type Advice struct {
	gorm.Model

	ID                  uint   `json:"id"`
	ReferenceId         string `json:"referenceId" gorm:"type:varchar(200)"`
	DebitAccountNo      string `json:"debitAccountNo" gorm:"type:varchar(200)"`
	CreditAccountNo     string `json:"creditAccountNo" gorm:"type:varchar(200)"`
	TransactionDate     string `json:"transactionDate" gorm:"type:varchar(20)"`
	OriginalReferenceId string `json:"originalReferenceId" gorm:"type:varchar(250)"`
	ClientId            string `json:"clientId" gorm:"type:varchar(250)"`
	TransactionTime     string `json:"transactionTime" gorm:"type:varchar(20)"`
}

func (*Advice) TableName() string {
	return "advice"
}
