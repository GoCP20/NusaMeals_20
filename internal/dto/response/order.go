package response

type GetOrderResponse struct {
	ID             uint   `json:"id"`
	UserID         uint   `json:"user_id"`
	MenuID         uint   `json:"menu_id"`
	MenuName       string `json:"menu_name"`
	MenuImages     string `json:"menu_images"`
	MenuCity       string `json:"menu_city"`
	MenuCalories   string `json:"menu_calories"`
	PriceMenu      int    `json:"price_menu"`
	Quantity       int    `json:"quantity"`
	TypeOrder      string `json:"type_order"`
	TableNumber    string `json:"table_number"`
	TotalPrice     int    `json:"total_price"`
	PaymentMethods string `json:"payment_methods"`
	OrderStatus    string `json:"order_status"`
	CreatedAt      string `json:"created_at"`
}

type GetOrderDetails struct {
	ID             uint   `json:"id"`
	UserID         uint   `json:"user_id"`
	Username       string `json:"username"`
	MenuID         uint   `json:"menu_id"`
	MenuName       string `json:"menu_name"`
	MenuImages     string `json:"menu_images"`
	MenuCity       string `json:"menu_city"`
	MenuCalories   string `json:"menu_calories"`
	PriceMenu      int    `json:"price_menu"`
	Quantity       int    `json:"quantity"`
	TypeOrder      string `json:"type_order"`
	TableNumber    string `json:"table_number"`
	TotalPrice     int    `json:"total_price"`
	PaymentMethods string `json:"payment_methods"`
	OrderStatus    string `json:"order_status"`
	CreatedAt      string `json:"created_at"`
}

type OrderUpdate struct {
	ID          uint   `json:"id"`
	UserID      uint   `json:"user_id"`
	MenuID      uint   `json:"menu_id"`
	TypeOrder   string `json:"type_order"`
	OrderStatus string `json:"order_status"`
	TotalPrice  int    `json:"total_price"`
	UpdatedAt   string `json:"updated_at"`
	CreatedAt   string `json:"created_at"`
}

type GetAllOrdersResponse struct {
	ID          uint   `json:"id"`
	UserID      uint   `json:"user_id"`
	MenuID      uint   `json:"menu_id"`
	Quantity    int    `json:"quantity"`
	TypeOrder   string `json:"type_order"`
	TotalPrice  int    `json:"total_price"`
	OrderStatus string `json:"order_status"`
	CreatedAt   string `json:"created_at"`
}
