package usecase

import (
	"errors"
	"reglog/internal/dto/request"
	"reglog/internal/dto/response"
	"reglog/internal/model"
	"reglog/internal/repository"
	"time"
)

type OrderUseCase interface {
	CreateOrder(request request.CreateOrder, timezone string) (response.GetOrderResponse, error)
	GetOrderByID(orderID uint, timezone string) (response.GetOrderDetails, error)
	GetAllOrders(timezone string) ([]response.GetAllOrdersResponse, error)
	GetOrdersByUserID(userID uint, timezone string) ([]response.GetOrderDetails, error)
	UpdateOrderStatus(orderID uint, status string) error
	DeleteOrderByID(orderID uint) error
}

type orderUseCase struct {
	OrderRepo repository.OrderRepository
	UserRepo  repository.UserRepository
	MenuRepo  repository.MenuRepository
}

func NewOrderUseCase(orderRepo repository.OrderRepository, userRepo repository.UserRepository, menuRepo repository.MenuRepository) OrderUseCase {
	return &orderUseCase{
		OrderRepo: orderRepo,
		UserRepo:  userRepo,
		MenuRepo:  menuRepo,
	}
}

func (uc *orderUseCase) GetAllOrders(timezone string) ([]response.GetAllOrdersResponse, error) {
	orders, err := uc.OrderRepo.GetAllOrders()
	if err != nil {
		return nil, err
	}

	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return nil, err
	}

	var orderResponses []response.GetAllOrdersResponse
	for _, order := range orders {
		createdAtFormatted := order.CreatedAt.In(loc).Format("2006-01-02 15:04")
		orderResponse := response.GetAllOrdersResponse{
			ID:          order.ID,
			UserID:      order.UserID,
			MenuID:      order.MenuID,
			Quantity:    order.Quantity,
			TypeOrder:   order.TypeOrder,
			TotalPrice:  order.TotalPrice,
			OrderStatus: order.OrderStatus,
			CreatedAt:   createdAtFormatted,
		}
		orderResponses = append(orderResponses, orderResponse)
	}

	return orderResponses, nil
}

func (uc *orderUseCase) CreateOrder(request request.CreateOrder, timezone string) (response.GetOrderResponse, error) {
	var user model.User
	var menu model.Menu
	var err error

	if request.UserID != 0 {
		user, err = uc.UserRepo.GetUserByID(request.UserID)
	} else if request.Username != "" {
		user, err = uc.UserRepo.GetUserByUsername(request.Username)
	} else {
		return response.GetOrderResponse{}, errors.New("user ID or username is required")
	}
	if err != nil {
		return response.GetOrderResponse{}, err
	}

	if request.MenuID != 0 {
		menu, err = uc.MenuRepo.GetMenuByID(request.MenuID)
	} else if request.MenuName != "" {
		menus, err := uc.MenuRepo.GetMenusByName(request.MenuName)
		if err != nil {
			return response.GetOrderResponse{}, err
		}
		if len(menus) == 0 {
			return response.GetOrderResponse{}, errors.New("menu not found")
		}
		menu = menus[0]
	} else {
		return response.GetOrderResponse{}, errors.New("menu ID or menu name is required")
	}
	if err != nil {
		return response.GetOrderResponse{}, err
	}

	totalPrice := menu.Price * request.Quantity

	order := &model.Order{
		UserID:         user.ID,
		MenuID:         menu.ID,
		Quantity:       request.Quantity,
		TypeOrder:      request.TypeOrder,
		TableNumber:    request.TableNumber,
		TotalPrice:     totalPrice,
		PaymentMethods: request.PaymentMethods,
		OrderStatus:    "New Order",
	}

	err = uc.OrderRepo.CreateOrder(order)
	if err != nil {
		return response.GetOrderResponse{}, err
	}
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return response.GetOrderResponse{}, err
	}
	currentTime := time.Now().In(location)
	createdAtFormatted := currentTime.Format("2006-01-02 15:04")

	orderResponse := response.GetOrderResponse{
		ID:             order.ID,
		UserID:         user.ID,
		MenuID:         menu.ID,
		MenuName:       menu.Name,
		MenuImages:     menu.Images,
		MenuCity:       menu.City,
		MenuCalories:   menu.Calories,
		Quantity:       order.Quantity,
		PriceMenu:      menu.Price,
		TypeOrder:      order.TypeOrder,
		TableNumber:    order.TableNumber,
		TotalPrice:     order.TotalPrice,
		PaymentMethods: order.PaymentMethods,
		OrderStatus:    order.OrderStatus,
		CreatedAt:      createdAtFormatted,
	}

	return orderResponse, nil
}

