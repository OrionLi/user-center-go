package handlers

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/anypb"
	pb "user-center-go/proto/userpb"
	"user-center-go/services" // 替换为你的service文件所在的包路径
)

// 只提供grpc服务

type Server struct {
	UserService services.UserService
}

func (s *Server) CreateUser(_ context.Context, req *pb.CreateUserRequest) (*pb.BaseResponse, error) {
	user := &pb.User{
		Username: req.GetUsername(),
		Account:  req.GetAccount(),
		Password: req.GetPassword(),
	}

	createdUser, err := s.UserService.CreateUser(user)
	if err != nil {
		return nil, err
	}

	userResult, _ := anypb.New(createdUser)
	res := &pb.BaseResponse{
		Data: userResult,
	}

	return res, nil
}

func (s *Server) GetUser(_ context.Context, req *pb.GetUserRequest) (*pb.BaseResponse, error) {
	userID := req.GetId()

	user, err := s.UserService.GetUserByID(uint(userID))
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	userResult, _ := anypb.New(user)
	res := &pb.BaseResponse{
		Data: userResult,
	}

	return res, nil
}
