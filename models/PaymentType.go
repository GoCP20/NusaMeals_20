package models

type PaymentType struct {
	ID   int    `gorm:"primaryKey" json:"id_payment_type"`
	Name string `json:"name"`
}
