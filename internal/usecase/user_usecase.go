package usecase

import (
	"errors"
	"fmt"
	"reglog/internal/common/middleware"
	"reglog/internal/dto/request"
	"reglog/internal/dto/response"
	"reglog/internal/model"
	"reglog/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	RegisterUser(request request.RegisterUser) error
	LoginUser(request request.LoginUser) (*response.LoginUser, error)
	GetAllUser() ([]response.User, error)
	GetUserByID(ID uint) (response.User, error)
	GetUserByUsername(username string) (response.User, error)
	GetUserByEmail(email string) (response.User, error)
	UpdateUser(ID uint, data request.UpdateUser) (response.User, error)
	DeleteUser(ID uint) error
}

type userUseCase struct {
	UserRepo    repository.UserRepository
	JwtProvider *middleware.JWTProvider
}

func NewUserUseCase(ur repository.UserRepository, jp *middleware.JWTProvider) UserUseCase {
	return &userUseCase{
		UserRepo:    ur,
		JwtProvider: jp,
	}
}

func (u *userUseCase) RegisterUser(request request.RegisterUser) error {
	// Check Available Username
	_, err := u.UserRepo.GetUserByUsername(request.Username)
	if err == nil {
		return errors.New("username is already Exists")
	}

	// Check Available Email
	_, err = u.UserRepo.GetUserByEmail(request.Email)
	if err == nil {
		return errors.New("email is already Exists")
	}

	// Encryption Password from string to Bcrypt
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to encrypt password")
	}

	// Transfer Object Request to Model User
	newUser := model.User{
		Username: request.Username,
		Email:    request.Email,
		Password: string(hashPassword),
		Role:     request.Role,
	}

	// Save User via repository
	err = u.UserRepo.CreateUser(newUser)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUseCase) LoginUser(request request.LoginUser) (*response.LoginUser, error) {
	var loginResponse response.LoginUser

	// Check if the provided value is an email or username
	user, err := u.UserRepo.GetUserByEmailOrUsername(request.EmailOrUsername)
	if err != nil {
		return nil, errors.New("user or email not found")
	}
	// validation password should match with bcrypt method
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	// Setup Token
	token, errToken := u.JwtProvider.CreateToken(user.Username, user.Role, user.ID)
	if errToken != nil {
		return nil, errors.New("failed to create token")
	}

	// Transfer model User to object Login Response
	loginResponse = response.LoginUser{
		ID:       user.ID,
		Username: user.Username,
		Token:    token,
	}

	return &loginResponse, nil
}

func (u *userUseCase) GetAllUser() ([]response.User, error) {
	var usersResponse []response.User

	users, err := u.UserRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		userResponse := response.User{
			ID:          user.ID,
			Name:        user.Name,
			Username:    user.Username,
			Email:       user.Email,
			Gender:      user.Gender,
			PhoneNumber: user.PhoneNumber,
			Picture:     user.Picture,
		}
		usersResponse = append(usersResponse, userResponse)
	}

	return usersResponse, nil
}

func (u *userUseCase) GetUserByID(ID uint) (response.User, error) {
	var userResponse response.User

	user, err := u.UserRepo.GetUserByID(ID)
	if err != nil {
		if err.Error() == "user not found" {
			return userResponse, fmt.Errorf("user not found with ID: %d", ID)
		}
		return userResponse, fmt.Errorf("failed to get user by ID: %w", err)
	}

	userResponse = response.User{
		ID:          user.ID,
		Name:        user.Name,
		Username:    user.Username,
		Email:       user.Email,
		Gender:      user.Gender,
		PhoneNumber: user.PhoneNumber,
		Picture:     user.Picture,
	}

	return userResponse, nil
}

func (u *userUseCase) GetUserByUsername(username string) (response.User, error) {
	var userResponse response.User

	user, err := u.UserRepo.GetUserByUsername(username)
	if err != nil {
		return userResponse, err
	}

	userResponse = response.User{
		ID:          user.ID,
		Name:        user.Name,
		Username:    user.Username,
		Email:       user.Email,
		Gender:      user.Gender,
		PhoneNumber: user.PhoneNumber,
		Picture:     user.Picture,
	}

	return userResponse, nil
}

func (u *userUseCase) GetUserByEmail(email string) (response.User, error) {
	var userResponse response.User

	user, err := u.UserRepo.GetUserByEmail(email)
	if err != nil {
		return userResponse, err
	}

	userResponse = response.User{
		ID:          user.ID,
		Name:        user.Name,
		Username:    user.Username,
		Email:       user.Email,
		Gender:      user.Gender,
		PhoneNumber: user.PhoneNumber,
		Picture:     user.Picture,
	}

	return userResponse, nil
}

func (u *userUseCase) UpdateUser(ID uint, request request.UpdateUser) (response.User, error) {
	var userResponse response.User

	_, err := u.UserRepo.GetUserByUsername(request.Username)
	if err == nil {
		return userResponse, errors.New("username already used")
	}

	_, err = u.UserRepo.GetUserByEmail(request.Email)
	if err == nil {
		return userResponse, errors.New("email already used")
	}

	// Encryption Password from string to Bcrypt
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return userResponse, errors.New("failed to encrypt password")
	}

	// Transfer object request to model User
	updateUser := model.User{
		Name:        request.Name,
		Username:    request.Username,
		Email:       request.Email,
		Password:    string(hashPassword),
		Gender:      request.Gender,
		PhoneNumber: request.PhoneNumber,
		Picture:     request.Picture,
	}

	err = u.UserRepo.UpdateUser(ID, updateUser)
	if err != nil {
		return userResponse, err
	}

	// Transfer Model User to response User
	userResponse = response.User{
		ID:          ID,
		Name:        updateUser.Name,
		Username:    updateUser.Username,
		Email:       updateUser.Email,
		Gender:      updateUser.Gender,
		PhoneNumber: updateUser.PhoneNumber,
		Picture:     updateUser.Picture,
	}

	return userResponse, nil
}

func (u *userUseCase) DeleteUser(ID uint) error {
	if err := u.UserRepo.DeleteUserByID(ID); err != nil {
		return err
	}

	return nil
}
