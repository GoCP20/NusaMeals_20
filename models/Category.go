package models

import "time"

type Category struct {
	ID        int       `gorm:"primaryKey" json:"id_category"`
	Name      string    `json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
