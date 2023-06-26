package model

import (
	"github.com/jinzhu/gorm"
)

type Menu struct {
	gorm.Model
	Name        string `json:"name" form:"name"`
	Price       int    `json:"price" form:"price"`
	Calories    string `json:"calories" form:"calories"`
	City        string `json:"city" form:"city"`
	Description string `json:"description" form:"description"`
	Ingredient  string `json:"ingredient" form:"ingredient"`
	Images      string `json:"images" form:"images"`
	CategoryID  uint   `json:"category_id" form:"category_id"`
	Category    Category
}
