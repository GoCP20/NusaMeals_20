package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"reglog/internal/dto/request"
	"reglog/internal/usecase"
)

type MenuController struct {
	MenuUseCase usecase.MenuUseCase
}

func NewMenuController(cu usecase.MenuUseCase) *MenuController {
	return &MenuController{
		MenuUseCase: cu,
	}
}

func (mc *MenuController) GetAllMenusController(c echo.Context) error {
	menus, err := mc.MenuUseCase.GetAllMenus()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all menus",
		"data":    menus,
	})
}

func (mc *MenuController) GetMenuController(c echo.Context) error {
	id := c.Param("id")
	menuID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	menu, err := mc.MenuUseCase.GetMenuByID(uint(menuID))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get menu",
		"data":    menu,
	})
}

func (mc *MenuController) GetMenusByNameController(c echo.Context) error {
	name := c.QueryParam("name")

	menus, err := mc.MenuUseCase.GetMenusByName(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get menus by name",
		"data":    menus,
	})
}

func (mc *MenuController) GetMenusByCategoryController(c echo.Context) error {
	categoryID, err := strconv.ParseUint(c.QueryParam("category"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	menus, err := mc.MenuUseCase.GetMenusByCategory(uint(categoryID))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get menus by category",
		"data":    menus,
	})
}

func (mc *MenuController) GetMenusByCategoryNameController(c echo.Context) error {
	categoryName := c.QueryParam("categoryName")

	menus, err := mc.MenuUseCase.GetMenusByCategoryName(categoryName)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get menus by category name",
		"data":    menus,
	})
}

func (mc *MenuController) CreateMenuController(c echo.Context) error {
	var req request.Menu
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = mc.MenuUseCase.CreateMenu(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create menu",
	})
}

func (mc *MenuController) UpdateMenuController(c echo.Context) error {
	id := c.Param("id")
	menuID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var req request.UpdateMenu
	err = c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = mc.MenuUseCase.UpdateMenu(uint(menuID), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update menu",
	})
}

func (mc *MenuController) DeleteMenuController(c echo.Context) error {
	id := c.Param("id")
	menuID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = mc.MenuUseCase.DeleteMenuByID(uint(menuID))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete menu",
	})
}
