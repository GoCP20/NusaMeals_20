package repository

import (
	"reglog/internal/model"

	"github.com/jinzhu/gorm"
)

type CategoryRepository interface {
	CreateCategory(data model.Category) error
	GetAllCategories() ([]model.Category, error)
	GetCategoryByID(ID uint) (model.Category, error)
	GetMenusByCategory(categoryID uint) ([]model.Menu, error)
	GetCategoryByName(name string) (model.Category, error)
	UpdateCategory(data model.Category) error
	DeleteCategory(ID uint) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) CreateCategory(data model.Category) error {
	err := r.db.Create(&data).Error
	return err
}

func (r *categoryRepository) GetAllCategories() ([]model.Category, error) {
	var categories []model.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) GetCategoryByID(ID uint) (model.Category, error) {
	var category model.Category
	err := r.db.First(&category, ID).Error
	return category, err
}

func (r *categoryRepository) GetMenusByCategory(categoryID uint) ([]model.Menu, error) {
	var menus []model.Menu
	err := r.db.Where("category_id = ?", categoryID).Find(&menus).Error
	return menus, err
}

func (r *categoryRepository) GetCategoryByName(name string) (model.Category, error) {
	var category model.Category
	err := r.db.Where("name = ?", name).Preload("Menus").First(&category).Error
	return category, err
}

func (r *categoryRepository) UpdateCategory(data model.Category) error {
	err := r.db.Save(&data).Error
	return err
}

func (r *categoryRepository) DeleteCategory(ID uint) error {
	err := r.db.Delete(&model.Category{}, ID).Error
	return err
}
