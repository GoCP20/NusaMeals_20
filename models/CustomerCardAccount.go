package models

type CustomerCardAccount struct {
	ID                int      `gorm:"primaryKey" json:"id_card"`
	CardTypeID        int      `json:"id_card_type"`
	CardType          CardType `gorm:"foreignKey:CardTypeID" json:"card_type"`
	CustomerID        int      `json:"id_customer"`
	Customer          Customer `gorm:"foreignKey:CustomerID" json:"customer"`
	CardAccountName   string   `json:"card_account_name"`
	CardAccountNumber string   `json:"card_account_number"`
}
