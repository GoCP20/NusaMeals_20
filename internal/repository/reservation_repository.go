package repository

import (
	"reglog/internal/model"

	"github.com/jinzhu/gorm"
)

type ReservationRepository interface {
	MakeReservation(data *model.Reservation) error
	GetReservationByID(ID uint) (*model.Reservation, error)
	GetReservationByName(name string) ([]*model.Reservation, error)
	GetAllReservations() ([]*model.Reservation, error)
	UpdateReservation(data *model.Reservation) error
	CancelReservation(ID uint) error
}

type reservationRepository struct {
	db *gorm.DB
}

func NewReservationRepository(db *gorm.DB) ReservationRepository {
	return &reservationRepository{db: db}
}

func (r *reservationRepository) MakeReservation(data *model.Reservation) error {
	err := r.db.Create(data).Error
	return err
}

func (r *reservationRepository) GetAllReservations() ([]*model.Reservation, error) {
	var reservations []*model.Reservation
	if err := r.db.Find(&reservations).Error; err != nil {
		return nil, err
	}
	return reservations, nil
}

func (r *reservationRepository) GetReservationByID(ID uint) (*model.Reservation, error) {
	reservation := &model.Reservation{}
	err := r.db.First(reservation, ID).Error
	return reservation, err
}

func (r *reservationRepository) GetReservationByName(name string) ([]*model.Reservation, error) {
	var reservations []*model.Reservation
	err := r.db.Where("name = ?", name).Find(&reservations).Error
	return reservations, err
}

func (r *reservationRepository) UpdateReservation(data *model.Reservation) error {
	err := r.db.Save(data).Error
	return err
}

func (r *reservationRepository) CancelReservation(ID uint) error {
	err := r.db.Delete(&model.Reservation{}, ID).Error
	return err
}
