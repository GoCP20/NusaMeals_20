package models

type Drink struct {
	ID            int      `gorm:"primaryKey" json:"id_drink"`
	CategoryID    int      `json:"id_category"`
	DrinkCategory Category `gorm:"foreignKey:CategoryID" json:"drink_category"`
	Name          string   `json:"name"`
	Category      string   `json:"category"`
	Country       string   `json:"country"`
	TotalCalorie  string   `json:"total_calorie"`
	Description   string   `json:"description"`
	Ingredient    string   `json:"ingredient"`
	Photos        string   `json:"photos"`
}
