package request

type OrderItemDTO struct {
	ID       uint    `json:"id"`
	OrderID  uint    `json:"order_id"`
	MenuID   uint    `json:"menu_id"`
	Quantity int     `json:"quantity"`
	Subtotal float64 `json:"subtotal"`
}
