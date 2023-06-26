package usecase

import (
	"errors"
	"reglog/internal/dto/request"
	"reglog/internal/dto/response"
	"reglog/internal/model"
	"reglog/internal/repository"
)

type CategoryUseCase interface {
	CreateCategory(req request.CreateCategory) error
	GetCategories() ([]string, error)
	GetCategoryByID(ID uint) (model.Category, error)
	GetMenusByCategoryID(categoryID uint) ([]response.GetMenuResponse, error)
	GetMenusByCategoryName(name string) ([]response.GetMenuResponse, error)
	UpdateCategory(ID uint, req request.UpdateCategory) error
	DeleteCategory(ID uint) error
}

type categoryUseCase struct {
	CategoryRepo repository.CategoryRepository
	MenuRepo     repository.MenuRepository
}

func NewCategoryUseCase(categoryRepo repository.CategoryRepository, menuRepo repository.MenuRepository) CategoryUseCase {
	return &categoryUseCase{
		CategoryRepo: categoryRepo,
		MenuRepo:     menuRepo,
	}
}

func (uc *categoryUseCase) CreateCategory(req request.CreateCategory) error {
	category := model.Category{
		Name: req.Name,
	}
	return uc.CategoryRepo.CreateCategory(category)
}

func (uc *categoryUseCase) GetCategories() ([]string, error) {
	categories, err := uc.CategoryRepo.GetAllCategories()
	if err != nil {
		return nil, err
	}

	var response []string
	for _, category := range categories {
		response = append(response, category.Name)
	}

	return response, nil
}

func (uc *categoryUseCase) GetCategoryByID(ID uint) (model.Category, error) {
	return uc.CategoryRepo.GetCategoryByID(ID)
}

func (uc *categoryUseCase) GetMenusByCategoryID(categoryID uint) ([]response.GetMenuResponse, error) {
	menus, err := uc.CategoryRepo.GetMenusByCategory(categoryID)
	if err != nil {
		return nil, err
	}

	category, err := uc.CategoryRepo.GetCategoryByID(categoryID)
	if err != nil {
		return nil, err
	}

	if category.ID == 0 {
		return nil, errors.New("category not found")
	}

	var responseMenus []response.GetMenuResponse
	for _, menu := range menus {
		responseMenu := response.GetMenuResponse{
			ID:          menu.ID,
			Name:        menu.Name,
			Price:       menu.Price,
			CategoryID:  menu.CategoryID,
			Category:    category.Name,
			Calories:    menu.Calories,
			City:        menu.City,
			Description: menu.Description,
			Ingredients: menu.Ingredient,
			Images:      menu.Images,
		}
		responseMenus = append(responseMenus, responseMenu)
	}

	return responseMenus, nil
}

func (uc *categoryUseCase) GetMenusByCategoryName(name string) ([]response.GetMenuResponse, error) {
	category, err := uc.CategoryRepo.GetCategoryByName(name)
	if err != nil {
		return nil, err
	}

	if category.ID == 0 {
		return nil, errors.New("category not found")
	}

	var responseMenus []response.GetMenuResponse
	for _, menu := range category.Menus {
		responseMenu := response.GetMenuResponse{
			ID:          menu.ID,
			Name:        menu.Name,
			Price:       menu.Price,
			CategoryID:  menu.CategoryID,
			Category:    category.Name,
			Calories:    menu.Calories,
			City:        menu.City,
			Description: menu.Description,
			Ingredients: menu.Ingredient,
			Images:      menu.Images,
		}
		responseMenus = append(responseMenus, responseMenu)
	}

	return responseMenus, nil
}

func (uc *categoryUseCase) UpdateCategory(ID uint, req request.UpdateCategory) error {
	category, err := uc.CategoryRepo.GetCategoryByID(ID)
	if err != nil {
		return err
	}

	// Update the category fields based on the request
	category.Name = req.Name

	return uc.CategoryRepo.UpdateCategory(category)
}

func (uc *categoryUseCase) DeleteCategory(ID uint) error {
	return uc.CategoryRepo.DeleteCategory(ID)
}
