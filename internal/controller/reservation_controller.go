package controller

import (
	"net/http"
	"reglog/internal/dto/request"
	"reglog/internal/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ReservationController struct {
	reservationUseCase usecase.ReservationUseCase
}

func NewReservationController(reservationUseCase usecase.ReservationUseCase) *ReservationController {
	return &ReservationController{
		reservationUseCase: reservationUseCase,
	}
}

func (c *ReservationController) MakeReservation(ctx echo.Context) error {
	var reservationRequest request.ReservationRequest
	if err := ctx.Bind(&reservationRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	createdReservation, err := c.reservationUseCase.MakeReservation(&reservationRequest)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, createdReservation)
}

func (c *ReservationController) UpdateReservation(ctx echo.Context) error {
	reservationID := ctx.QueryParam("id")
	id, err := strconv.ParseUint(reservationID, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	var updateRequest request.UpdateReservation
	if err := ctx.Bind(&updateRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	updatedReservation, err := c.reservationUseCase.UpdateReservation(uint(id), &updateRequest)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, updatedReservation)
}

func (c *ReservationController) CancelReservation(ctx echo.Context) error {
	reservationID := ctx.QueryParam("id")
	id, err := strconv.ParseUint(reservationID, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid reservation ID"})
	}

	err = c.reservationUseCase.CancelReservation(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"message": "Reservation canceled successfully"})
}

func (c *ReservationController) GetReservationByID(ctx echo.Context) error {
	reservationID := ctx.QueryParam("id")
	id, err := strconv.ParseUint(reservationID, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	reservation, err := c.reservationUseCase.GetReservationByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, reservation)
}

func (c *ReservationController) GetReservationByName(ctx echo.Context) error {
	name := ctx.QueryParam("name")

	reservations, err := c.reservationUseCase.GetReservationByName(name)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, reservations)
}

func (c *ReservationController) GetAllReservations(ctx echo.Context) error {
	reservations, err := c.reservationUseCase.GetAllReservations()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, reservations)
}
