package models

type Reservation struct {
	ID           int    `gorm:"primaryKey" json:"id_reservation"`
	CustomerName string `json:"customer_name"`
	PhoneNumber  string `json:"phone_number"`
	Date         string `json:"date"`
	TimeIn       string `json:"time_in"`
	TimeOut      string `json:"time_out"`
	Action       string `json:"action"`
}
