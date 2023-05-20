package models

import "time"

type Payment struct {
	ID             int         `gorm:"primaryKey" json:"id_payment"`
	OrderID        int         `json:"id_order"`
	Order          Order       `gorm:"foreignKey:OrderID" json:"order"`
	PaymentTypeID  int         `json:"id_payment_type"`
	PaymentType    PaymentType `gorm:"foreignKey:PaymentTypeID" json:"payment_type"`
	TotalAmount    float64     `json:"total_amount"`
	PaymentDate    time.Time   `json:"payment_date"`
	PaymentStatus  string      `json:"payment_status"`
	ApprovalStatus string      `json:"approval_status"`
}
