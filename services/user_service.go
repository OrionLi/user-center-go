package services

import (
	"errors"
	"gorm.io/gorm"
	"user-center-go/database"
	pb "user-center-go/proto/userpb"
)

type UserService interface {
	CreateUser(user *pb.User) (*pb.User, error)
	GetUserByID(id uint) (*pb.User, error)
	GetUserByUsername(username string) (*pb.User, error)
	UpdateUser(user *pb.User) (*pb.User, error)
	DeleteUser(id uint) error
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) CreateUser(user *pb.User) (*pb.User, error) {
	err := database.DB.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetUserByID(id uint) (*pb.User, error) {
	var user pb.User
	err := database.DB.First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 用户不存在
		}
		return nil, err
	}
	return &user, nil
}

func (s *userService) GetUserByUsername(username string) (*pb.User, error) {
	var user pb.User
	err := database.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 用户不存在
		}
		return nil, err
	}
	return &user, nil
}

func (s *userService) UpdateUser(user *pb.User) (*pb.User, error) {
	err := database.DB.Save(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) DeleteUser(id uint) error {
	err := database.DB.Delete(&pb.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
