package models

type OrderCancel struct {
	ID         int      `gorm:"primaryKey" json:"id_order_cancel"`
	OrderID    int      `json:"id_order"`
	Order      Order    `gorm:"foreignKey:OrderID" json:"order"`
	CustomerID int      `json:"id_customer"`
	Customer   Customer `gorm:"foreignKey:CustomerID" json:"customer"`
	Reason     string   `json:"reason"`
}
