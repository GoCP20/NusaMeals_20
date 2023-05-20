package models

import "time"

type Order struct {
	ID            int         `gorm:"primaryKey" json:"id_order"`
	CustomerID    int         `json:"id_customer"`
	Customer      Customer    `gorm:"foreignKey:CustomerID" json:"customer"`
	MenuID        int         `json:"id_menu"`
	Menu          Menu        `gorm:"foreignKey:MenuID" json:"menu"`
	PaymentTypeID int         `json:"id_payment_type"`
	PaymentType   PaymentType `gorm:"foreignKey:PaymentTypeID" json:"payment_type"`
	CustomerName  string      `json:"customer_name"`
	CustomerPhone string      `json:"customer_phone"`
	Type          string      `json:"type"`
	OrderStatus   string      `json:"order_status"`
	OrderDate     time.Time   `json:"order_date"`
	PaymentStatus string      `json:"payment_status"`
	TotalPrice    float64     `json:"total_price"`
	Action        string      `json:"action"`
}
