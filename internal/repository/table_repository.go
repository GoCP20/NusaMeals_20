package repository

import (
	"reglog/internal/model"

	"github.com/jinzhu/gorm"
)

type TableRepository interface {
	AddTable(data *model.Table) error
	GetAllTables() ([]model.Table, error)
	GetTableByID(ID uint) (*model.Table, error)
	UpdateTable(data *model.Table) error
	DeleteTable(ID uint) error
}

type tableRepository struct {
	db *gorm.DB
}

func NewTableRepository(db *gorm.DB) TableRepository {
	return &tableRepository{db: db}
}

func (r *tableRepository) AddTable(data *model.Table) error {
	err := r.db.Create(data).Error
	return err
}

func (r *tableRepository) GetAllTables() ([]model.Table, error) {
	var tables []model.Table
	err := r.db.Find(&tables).Error
	return tables, err
}

func (r *tableRepository) GetTableByID(ID uint) (*model.Table, error) {
	table := &model.Table{}
	err := r.db.First(table, ID).Error
	return table, err
}

func (r *tableRepository) UpdateTable(data *model.Table) error {
	err := r.db.Save(data).Error
	return err
}

func (r *tableRepository) DeleteTable(ID uint) error {
	err := r.db.Delete(&model.Table{}, ID).Error
	return err
}
