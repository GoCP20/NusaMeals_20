package controller

import (
	"net/http"
	"reglog/internal/dto/request"
	"reglog/internal/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PaymentController struct {
	paymentUseCase usecase.PaymentUseCase
}

func NewPaymentController(paymentUseCase usecase.PaymentUseCase) *PaymentController {
	return &PaymentController{
		paymentUseCase: paymentUseCase,
	}
}

func (c *PaymentController) CreatePayment(ctx echo.Context) error {
	var paymentDTO request.Payment
	if err := ctx.Bind(&paymentDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	createdPayment, err := c.paymentUseCase.CreatePayment(&paymentDTO)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, createdPayment)
}

func (c *PaymentController) UpdatePayment(ctx echo.Context) error {
	paymentID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	var paymentDTO request.Payment
	if err := ctx.Bind(&paymentDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	updatedPayment, err := c.paymentUseCase.UpdatePayment(uint(paymentID), &paymentDTO)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, updatedPayment)
}

func (c *PaymentController) UpdatePaymentByAdmin(ctx echo.Context) error {
	paymentIDStr := ctx.QueryParam("id")
	paymentID, err := strconv.ParseUint(paymentIDStr, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	var paymentUpdateDTO request.PaymentUpdate
	if err := ctx.Bind(&paymentUpdateDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	updatedPayment, err := c.paymentUseCase.UpdatePaymentByAdmin(uint(paymentID), &paymentUpdateDTO)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"message": "Updated payment status successfully!", "data": updatedPayment})
}

func (c *PaymentController) DeletePayment(ctx echo.Context) error {
	paymentID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid payment ID"})
	}

	err = c.paymentUseCase.DeletePayment(uint(paymentID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"message": "Payment deleted successfully"})
}

func (c *PaymentController) GetPaymentByID(ctx echo.Context) error {
	paymentID, err := strconv.ParseUint(ctx.QueryParam("ID"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	payment, err := c.paymentUseCase.GetPaymentByID(uint(paymentID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, payment)
}

func (c *PaymentController) GetPaymentByOrderID(ctx echo.Context) error {
	orderID, err := strconv.ParseUint(ctx.QueryParam("ID"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	payments, err := c.paymentUseCase.GetPaymentByOrderID(uint(orderID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, payments)
}

func (c *PaymentController) GetPaymentByUsername(ctx echo.Context) error {
	username := ctx.QueryParam("username")

	payments, err := c.paymentUseCase.GetPaymentByUsername(username)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, payments)
}

func (c *PaymentController) GetAllPayments(ctx echo.Context) error {
	payments, err := c.paymentUseCase.GetAllPayments()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to fetch payments",
		})
	}

	return ctx.JSON(http.StatusOK, payments)
}
