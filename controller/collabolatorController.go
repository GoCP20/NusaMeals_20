package controller

import (
	"loginnregister/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func CreateCollabolatorController(c echo.Context) error {
	id := GetUserId(c)
	taskId, _ := strconv.Atoi(c.Param("id"))
	err := usecase.CreateCollabolator(uint(taskId), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create collabolator",
	})
}
