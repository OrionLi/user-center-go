package handlers

import (
	"context"
	"fmt"
	"user-center-go/models"
	pb "user-center-go/proto/userpb"
	"user-center-go/services" // 替换为你的service文件所在的包路径
)

// 只提供grpc服务

type Server struct {
	UserService services.UserService
}

func (s *Server) CreateUser(_ context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := &models.User{
		Username: req.GetUsername(),
		Account:  req.GetAccount(),
		Password: req.GetPassword(),
	}

	createdUser, err := s.UserService.CreateUser(user)
	if err != nil {
		return nil, err
	}

	res := &pb.CreateUserResponse{
		Id: uint32(createdUser.ID),
	}

	return res, nil
}

func (s *Server) GetUser(_ context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	userID := req.GetId()

	user, err := s.UserService.GetUserByID(uint(userID))
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	res := &pb.GetUserResponse{
		Id:       uint32(user.ID),
		Username: user.Username,
		Account:  user.Account,
	}

	return res, nil
}
