package response

type GetItem struct {
	ID             uint   `json:"id"`
	Name           string `json:"name" form:"name"`
	Description    string `json:"description" form:"description"`
	Stock          int    `json:"stock" form:"stock"`
	Price          int    `json:"price" form:"price"`
	CategoryTypeID uint   `json:"categoryType_id" form:"categoryType_id"`
	CategoryType   string `json:"categoryType" form:"categoryType"`
}