func (uc *orderUseCase) GetOrderByID(orderID uint, timezone string) (response.GetOrderDetails, error) {
	order, err := uc.OrderRepo.GetOrderByID(orderID)
	if err != nil {
		return response.GetOrderDetails{}, err
	}

	user, err := uc.UserRepo.GetUserByID(order.UserID)
	if err != nil {
		return response.GetOrderDetails{}, err
	}

	menu, err := uc.MenuRepo.GetMenuByID(order.MenuID)
	if err != nil {
		return response.GetOrderDetails{}, err
	}

	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return response.GetOrderDetails{}, err
	}

	createdAtFormatted := order.CreatedAt.In(loc).Format("2006-01-02 15:04")
	orderResponse := response.GetOrderDetails{
		ID:             order.ID,
		UserID:         user.ID,
		Username:       user.Username,
		MenuID:         menu.ID,
		MenuName:       menu.Name,
		MenuImages:     menu.Images,
		MenuCity:       menu.City,
		MenuCalories:   menu.Calories,
		Quantity:       order.Quantity,
		PriceMenu:      menu.Price,
		TypeOrder:      order.TypeOrder,
		TableNumber:    order.TableNumber,
		TotalPrice:     order.TotalPrice,
		PaymentMethods: order.PaymentMethods,
		OrderStatus:    order.OrderStatus,
		CreatedAt:      createdAtFormatted,
	}

	return orderResponse, nil
}

func (uc *orderUseCase) GetOrdersByUserID(userID uint, timezone string) ([]response.GetOrderDetails, error) {
	var responseOrders []response.GetOrderDetails

	orders, err := uc.OrderRepo.GetOrdersByUserID(userID)
	if err != nil {
		return nil, err
	}

	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return nil, err
	}

	for _, order := range orders {
		user, err := uc.UserRepo.GetUserByID(order.UserID)
		if err != nil {
			return nil, err
		}

		menu, err := uc.MenuRepo.GetMenuByID(order.MenuID)
		if err != nil {
			return nil, err
		}

		createdAtFormatted := order.CreatedAt.In(loc).Format("2006-01-02 15:04")

		orderResponse := response.GetOrderDetails{
			ID:             order.ID,
			UserID:         user.ID,
			Username:       user.Username,
			MenuID:         menu.ID,
			MenuName:       menu.Name,
			MenuImages:     menu.Images,
			MenuCity:       menu.City,
			MenuCalories:   menu.Calories,
			Quantity:       order.Quantity,
			PriceMenu:      menu.Price,
			TypeOrder:      order.TypeOrder,
			TableNumber:    order.TableNumber,
			TotalPrice:     order.TotalPrice,
			PaymentMethods: order.PaymentMethods,
			OrderStatus:    order.OrderStatus,
			CreatedAt:      createdAtFormatted,
		}

		responseOrders = append(responseOrders, orderResponse)
	}

	return responseOrders, nil
}

func (uc *orderUseCase) UpdateOrderStatus(orderID uint, status string) error {
	err := uc.OrderRepo.UpdateOrderStatus(orderID, status)
	if err != nil {
		return err
	}

	return nil
}

func (uc *orderUseCase) DeleteOrderByID(orderID uint) error {
	err := uc.OrderRepo.DeleteOrder(orderID)
	if err != nil {
		return err
	}

	return nil
}
