package models

type CustomerEwalletAccount struct {
	ID                   int         `gorm:"primaryKey" json:"id_ewallet"`
	EwalletTypeID        int         `json:"id_ewallet_type"`
	EwalletType          EwalletType `gorm:"foreignKey:EwalletTypeID" json:"ewallet_type"`
	CustomerID           int         `json:"id_customer"`
	Customer             Customer    `gorm:"foreignKey:CustomerID" json:"customer"`
	EwalletAccountName   string      `json:"ewallet_account_name"`
	EwalletAccountNumber string      `json:"ewallet_account_number"`
}
