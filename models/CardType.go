package models

type CardType struct {
	ID   int    `gorm:"primaryKey" json:"id_card_type"`
	Name string `json:"name"`
}
