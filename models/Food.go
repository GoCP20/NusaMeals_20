package models

type Food struct {
	ID           int      `gorm:"primaryKey" json:"id_food"`
	CategoryID   int      `json:"id_category"`
	FoodCategory Category `gorm:"foreignKey:CategoryID" json:"food_category"`
	Name         string   `json:"name"`
	Category     string   `json:"category"`
	Country      string   `json:"country"`
	TotalCalorie string   `json:"total_calorie"`
	Description  string   `json:"description"`
	Ingredient   string   `json:"ingredient"`
	Photos       string   `json:"photos"`
}
