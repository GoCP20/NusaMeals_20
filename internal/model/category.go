package model

type Category struct {
	ID    int    `gorm:"primary_key" json:"id" auto_increment:"true"`
	Name  string `json:"name" form:"name"`
	Menus []Menu `json:"menus" gorm:"foreignkey:CategoryID"`
}
