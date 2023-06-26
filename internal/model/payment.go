package model

import (
	"github.com/jinzhu/gorm"
)

type Payment struct {
	gorm.Model
	OrderID       uint `json:"order_id" gorm:"column:order_id"`
	Order         Order
	Amount        float64 `json:"amount" gorm:"column:amount"`
	PaymentStatus string  `json:"payment_status" gorm:"column:payment_status"`
}
