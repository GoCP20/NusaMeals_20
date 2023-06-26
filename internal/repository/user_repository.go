package repository

import (
	"fmt"
	"reglog/internal/model"
	"strings"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	CreateUser(data model.User) error
	GetAllUsers() ([]model.User, error)
	GetUserByID(ID uint) (model.User, error)
	GetUserByUsername(username string) (model.User, error)
	GetUserByEmail(email string) (model.User, error)
	GetUserByEmailOrUsername(emailOrUsername string) (model.User, error)
	UpdateUser(ID uint, data model.User) error
	DeleteUserByID(ID uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(data model.User) error {
	if err := r.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User

	if err := r.db.Model(&model.User{}).Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (r *userRepository) GetUserByID(ID uint) (model.User, error) {
	var user model.User

	if err := r.db.First(&user, ID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			// Record not found
			return user, fmt.Errorf("user not found")
		}
		// Other database error
		return user, err
	}

	return user, nil
}

func (r *userRepository) GetUserByUsername(username string) (model.User, error) {
	var user model.User

	err := r.db.Model(&user).Where("username = ?", username).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) GetUserByEmail(email string) (model.User, error) {
	var user model.User

	if err := r.db.Model(&user).Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) GetUserByEmailOrUsername(emailOrUsername string) (model.User, error) {
	var user model.User

	// Check if the provided value is an email or username
	if strings.Contains(emailOrUsername, "@") {
		// Search by email
		if err := r.db.Model(&user).Where("email = ?", emailOrUsername).First(&user).Error; err != nil {
			return user, err
		}
	} else {
		// Search by username
		if err := r.db.Model(&user).Where("username = ?", emailOrUsername).First(&user).Error; err != nil {
			return user, err
		}
	}

	return user, nil
}

func (r *userRepository) UpdateUser(ID uint, data model.User) error {
	var user model.User

	if err := r.db.Model(&user).Where("id = ?", ID).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) DeleteUserByID(ID uint) error {
	var user model.User

	if err := r.db.Where("id = ", ID).Find(&user).Unscoped().Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
