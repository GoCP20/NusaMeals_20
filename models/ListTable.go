package models

type ListTable struct {
	ID       int    `gorm:"primaryKey" json:"id_list_table"`
	Number   int    `json:"number_of_table"`
	Seat     int    `json:"seat"`
	Type     string `json:"type"`
	Status   string `json:"status"`
	Location string `json:"location"`
	Photos   string `json:"photos"`
}
