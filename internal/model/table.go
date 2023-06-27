package model

import (
	"github.com/jinzhu/gorm"
)

type Table struct {
	gorm.Model
	NumberTable   int    `json:"number_of_table" gorm:"column:number_of_table"`
	Seats         int    `json:"seats" gorm:"column:seats"`
	Type          string `json:"type" gorm:"column:type"`
	TableLocation string `json:"table_location" gorm:"column:table_location"`
	Picture       string `json:"picture" gorm:"column:picture"`
}
