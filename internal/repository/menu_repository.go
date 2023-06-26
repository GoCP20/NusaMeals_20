package repository

import (
	"reglog/internal/model"

	"github.com/jinzhu/gorm"
)

type MenuRepository interface {
	GetMenus() ([]model.Menu, error)
	GetAllMenus() ([]model.Menu, error)
	GetMenuByID(id uint) (model.Menu, error)
	GetMenusByName(name string) ([]model.Menu, error)
	GetMenusByCategory(categoryID uint) ([]model.Menu, error)
	GetMenusByCategoryName(categoryName string) ([]model.Menu, error)
	CreateMenu(req *model.Menu) error
	UpdateMenu(menu *model.Menu) error
	DeleteMenuByID(id uint) error
}

type menuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) *menuRepository {
	return &menuRepository{db: db}
}

func (r *menuRepository) GetMenus() ([]model.Menu, error) {
	var menus []model.Menu
	err := r.db.Preload("Category").Find(&menus).Error
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (r *menuRepository) GetAllMenus() ([]model.Menu, error) {
	var menus []model.Menu
	err := r.db.Find(&menus).Error
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (r *menuRepository) GetMenuByID(id uint) (model.Menu, error) {
	var menu model.Menu
	err := r.db.Preload("Category").Where("id = ?", id).First(&menu).Error
	if err != nil {
		return model.Menu{}, err
	}
	return menu, nil
}

func (r *menuRepository) GetMenusByName(name string) ([]model.Menu, error) {
	var menus []model.Menu
	err := r.db.Preload("Category").Where("name = ?", name).Find(&menus).Error
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (r *menuRepository) GetMenusByCategory(categoryID uint) ([]model.Menu, error) {
	var menus []model.Menu
	err := r.db.Preload("Category").Where("category_id = ?", categoryID).Find(&menus).Error
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (r *menuRepository) GetMenusByCategoryName(name string) ([]model.Menu, error) {
	var menus []model.Menu
	err := r.db.Preload("Category").Where("name = ?", name).Find(&menus).Error
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (r *menuRepository) CreateMenu(menu *model.Menu) error {
	err := r.db.Create(menu).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *menuRepository) UpdateMenu(menu *model.Menu) error {
	err := r.db.Save(&menu).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *menuRepository) DeleteMenuByID(id uint) error {
	err := r.db.Where("id = ?", id).Delete(&model.Menu{}).Error
	if err != nil {
		return err
	}
	return nil
}
