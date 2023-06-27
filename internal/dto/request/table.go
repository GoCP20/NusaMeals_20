package request

type CreateTable struct {
	NumberTable   int    `json:"number_of_table"`
	Seats         int    `json:"seats"`
	Type          string `json:"type"`
	TableLocation string `json:"table_location"`
	Picture       string `json:"picture"`
}

type UpdateTable struct {
	NumberTable   int    `json:"number_of_table"`
	Seats         int    `json:"seats"`
	Type          string `json:"type"`
	TableLocation string `json:"table_location"`
	Picture       string `json:"picture"`
}
