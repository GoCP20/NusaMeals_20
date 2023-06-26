package controller

import (
	"net/http"
	"reglog/internal/dto/request"
	"reglog/internal/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	CategoryUseCase usecase.CategoryUseCase
}

func NewCategoryController(cu usecase.CategoryUseCase) *CategoryController {
	return &CategoryController{
		CategoryUseCase: cu,
	}
}

func (cc *CategoryController) GetCategoriesController(c echo.Context) error {
	categories, err := cc.CategoryUseCase.GetCategories()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all categories",
		"data":    categories,
	})
}

func (cc *CategoryController) GetMenusByCategoryIDController(c echo.Context) error {
	categoryID, err := strconv.ParseUint(c.QueryParam("categoryID"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	menus, err := cc.CategoryUseCase.GetMenusByCategoryID(uint(categoryID))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get menus by category",
		"data":    menus,
	})
}

func (cc *CategoryController) GetMenusByCategoryNameController(c echo.Context) error {
	categoryName := c.QueryParam("category_name")
	menus, err := cc.CategoryUseCase.GetMenusByCategoryName(categoryName)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get menus by category",
		"data":    menus,
	})
}

func (cc *CategoryController) CreateCategoryController(c echo.Context) error {
	var req request.CreateCategory
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := cc.CategoryUseCase.CreateCategory(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create category",
	})
}

func (cc *CategoryController) UpdateCategoryController(c echo.Context) error {
	categoryID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var req request.UpdateCategory
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = cc.CategoryUseCase.UpdateCategory(uint(categoryID), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update category",
	})
}

func (cc *CategoryController) DeleteCategoryController(c echo.Context) error {
	categoryID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = cc.CategoryUseCase.DeleteCategory(uint(categoryID))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete category",
	})
}
