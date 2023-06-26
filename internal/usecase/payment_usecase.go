package usecase

import (
	"reglog/internal/dto/request"
	"reglog/internal/dto/response"
	"reglog/internal/model"
	"reglog/internal/repository"
)

type PaymentUseCase interface {
	CreatePayment(paymentRequest *request.Payment) (*response.Payment, error)
	UpdatePayment(paymentID uint, paymentRequest *request.Payment) (*response.Payment, error)
	UpdatePaymentByAdmin(paymentID uint, paymentRequest *request.PaymentUpdate) (*response.PaymentUpdate, error)
	DeletePayment(paymentID uint) error
	GetPaymentByID(paymentID uint) (*response.Payment, error)
	GetAllPayments() ([]response.GetAllPayment, error)
	GetPaymentByOrderID(orderID uint) ([]response.Payment, error)
	GetPaymentByUsername(username string) ([]response.Payment, error)
}

type paymentUseCase struct {
	paymentRepo repository.PaymentRepository
	orderRepo   repository.OrderRepository
	userRepo    repository.UserRepository
}

func NewPaymentUseCase(paymentRepo repository.PaymentRepository, orderRepo repository.OrderRepository, userRepo repository.UserRepository) PaymentUseCase {
	return &paymentUseCase{
		paymentRepo: paymentRepo,
		orderRepo:   orderRepo,
		userRepo:    userRepo,
	}
}

func (u *paymentUseCase) CreatePayment(paymentRequest *request.Payment) (*response.Payment, error) {
	order, err := u.orderRepo.GetOrderByID(paymentRequest.OrderID)
	if err != nil {
		return nil, err
	}

	user, err := u.userRepo.GetUserByID(order.UserID)
	if err != nil {
		return nil, err
	}

	payment := &model.Payment{
		OrderID:       paymentRequest.OrderID,
		Amount:        paymentRequest.Amount,
		PaymentStatus: "Not Yet Paid",
	}

	createdPayment, err := u.paymentRepo.CreatePayment(payment)
	if err != nil {
		return nil, err
	}

	createdPaymentResponse := &response.Payment{
		ID:             createdPayment.ID,
		OrderID:        createdPayment.OrderID,
		UserID:         order.UserID,
		Username:       user.Username,
		TypeOrder:      order.TypeOrder,
		TableNumber:    order.TableNumber,
		TotalPrice:     order.TotalPrice,
		PaymentMethods: order.PaymentMethods,
		Amount:         createdPayment.Amount,
		PaymentStatus:  createdPayment.PaymentStatus,
	}

	return createdPaymentResponse, nil
}

func (u *paymentUseCase) UpdatePayment(paymentID uint, paymentRequest *request.Payment) (*response.Payment, error) {
	payment, err := u.paymentRepo.GetPaymentByID(paymentID)
	if err != nil {
		return nil, err
	}

	payment.Amount = paymentRequest.Amount

	updatedPayment, err := u.paymentRepo.UpdatePayment(payment)
	if err != nil {
		return nil, err
	}

	order, err := u.orderRepo.GetOrderByID(updatedPayment.OrderID)
	if err != nil {
		return nil, err
	}

	updatedPaymentResponse := &response.Payment{
		ID:             updatedPayment.ID,
		OrderID:        updatedPayment.OrderID,
		UserID:         order.UserID,
		Username:       order.User.Username,
		TypeOrder:      order.TypeOrder,
		TableNumber:    order.TableNumber,
		TotalPrice:     order.TotalPrice,
		PaymentMethods: order.PaymentMethods,
		Amount:         updatedPayment.Amount,
		PaymentStatus:  updatedPayment.PaymentStatus,
	}

	return updatedPaymentResponse, nil
}

func (u *paymentUseCase) UpdatePaymentByAdmin(paymentID uint, paymentRequest *request.PaymentUpdate) (*response.PaymentUpdate, error) {
	payment, err := u.paymentRepo.GetPaymentByID(paymentID)
	if err != nil {
		return nil, err
	}

	payment.PaymentStatus = paymentRequest.PaymentStatus
	// Update other fields accordingly

	updatedPayment, err := u.paymentRepo.UpdatePayment(payment)
	if err != nil {
		return nil, err
	}

	order, err := u.orderRepo.GetOrderByID(updatedPayment.OrderID)
	if err != nil {
		return nil, err
	}

	user, err := u.userRepo.GetUserByID(order.UserID)
	if err != nil {
		return nil, err
	}

	updatedPaymentResponse := &response.PaymentUpdate{
		ID:             updatedPayment.ID,
		OrderID:        updatedPayment.OrderID,
		UserID:         user.ID,
		Username:       user.Username,
		TotalPrice:     order.TotalPrice,
		TypeOrder:      order.TypeOrder,
		TableNumber:    order.TableNumber,
		PaymentMethods: order.PaymentMethods,
		Amount:         updatedPayment.Amount,
		PaymentStatus:  updatedPayment.PaymentStatus,
		// Set other fields accordingly
	}

	return updatedPaymentResponse, nil
}

