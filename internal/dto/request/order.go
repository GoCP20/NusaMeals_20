package request

type UpdateOrder struct {
	OrderStatus string `json:"order_status"`
}

type CreateOrderRequest struct {
	Username  string `json:"username"`
	MenuName  string `json:"menu_name"`
	PriceMenu int    `json:"price_menu"`
	TypeOrder string `json:"type_order"`
	Quantity  int    `json:"quantity"`
}

type CreateOrder struct {
	ID             uint   `json:"id"`
	UserID         uint   `json:"user_id"`
	Username       string `json:"username"`
	MenuID         uint   `json:"menu_id"`
	MenuName       string `json:"menu_name"`
	Quantity       int    `json:"quantity"`
	TableNumber    string `json:"table_number"`
	TotalPrice     int    `json:"total_price"`
	TypeOrder      string `json:"type_order"`
	PaymentMethods string `json:"payment_methods"`
	OrderStatus    string `json:"order_status"`
}
