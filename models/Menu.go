package models

import "time"

type Menu struct {
	ID         int       `gorm:"primaryKey" json:"id_menu"`
	CategoryID int       `json:"id_category"`
	Category   Category  `gorm:"foreignKey:CategoryID" json:"category"`
	Name       string    `json:"name"`
	Price      float64   `json:"price"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