func (u *paymentUseCase) DeletePayment(paymentID uint) error {
	return u.paymentRepo.DeletePayment(paymentID)
}

func (u *paymentUseCase) GetPaymentByID(paymentID uint) (*response.Payment, error) {
	payment, err := u.paymentRepo.GetPaymentByID(paymentID)
	if err != nil {
		return nil, err
	}

	order, err := u.orderRepo.GetOrderByID(payment.OrderID)
	if err != nil {
		return nil, err
	}

	user, err := u.userRepo.GetUserByID(order.UserID)
	if err != nil {
		return nil, err
	}

	paymentResponse := &response.Payment{
		ID:             payment.ID,
		OrderID:        payment.OrderID,
		UserID:         order.UserID,
		Username:       user.Username,
		TotalPrice:     order.TotalPrice,
		TypeOrder:      order.TypeOrder,
		TableNumber:    order.TableNumber,
		PaymentMethods: order.PaymentMethods,
		Amount:         payment.Amount,
		PaymentStatus:  payment.PaymentStatus,
		// Set other fields accordingly
	}

	return paymentResponse, nil
}

func (u *paymentUseCase) GetAllPayments() ([]response.GetAllPayment, error) {
	payments, err := u.paymentRepo.GetAllPayments()
	if err != nil {
		return nil, err
	}

	var paymentResponses []response.GetAllPayment
	for _, payment := range payments {
		paymentResponse := response.GetAllPayment{
			ID:            payment.ID,
			OrderID:       payment.OrderID,
			Amount:        payment.Amount,
			PaymentStatus: payment.PaymentStatus,
		}

		paymentResponses = append(paymentResponses, paymentResponse)
	}

	return paymentResponses, nil
}

func (u *paymentUseCase) GetPaymentByOrderID(orderID uint) ([]response.Payment, error) {
	allPayments, err := u.paymentRepo.GetAllPayments()
	if err != nil {
		return nil, err
	}

	var paymentResponses []response.Payment
	for _, payment := range allPayments {
		if payment.OrderID == orderID {
			order, err := u.orderRepo.GetOrderByID(payment.OrderID)
			if err != nil {
				return nil, err
			}

			user, err := u.userRepo.GetUserByID(order.UserID)
			if err != nil {
				return nil, err
			}

			paymentResponse := response.Payment{
				ID:             payment.ID,
				OrderID:        payment.OrderID,
				UserID:         order.UserID,
				Username:       user.Username,
				TotalPrice:     order.TotalPrice,
				TypeOrder:      order.TypeOrder,
				TableNumber:    order.TableNumber,
				PaymentMethods: order.PaymentMethods,
				Amount:         payment.Amount,
				PaymentStatus:  payment.PaymentStatus,
				// Set other fields accordingly
			}

			paymentResponses = append(paymentResponses, paymentResponse)
		}
	}

	return paymentResponses, nil
}

func (u *paymentUseCase) GetPaymentByUsername(username string) ([]response.Payment, error) {
	user, err := u.userRepo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	payments, err := u.paymentRepo.GetPaymentByUsername(username)
	if err != nil {
		return nil, err
	}

	var paymentResponses []response.Payment
	for _, payment := range payments {
		order, err := u.orderRepo.GetOrderByID(payment.OrderID)
		if err != nil {
			return nil, err
		}

		paymentResponse := response.Payment{
			ID:             payment.ID,
			OrderID:        payment.OrderID,
			UserID:         order.UserID,
			Username:       user.Username,
			TotalPrice:     order.TotalPrice,
			TypeOrder:      order.TypeOrder,
			TableNumber:    order.TableNumber,
			PaymentMethods: order.PaymentMethods,
			Amount:         payment.Amount,
			PaymentStatus:  payment.PaymentStatus,
		}

		paymentResponses = append(paymentResponses, paymentResponse)
	}

	return paymentResponses, nil
}
