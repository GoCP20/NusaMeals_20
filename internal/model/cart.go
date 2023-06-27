package model

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	ID        uint           `gorm:"primaryKey" json:"cart_id"`
	UserID    uint           `gorm:"unique" json:"user_id"`
	Items     []CartItem     `json:"items" gorm:"foreignkey:CartID"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type CartItem struct {
	ID        uint    `gorm:"primaryKey" json:"item_id"`
	CartID    uint    `gorm:"index" json:"cart_id"`
	MenuID    uint    `gorm:"index" json:"menu_id"`
	Menu      Menu    `json:"menu"`
	Quantity  int     `json:"quantity" gorm:"column:quantity"`
	Subtotal  float64 `json:"subtotal" gorm:"column:subtotal"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
