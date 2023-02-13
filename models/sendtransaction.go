package models

import "gorm.io/gorm"

type STransaction struct {
	gorm.Model

	ID       uint   `json:"id"`
	ItemId   string `json:"itemId" gorm:"type:varchar(7)"`
	Nama     string `json:"nama" gorm:"type:varchar(50)"`
	Quantity int    `json:"quantity" gorm:"type:integer"`
	OrderID  uint   `json:"orderId"`
}

type STransactionjadi struct {
	ID       uint   `json:"id" `
	ItemId   string `json:"itemId" gorm:"type:varchar(7)"`
	Nama     string `json:"nama" gorm:"type:varchar(50)"`
	Quantity int    `json:"quantity" gorm:"type:integer"`
	OrderID  uint   `json:"orderId" `
}

func (*STransaction) TableName() string {
	return "sendtransaction"
}

func (*STransactionjadi) TableName() string {
	return "sendtransaction"
}
