package models

type EwalletType struct {
	ID   int    `gorm:"primaryKey" json:"id_ewallet_type"`
	Name string `json:"name"`
}
