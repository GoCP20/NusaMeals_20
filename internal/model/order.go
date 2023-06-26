package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID             uint `gorm:"primaryKey" json:"order_id"`
	UserID         uint `json:"user_id" gorm:"column:user_id"`
	User           User
	MenuID         uint `json:"menu_id" gorm:"column:menu_id"`
	Menu           Menu
	TypeOrder      string         `json:"type_order" gorm:"column:type_order"`
	Quantity       int            `json:"quantity" gorm:"column:quantity"`
	TableNumber    string         `json:"table_number" gorm:"column:table_number"`
	TotalPrice     int            `json:"total_price" gorm:"column:total_price"`
	PaymentMethods string         `json:"payment_methods" gorm:"column:payment_methods"`
	PaymentStatus  string         `json:"payment_status" gorm:"column:payment_status"`
	OrderStatus    string         `json:"order_status" gorm:"column:order_status"`
	CreatedAt      time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
