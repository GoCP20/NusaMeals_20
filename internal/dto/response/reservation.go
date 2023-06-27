package response

type CreateReservationResponse struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	PhoneNumber    string `json:"phone_number"`
	Date           string `json:"date"`
	StartEnd       string `json:"start_end"`
	Agenda         string `json:"agenda"`
	NumberOfPeople string `json:"number_of_people"`
}

type UpdateReservationResponse struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	PhoneNumber    string `json:"phone_number"`
	Date           string `json:"date"`
	StartEnd       string `json:"start_end"`
	Agenda         string `json:"agenda"`
	NumberOfPeople string `json:"number_of_people"`
}
