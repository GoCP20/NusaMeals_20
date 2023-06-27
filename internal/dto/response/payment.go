package response

type Payment struct {
	ID             uint    `json:"id"`
	OrderID        uint    `json:"order_id"`
	UserID         uint    `json:"user_id"`
	Username       string  `json:"username"`
	TotalPrice     int     `json:"total_price"`
	TypeOrder      string  `json:"type_order"`
	TableNumber    string  `json:"table_number"`
	PaymentMethods string  `json:"payment_methods"`
	Amount         float64 `json:"amount"`
	PaymentStatus  string  `json:"payment_status"`
}

type PaymentUpdate struct {
	ID             uint    `json:"id"`
	OrderID        uint    `json:"order_id"`
	UserID         uint    `json:"user_id"`
	Username       string  `json:"username"`
	TotalPrice     int     `json:"total_price"`
	TypeOrder      string  `json:"type_order"`
	TableNumber    string  `json:"table_number"`
	PaymentMethods string  `json:"payment_methods"`
	Amount         float64 `json:"amount"`
	PaymentStatus  string  `json:"payment_status"`
}

type UpdatePayment struct {
	ID             uint    `json:"id"`
	OrderID        uint    `json:"order_id"`
	UserID         uint    `json:"user_id"`
	Username       string  `json:"username"`
	TotalPrice     int     `json:"total_price"`
	TypeOrder      string  `json:"type_order"`
	TableNumber    string  `json:"table_number"`
	PaymentMethods string  `json:"payment_methods"`
	Amount         float64 `json:"amount"`
	PaymentStatus  string  `json:"payment_status"`
}

type GetAllPayment struct {
	ID            uint    `json:"id"`
	OrderID       uint    `json:"order_id"`
	Amount        float64 `json:"amount"`
	PaymentStatus string  `json:"payment_status"`
}
