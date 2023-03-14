package models

import "gorm.io/gorm"

type GetDetail struct {
	gorm.Model

	// ID            uint   `json:"id"`

	TransactionID string `json:"transactionID" gorm:"type:varchar(50)"`
	AccountNumber string `json:"accountNumber" gorm:"type:varchar(50)"`
	TransactCategory string `json:"transactCategory" gorm:"type:varchar(50)"`
	// CaseID string `json:"caseID" gorm:"type:varchar(50)"`
	// SKN Skn `gorm:"references:reference_id"`
	// InternalTransfer InternalTransfer `gorm:"references:TransactionID"`

}

func (*GetDetail) TableName() string {
	return "getdetail"
}
