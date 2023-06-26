package request

type Payment struct {
	ID             uint    `json:"id"`
	OrderID        uint    `json:"order_id"`
	UserID         uint    `json:"user_id"`
	Username       string  `json:"username"`
	MenuName       string  `json:"menu_name"`
	TotalPrice     int     `json:"total_price"`
	PaymentMethods string  `json:"payment_methods"`
	Amount         float64 `json:"amount"`
	Status         string  `json:"status"`
	PaymentType    string  `json:"payment_type"`
}

type PaymentUpdate struct {
	OrderID       uint    `json:"order_id"`
	UserID        uint    `json:"user_id"`
	Amount        float64 `json:"amount"`
	PaymentStatus string  `json:"payment_status"`
	Method        string  `json:"method"`
	PaymentType   string  `json:"payment_type"`
}
