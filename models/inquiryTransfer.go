package models

import "gorm.io/gorm"

type InquiryTransfer struct {
	gorm.Model

	ID          uint   `json:"id"`
	AccountNo   string `json:"accountNo" gorm:"type:varchar(250)"`
	ReferenceId string `json:"referenceId" gorm:"type:varchar(250)"`
}

func (*InquiryTransfer) TableName() string {
	return "inquiryTransfer"
}
