package data

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"reglog/internal/model"
)

func UserFaker(db *gorm.DB) *model.User {
	password, _ := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	return &model.User{
		Name:     "admin",
		Username: "admin",
		Email:    "admin",
		Password: string(password),
		Role:     "admin",
	}
}
