package usecase

import (
	"reglog/internal/dto/request"
	"reglog/internal/dto/response"
	"reglog/internal/model"
	"reglog/internal/repository"
)

type TableUseCase interface {
	AddTable(tableRequest *request.CreateTable) (*response.CreateTableResponse, error)
	GetAllTables() ([]response.CreateTableResponse, error)
	GetTableByID(tableID uint) (*response.CreateTableResponse, error)
	UpdateTable(tableID uint, tableRequest *request.UpdateTable) (*response.UpdateTableResponse, error)
	DeleteTable(tableID uint) error
}

type tableUseCase struct {
	tableRepo repository.TableRepository
}

func NewTableUseCase(tableRepo repository.TableRepository) TableUseCase {
	return &tableUseCase{
		tableRepo: tableRepo,
	}
}

func (u *tableUseCase) AddTable(tableRequest *request.CreateTable) (*response.CreateTableResponse, error) {
	table := model.Table{
		NumberTable:   tableRequest.NumberTable,
		Seats:         tableRequest.Seats,
		Type:          tableRequest.Type,
		TableLocation: tableRequest.TableLocation,
		Picture:       tableRequest.Picture,
	}

	err := u.tableRepo.AddTable(&table)
	if err != nil {
		return nil, err
	}

	tableResponse := &response.CreateTableResponse{
		ID:            table.ID,
		NumberTable:   table.NumberTable,
		Seats:         table.Seats,
		Type:          table.Type,
		TableLocation: table.TableLocation,
		Picture:       table.Picture,
	}

	return tableResponse, nil
}

func (u *tableUseCase) GetAllTables() ([]response.CreateTableResponse, error) {
	tables, err := u.tableRepo.GetAllTables()
	if err != nil {
		return nil, err
	}

	var tableResponses []response.CreateTableResponse
	for _, table := range tables {
		tableResponse := response.CreateTableResponse{
			ID:            table.ID,
			NumberTable:   table.NumberTable,
			Seats:         table.Seats,
			Type:          table.Type,
			TableLocation: table.TableLocation,
			Picture:       table.Picture,
		}

		tableResponses = append(tableResponses, tableResponse)
	}

	return tableResponses, nil
}

func (u *tableUseCase) GetTableByID(tableID uint) (*response.CreateTableResponse, error) {
	table, err := u.tableRepo.GetTableByID(tableID)
	if err != nil {
		return nil, err
	}

	tableResponse := &response.CreateTableResponse{
		ID:            table.ID,
		NumberTable:   table.NumberTable,
		Seats:         table.Seats,
		Type:          table.Type,
		TableLocation: table.TableLocation,
		Picture:       table.Picture,
	}

	return tableResponse, nil
}

func (u *tableUseCase) UpdateTable(tableID uint, tableRequest *request.UpdateTable) (*response.UpdateTableResponse, error) {
	table, err := u.tableRepo.GetTableByID(tableID)
	if err != nil {
		return nil, err
	}

	table.NumberTable = tableRequest.NumberTable
	table.Seats = tableRequest.Seats
	table.Type = tableRequest.Type
	table.TableLocation = tableRequest.TableLocation
	table.Picture = tableRequest.Picture

	err = u.tableRepo.UpdateTable(table)
	if err != nil {
		return nil, err
	}

	tableResponse := &response.UpdateTableResponse{
		ID:            table.ID,
		NumberTable:   table.NumberTable,
		Seats:         table.Seats,
		Type:          table.Type,
		TableLocation: table.TableLocation,
		Picture:       table.Picture,
	}

	return tableResponse, nil
}

func (u *tableUseCase) DeleteTable(tableID uint) error {
	err := u.tableRepo.DeleteTable(tableID)
	return err
}
