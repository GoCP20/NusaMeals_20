package response

type GetMenuResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	CategoryID  uint   `json:"category_id"`
	Category    string `json:"category"`
	Calories    string `json:"calories"`
	City        string `json:"city"`
	Description string `json:"description"`
	Ingredients string `json:"ingredients"`
	Images      string `json:"images"`
}

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type GetCategoryResponse struct {
	ID    uint              `json:"id"`
	Name  string            `json:"name"`
	Menus []GetMenuResponse `json:"menus"`
}
