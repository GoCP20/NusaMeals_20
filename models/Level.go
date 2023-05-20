package models

type Level struct {
	ID   int    `gorm:"primaryKey" json:"id_level"`
	Name string `json:"level_name"`
}
