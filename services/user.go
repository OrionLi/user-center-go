package services

import (
	"errors"
	"gorm.io/gorm"

	"github.com/OrionLi/user-center-go/database"
	"github.com/OrionLi/user-center-go/models"
)

type UserService interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(id uint) error
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) CreateUser(user *models.User) (*models.User, error) {
	err := database.DB.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := database.DB.First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 用户不存在
		}
		return nil, err
	}
	return &user, nil
}

func (s *userService) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 用户不存在
		}
		return nil, err
	}
	return &user, nil
}

func (s *userService) UpdateUser(user *models.User) (*models.User, error) {
	err := database.DB.Save(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) DeleteUser(id uint) error {
	err := database.DB.Delete(&models.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
