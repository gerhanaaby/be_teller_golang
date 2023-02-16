package models

import "gorm.io/gorm"

type InternalTransfer struct {
	gorm.Model

	ID               uint   `json:"id"`
	ReferenceId      string `json:"referenceId" gorm:"type:varchar(200)"`
	DebitAccountNo   string `json:"debitAccountNo" gorm:"type:varchar(200)"`
	CreditAccountNo  string `json:"creditAccountNo" gorm:"type:varchar(200)"`
	CreditAmount     string `json:"creditAmount" goarm:"type:varchar(250)"`
	CreditCurrency   string `json:"creditCurrency" gorm:"type:varchar(10)"`
	TransactionDate  string `json:"transactionDate" gorm:"type:varchar(20)"`
	Remark           string `json:"remark" gorm:"type:varchar(250)"`
	BeneficiaryName  string `json:"beneficiaryName" gorm:"type:varchar(250)"`
	DebitAccountName string `json:"debitAccountName" gorm:"type:varchar(250)"`
}

func (*InternalTransfer) TableName() string {
	return "internaltransfer"
}
