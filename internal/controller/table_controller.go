package controller

import (
	"net/http"
	"reglog/internal/dto/request"
	"reglog/internal/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TableController struct {
	tableUseCase usecase.TableUseCase
}

func NewTableController(tableUseCase usecase.TableUseCase) *TableController {
	return &TableController{
		tableUseCase: tableUseCase,
	}
}

func (c *TableController) AddTable(ctx echo.Context) error {
	var tableRequest request.CreateTable
	if err := ctx.Bind(&tableRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	createdTable, err := c.tableUseCase.AddTable(&tableRequest)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, createdTable)
}

func (c *TableController) UpdateTable(ctx echo.Context) error {
	tableID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	var tableRequest request.UpdateTable
	if err := ctx.Bind(&tableRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	updatedTable, err := c.tableUseCase.UpdateTable(uint(tableID), &tableRequest)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, updatedTable)
}

func (c *TableController) DeleteTable(ctx echo.Context) error {
	tableID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid table ID"})
	}

	err = c.tableUseCase.DeleteTable(uint(tableID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"message": "Table deleted successfully"})
}

func (c *TableController) GetTableByID(ctx echo.Context) error {
	tableID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	table, err := c.tableUseCase.GetTableByID(uint(tableID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, table)
}

func (c *TableController) GetAllTables(ctx echo.Context) error {
	tables, err := c.tableUseCase.GetAllTables()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, tables)
}
