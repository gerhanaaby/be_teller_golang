package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model

	ID       uint   `json:"id"`
	ItemId   string `json:"itemId" gorm:"type:varchar(7)"`
	Nama     string `json:"nama" gorm:"type:varchar(50)"`
	Quantity int    `json:"quantity" gorm:"type:integer"`
	OrderID  uint   `json:"orderId"`
}

type Itemjadi struct {
	ID       uint   `json:"id" `
	ItemId   string `json:"itemId" gorm:"type:varchar(7)"`
	Nama     string `json:"nama" gorm:"type:varchar(50)"`
	Quantity int    `json:"quantity" gorm:"type:integer"`
	OrderID  uint   `json:"orderId" `
}

func (*Item) TableName() string {
	return "item"
}

func (*Itemjadi) TableName() string {
	return "item"
}
