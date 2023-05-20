package controller

import (
	"net/http"

	"loginnregister/models/payload"
	"loginnregister/usecase"

	"github.com/labstack/echo"
)

func GetUserController(c echo.Context) error {
	id := GetUserId(c)
	user, err := usecase.GetUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get the user",
		"data":    user,
	})
}

func UpdateUserController(c echo.Context) error {
	var req payload.UpdateUser
	id := GetUserId(c)

	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return err
	}
	err := usecase.UpdateUser(id, &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
	})
}

func DeleteUserController(c echo.Context) error {
	id := GetUserId(c)

	err := usecase.DeleteUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}
