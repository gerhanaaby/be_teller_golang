package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model

	OrderNo     int       `json:"orderNo" gorm:"type:bigint"`
	DateOrder   time.Time `json:"dateOrder"`
	TrackingNo  string    `json:"trackingNo" gorm:"type:varchar(8)"`
	Customers   Customer  `gorm:"foreignKey:CustomerID"`
	TotalPrice  int       `json:"totalPrice" gorm:"type:integer"`
	Items       []Item
	DiscountTag string `json:"discountTag"`
	CustomerID  uint
}

type Orderjadi struct {
	ID          uint
	OrderNo     int          `json:"orderNo" gorm:"type:bigint"`
	DateOrder   time.Time    `json:"dateOrder"`
	TrackingNo  string       `json:"trackingNo" gorm:"type:varchar(8)"`
	Customers   Customerjadi `gorm:"foreignKey:CustomerID"`
	TotalPrice  int          `json:"totalPrice" gorm:"type:integer"`
	Items       []Itemjadi   `gorm:"foreignKey:OrderID"`
	DiscountTag string       `json:"discountTag"`
	CustomerID  uint
}

func (*Order) TableName() string {
	return "order"
}

func (*Orderjadi) TableName() string {
	return "order"
}
