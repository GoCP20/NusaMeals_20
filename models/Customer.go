package models

import "time"

type Customer struct {
	ID               int       `gorm:"primaryKey" json:"id_customer"`
	UserID           int       `json:"id_user"`
	User             User      `gorm:"foreignKey:UserID" json:"user"`
	FullName         string    `json:"full_name"`
	Email            string    `json:"email"`
	Gender           string    `json:"gender"`
	PhoneNumber      string    `json:"phone_number"`
	Address          string    `json:"address"`
	RegistrationDate time.Time `gorm:"autoCreateTime" json:"registration_date"`
}
