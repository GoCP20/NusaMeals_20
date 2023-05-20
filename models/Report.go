package models

import "time"

type Report struct {
	ID          int       `gorm:"primaryKey" json:"id_report"`
	Cash        float64   `json:"cash"`
	DebitCard   float64   `json:"debit_card"`
	EWallet     float64   `json:"e_wallet"`
	TotalAmount float64   `json:"total_amount"`
	Date        time.Time `json:"date"`
}
