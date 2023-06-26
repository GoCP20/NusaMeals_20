package controller

import (
	"net/http"
	"reglog/internal/dto/request"
	"reglog/internal/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserUseCase usecase.UserUseCase
}

func NewUserController(uc usecase.UserUseCase) *UserController {
	return &UserController{
		UserUseCase: uc,
	}
}

func (h *UserController) GetAllUser(c echo.Context) error {
	users, err := h.UserUseCase.GetAllUser()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func (h *UserController) GetUserByID(c echo.Context) error {
	IDStr := c.QueryParam("ID")
	ID, err := strconv.ParseUint(IDStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	user, err := h.UserUseCase.GetUserByID(uint(ID))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserController) GetUserByUsername(c echo.Context) error {
	username := c.QueryParam("username")

	user, err := h.UserUseCase.GetUserByUsername(username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserController) GetUserByEmail(c echo.Context) error {
	email := c.QueryParam("email")

	user, err := h.UserUseCase.GetUserByEmail(email)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserController) UpdateUser(c echo.Context) error {
	ID := c.Param("id")

	// Parse request body to get update data
	var updateUser request.UpdateUser
	if err := c.Bind(&updateUser); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Convert ID to uint
	idUint, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}
	userID := uint(idUint)

	// Update user
	updatedUser, err := h.UserUseCase.UpdateUser(userID, updateUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, updatedUser)
}
