package repository

import (
	"reglog/internal/model"

	"github.com/jinzhu/gorm"
)

type PaymentRepository interface {
	CreatePayment(payment *model.Payment) (*model.Payment, error)
	UpdatePayment(payment *model.Payment) (*model.Payment, error)
	DeletePayment(paymentID uint) error
	GetPaymentByID(paymentID uint) (*model.Payment, error)
	GetAllPayments() ([]model.Payment, error)
	GetPaymentByOrderID(orderID uint) ([]model.Payment, error)
	GetPaymentByUsername(username string) ([]model.Payment, error)
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{
		db: db,
	}
}

func (r *paymentRepository) CreatePayment(payment *model.Payment) (*model.Payment, error) {
	result := r.db.Create(payment)
	if result.Error != nil {
		return nil, result.Error
	}
	return payment, nil
}
func (r *paymentRepository) GetPaymentByOrderID(orderID uint) ([]model.Payment, error) {
	var payments []model.Payment
	err := r.db.Where("order_id = ?", orderID).Find(&payments).Error
	if err != nil {
		return nil, err
	}
	return payments, nil
}

func (r *paymentRepository) GetPaymentByUsername(username string) ([]model.Payment, error) {
	var payments []model.Payment
	err := r.db.Table("payments").
		Select("payments.*").
		Joins("JOIN orders ON payments.order_id = orders.id").
		Joins("JOIN users ON orders.user_id = users.id").
		Where("users.username = ?", username).
		Find(&payments).Error
	if err != nil {
		return nil, err
	}
	return payments, nil
}

func (r *paymentRepository) UpdatePayment(payment *model.Payment) (*model.Payment, error) {
	result := r.db.Save(payment)
	if result.Error != nil {
		return nil, result.Error
	}
	return payment, nil
}

func (r *paymentRepository) DeletePayment(paymentID uint) error {
	return r.db.Delete(&model.Payment{}, paymentID).Error
}

func (r *paymentRepository) GetPaymentByID(paymentID uint) (*model.Payment, error) {
	var payment model.Payment
	err := r.db.First(&payment, paymentID).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *paymentRepository) GetAllPayments() ([]model.Payment, error) {
	var payments []model.Payment
	err := r.db.Find(&payments).Error
	if err != nil {
		return nil, err
	}
	return payments, nil
}
