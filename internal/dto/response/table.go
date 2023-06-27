package response

type CreateTableResponse struct {
	ID            uint   `json:"id"`
	NumberTable   int    `json:"number_of_table"`
	Seats         int    `json:"seats"`
	Type          string `json:"type"`
	TableLocation string `json:"table_location"`
	Picture       string `json:"picture"`
}

type UpdateTableResponse struct {
	ID            uint   `json:"id"`
	NumberTable   int    `json:"number_of_table"`
	Seats         int    `json:"seats"`
	Type          string `json:"type"`
	TableLocation string `json:"table_location"`
	Picture       string `json:"picture"`
}
