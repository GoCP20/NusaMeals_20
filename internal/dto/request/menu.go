package request

type Menu struct {
	Name        string `json:"name" form:"name" binding:"required"`
	Price       int    `json:"price" form:"price" binding:"required"`
	Calories    string `json:"calories" form:"calories"`
	City        string `json:"city" form:"city"`
	Description string `json:"description" form:"description"`
	Ingredient  string `json:"ingredient" form:"ingredient"`
	Images      string `json:"images" form:"images"`
	CategoryID  uint   `json:"category_id" form:"category_id" validate:"required"`
}

type UpdateMenu struct {
	Name        string `json:"name" form:"name" binding:"required"`
	Price       int    `json:"price" form:"price" binding:"required"`
	Calories    string `json:"calories" form:"calories"`
	City        string `json:"city" form:"city"`
	Description string `json:"description" form:"description"`
	Ingredient  string `json:"ingredient" form:"ingredient"`
	Images      string `json:"images" form:"images"`
	CategoryID  uint   `json:"category_id" form:"category_id" validate:"required"`
}
