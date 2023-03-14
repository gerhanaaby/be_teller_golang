package models

import "gorm.io/gorm"

type InternalTransfer struct {
	gorm.Model

	// ID               uint   `json:"id"`
	CaseID 			 string `json:"caseID" gorm:"type:varchar(50);default:'PB-'||TO_CHAR(NOW()::date, 'yyyymmdd')||to_char(nextval('itf_tix_seq'::regclass), 'FM0000000000')"`
	ReferenceId      string `json:"referenceId" gorm:"type:varchar(250);uniqueIndex"`
	// ReferenceId      string `json:"referenceId" gorm:"type:varchar(250)"`
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
