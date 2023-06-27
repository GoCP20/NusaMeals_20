package usecase

import (
	"reglog/internal/dto/request"
	"reglog/internal/dto/response"
	"reglog/internal/model"
	"reglog/internal/repository"
	"time"
)

type ReservationUseCase interface {
	MakeReservation(reservationRequest *request.ReservationRequest) (*response.CreateReservationResponse, error)
	GetReservationByID(reservationID uint) (*response.CreateReservationResponse, error)
	GetReservationByName(name string) ([]*response.CreateReservationResponse, error)
	UpdateReservation(reservationID uint, updateRequest *request.UpdateReservation) (*response.UpdateReservationResponse, error)
	CancelReservation(reservationID uint) error
	GetAllReservations() ([]*response.CreateReservationResponse, error)
}

type reservationUseCase struct {
	reservationRepo repository.ReservationRepository
}

func NewReservationUseCase(reservationRepo repository.ReservationRepository) ReservationUseCase {
	return &reservationUseCase{
		reservationRepo: reservationRepo,
	}
}

func (uc *reservationUseCase) MakeReservation(reservationRequest *request.ReservationRequest) (*response.CreateReservationResponse, error) {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return nil, err
	}

	currentTime := time.Now().In(location)
	reservation := model.Reservation{
		Name:           reservationRequest.Name,
		PhoneNumber:    reservationRequest.PhoneNumber,
		Date:           currentTime, // Assign time.Time directly
		StartEnd:       reservationRequest.StartEnd,
		Agenda:         reservationRequest.Agenda,
		NumberOfPeople: reservationRequest.NumberOfPeople,
	}

	err = uc.reservationRepo.MakeReservation(&reservation)
	if err != nil {
		return nil, err
	}

	reservationResponse := &response.CreateReservationResponse{
		ID:             reservation.ID,
		Name:           reservation.Name,
		PhoneNumber:    reservation.PhoneNumber,
		Date:           reservation.Date.Format("2006-01-02"), // Format as string (YYYY-MM-DD)
		StartEnd:       reservation.StartEnd,
		Agenda:         reservation.Agenda,
		NumberOfPeople: reservation.NumberOfPeople,
	}

	return reservationResponse, nil
}

func (uc *reservationUseCase) GetReservationByID(reservationID uint) (*response.CreateReservationResponse, error) {
	reservation, err := uc.reservationRepo.GetReservationByID(reservationID)
	if err != nil {
		return nil, err
	}

	reservationResponse := &response.CreateReservationResponse{
		ID:             reservation.ID,
		Name:           reservation.Name,
		PhoneNumber:    reservation.PhoneNumber,
		Date:           reservation.Date.Format("2006-01-02"), // Format as string (YYYY-MM-DD)
		StartEnd:       reservation.StartEnd,
		Agenda:         reservation.Agenda,
		NumberOfPeople: reservation.NumberOfPeople,
	}

	return reservationResponse, nil
}

func (uc *reservationUseCase) GetReservationByName(name string) ([]*response.CreateReservationResponse, error) {
	reservations, err := uc.reservationRepo.GetReservationByName(name)
	if err != nil {
		return nil, err
	}

	var reservationResponses []*response.CreateReservationResponse
	for _, reservation := range reservations {
		reservationResponse := &response.CreateReservationResponse{
			ID:             reservation.ID,
			Name:           reservation.Name,
			PhoneNumber:    reservation.PhoneNumber,
			Date:           reservation.Date.Format("2006-01-02"), // Format as string (YYYY-MM-DD)
			StartEnd:       reservation.StartEnd,
			Agenda:         reservation.Agenda,
			NumberOfPeople: reservation.NumberOfPeople,
			// Add other fields if needed
		}

		reservationResponses = append(reservationResponses, reservationResponse)
	}

	return reservationResponses, nil
}

func (uc *reservationUseCase) UpdateReservation(reservationID uint, updateRequest *request.UpdateReservation) (*response.UpdateReservationResponse, error) {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return nil, err
	}

	reservation, err := uc.reservationRepo.GetReservationByID(reservationID)
	if err != nil {
		return nil, err
	}

	reservation.Name = updateRequest.Name
	reservation.PhoneNumber = updateRequest.PhoneNumber
	reservation.Date = time.Now().In(location) // Assign time.Time directly
	reservation.StartEnd = updateRequest.StartEnd
	reservation.Agenda = updateRequest.Agenda
	reservation.NumberOfPeople = updateRequest.NumberOfPeople

	err = uc.reservationRepo.UpdateReservation(reservation)
	if err != nil {
		return nil, err
	}

	reservationResponse := &response.UpdateReservationResponse{
		ID:             reservation.ID,
		Name:           reservation.Name,
		PhoneNumber:    reservation.PhoneNumber,
		Date:           reservation.Date.Format("2006-01-02"), // Format as string (YYYY-MM-DD)
		StartEnd:       reservation.StartEnd,
		Agenda:         reservation.Agenda,
		NumberOfPeople: reservation.NumberOfPeople,
		// Add other fields if needed
	}

	return reservationResponse, nil
}

func (uc *reservationUseCase) CancelReservation(reservationID uint) error {
	err := uc.reservationRepo.CancelReservation(reservationID)
	return err
}

func (uc *reservationUseCase) GetAllReservations() ([]*response.CreateReservationResponse, error) {
	reservations, err := uc.reservationRepo.GetAllReservations()
	if err != nil {
		return nil, err
	}

	var reservationResponses []*response.CreateReservationResponse
	for _, reservation := range reservations {
		reservationResponse := &response.CreateReservationResponse{
			ID:             reservation.ID,
			Name:           reservation.Name,
			PhoneNumber:    reservation.PhoneNumber,
			Date:           reservation.Date.Format("2006-01-02"), // Format as string (YYYY-MM-DD)
			StartEnd:       reservation.StartEnd,
			Agenda:         reservation.Agenda,
			NumberOfPeople: reservation.NumberOfPeople,
			// Add other fields if needed
		}

		reservationResponses = append(reservationResponses, reservationResponse)
	}

	return reservationResponses, nil
}
