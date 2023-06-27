package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Reservation struct {
	gorm.Model
	Name           string    `json:"name" gorm:"column:name"`
	PhoneNumber    string    `json:"phone_number" gorm:"column:phone_number"`
	Date           time.Time `json:"date" gorm:"column:date"`
	StartEnd       string    `json:"start_end" gorm:"column:start_end"`
	Agenda         string    `json:"agenda" gorm:"column:agenda"`
	NumberOfPeople string    `json:"number_of_people" gorm:"column:number_of_people"`
}
