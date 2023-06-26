package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name        string `gorm:"not null" json:"name" form:"name"`
	Username    string `gorm:"unique;not null" json:"username" form:"username"`
	Email       string `gorm:"unique;not null" json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	Gender      string `json:"gender" form:"gender"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Picture     string `gorm:"null" json:"picture" form:"picture"`
	Role        string `gorm:"type:enum('user','admin');default:'user'" json:"role" form:"role"`
}
