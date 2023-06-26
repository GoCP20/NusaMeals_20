package usecase

import (
	"reglog/internal/dto/request"
	"reglog/internal/dto/response"
	"reglog/internal/model"
	"reglog/internal/repository"
)

type MenuUseCase interface {
	CreateMenu(request request.Menu) error
	GetMenuByID(ID uint) (response.GetMenuResponse, error)
	GetAllMenus() ([]response.GetMenuResponse, error)
	GetMenusByName(name string) ([]response.GetMenuResponse, error)
	GetMenusByCategory(categoryTypeID uint) ([]response.GetMenuResponse, error)
	GetMenusByCategoryName(categoryName string) ([]response.GetMenuResponse, error)
	UpdateMenu(ID uint, request request.UpdateMenu) error
	DeleteMenuByID(ID uint) error
}

type menuUseCase struct {
	MenuRepo repository.MenuRepository
}

func NewMenuUseCase(menuRepo repository.MenuRepository) MenuUseCase {
	return &menuUseCase{
		MenuRepo: menuRepo,
	}
}

func (uc *menuUseCase) CreateMenu(request request.Menu) error {
	menu := &model.Menu{
		Name:        request.Name,
		Price:       request.Price,
		CategoryID:  request.CategoryID,
		Calories:    request.Calories,
		City:        request.City,
		Description: request.Description,
		Ingredient:  request.Ingredient,
		Images:      request.Images,
	}

	err := uc.MenuRepo.CreateMenu(menu)
	if err != nil {
		return err
	}
	return nil
}

func (uc *menuUseCase) GetMenuByID(ID uint) (response.GetMenuResponse, error) {
	menu, err := uc.MenuRepo.GetMenuByID(ID)
	if err != nil {
		return response.GetMenuResponse{}, err
	}

	menuResponse := response.GetMenuResponse{
		ID:          menu.ID,
		CategoryID:  menu.CategoryID,
		Name:        menu.Name,
		Price:       menu.Price,
		Category:    menu.Category.Name,
		Calories:    menu.Calories,
		City:        menu.City,
		Description: menu.Description,
		Ingredients: menu.Ingredient,
		Images:      menu.Images,
	}

	return menuResponse, nil
}

func (uc *menuUseCase) GetAllMenus() ([]response.GetMenuResponse, error) {
	var responseMenu []response.GetMenuResponse

	menus, err := uc.MenuRepo.GetMenus()
	if err != nil {
		return nil, err
	}

	for _, item := range menus {
		response := response.GetMenuResponse{
			ID:          item.ID,
			CategoryID:  item.CategoryID,
			Name:        item.Name,
			Price:       item.Price,
			Category:    item.Category.Name,
			Calories:    item.Calories,
			City:        item.City,
			Description: item.Description,
			Ingredients: item.Ingredient,
			Images:      item.Images,
		}

		responseMenu = append(responseMenu, response)
	}

	return responseMenu, nil
}

func (uc *menuUseCase) GetMenusByName(name string) ([]response.GetMenuResponse, error) {
	menus, err := uc.MenuRepo.GetMenusByName(name)
	if err != nil {
		return nil, err
	}

	var responseMenus []response.GetMenuResponse
	for _, menu := range menus {
		response := response.GetMenuResponse{
			ID:          menu.ID,
			CategoryID:  menu.CategoryID,
			Name:        menu.Name,
			Price:       menu.Price,
			Category:    menu.Category.Name,
			Calories:    menu.Calories,
			City:        menu.City,
			Description: menu.Description,
			Ingredients: menu.Ingredient,
			Images:      menu.Images,
		}
		responseMenus = append(responseMenus, response)
	}

	return responseMenus, nil
}

func (uc *menuUseCase) GetMenusByCategory(categoryID uint) ([]response.GetMenuResponse, error) {
	menus, err := uc.MenuRepo.GetMenusByCategory(categoryID)
	if err != nil {
		return nil, err
	}

	var responseMenus []response.GetMenuResponse
	for _, menu := range menus {
		response := response.GetMenuResponse{
			ID:          menu.ID,
			CategoryID:  menu.CategoryID,
			Name:        menu.Name,
			Price:       menu.Price,
			Category:    menu.Category.Name,
			Calories:    menu.Calories,
			City:        menu.City,
			Description: menu.Description,
			Ingredients: menu.Ingredient,
			Images:      menu.Images,
		}
		responseMenus = append(responseMenus, response)
	}

	return responseMenus, nil
}

func (uc *menuUseCase) GetMenusByCategoryName(categoryName string) ([]response.GetMenuResponse, error) {
	menus, err := uc.MenuRepo.GetMenusByCategoryName(categoryName)
	if err != nil {
		return nil, err
	}

	var responseMenus []response.GetMenuResponse
	for _, menu := range menus {
		response := response.GetMenuResponse{
			ID:          menu.ID,
			CategoryID:  menu.CategoryID,
			Name:        menu.Name,
			Price:       menu.Price,
			Category:    menu.Category.Name,
			Calories:    menu.Calories,
			City:        menu.City,
			Description: menu.Description,
			Ingredients: menu.Ingredient,
			Images:      menu.Images,
		}
		responseMenus = append(responseMenus, response)
	}

	return responseMenus, nil
}

func (uc *menuUseCase) UpdateMenu(ID uint, request request.UpdateMenu) error {
	menu, err := uc.MenuRepo.GetMenuByID(ID)
	if err != nil {
		return err
	}

	menu.Name = request.Name
	menu.Price = request.Price
	menu.Calories = request.Calories
	menu.City = request.City
	menu.Description = request.Description
	menu.Ingredient = request.Ingredient
	menu.Images = request.Images
	menu.CategoryID = request.CategoryID

	err = uc.MenuRepo.UpdateMenu(&menu)
	if err != nil {
		return err
	}
	return nil
}

func (uc *menuUseCase) DeleteMenuByID(ID uint) error {
	err := uc.MenuRepo.DeleteMenuByID(ID)
	if err != nil {
		return err
	}
	return nil
}
