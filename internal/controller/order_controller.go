package controller

import (
	"net/http"
	"reglog/internal/dto/request"
	"reglog/internal/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderController struct {
	orderUseCase usecase.OrderUseCase
}

func NewOrderController(orderUseCase usecase.OrderUseCase) *OrderController {
	return &OrderController{
		orderUseCase: orderUseCase,
	}
}

func (oc *OrderController) GetAllOrders(c echo.Context) error {
	timezone := c.Request().Header.Get("Timezone")
	orders, err := oc.orderUseCase.GetAllOrders(timezone)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to get orders",
		})
	}

	return c.JSON(http.StatusOK, orders)
}

func (oc *OrderController) CreateOrder(c echo.Context) error {
	orderDTO := new(request.CreateOrder)
	if err := c.Bind(orderDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request payload",
		})
	}

	timezone := c.Request().Header.Get("Timezone")
	orderResponse, err := oc.orderUseCase.CreateOrder(*orderDTO, timezone)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to create order",
		})
	}

	return c.JSON(http.StatusOK, orderResponse)
}

func (oc *OrderController) GetOrderByID(c echo.Context) error {
	orderID := c.Param("orderID")
	// Convert orderID string to uint
	orderIDUint, err := strconv.ParseUint(orderID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid order ID",
		})
	}

	timezone := c.Request().Header.Get("Timezone")
	order, err := oc.orderUseCase.GetOrderByID(uint(orderIDUint), timezone)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": "Order not found",
		})
	}

	return c.JSON(http.StatusOK, order)
}

func (oc *OrderController) UpdateOrderStatus(c echo.Context) error {
	orderID := c.Param("orderID")
	// Convert orderID string to uint
	orderIDUint, err := strconv.ParseUint(orderID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid order ID",
		})
	}

	statusDTO := new(request.UpdateOrder)
	if err := c.Bind(statusDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request payload",
		})
	}

	err = oc.orderUseCase.UpdateOrderStatus(uint(orderIDUint), statusDTO.OrderStatus)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to update order status",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Order status updated successfully",
	})
}

func (oc *OrderController) DeleteOrder(c echo.Context) error {
	orderID := c.Param("orderID")
	// Convert orderID string to uint
	orderIDUint, err := strconv.ParseUint(orderID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid order ID",
		})
	}

	err = oc.orderUseCase.DeleteOrderByID(uint(orderIDUint))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to delete order",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "The order deleted successfully",
	})
}

func (oc *OrderController) GetOrdersByUserID(c echo.Context) error {
	userID := c.Param("userID")
	// Convert userID string to uint
	userIDUint, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid user ID",
		})
	}

	timezone := c.Request().Header.Get("Timezone")
	orders, err := oc.orderUseCase.GetOrdersByUserID(uint(userIDUint), timezone)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to get orders",
		})
	}

	return c.JSON(http.StatusOK, orders)
}
