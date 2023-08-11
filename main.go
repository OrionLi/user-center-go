package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"user-center-go/database"
	"user-center-go/handlers"
	pb "user-center-go/proto/userpb"
	"user-center-go/services"
)

func main() {
	// 初始化数据库连接
	database.InitSqliteDB()

	// 创建UserService实例
	userService := services.NewUserService()

	// 创建gRPC服务器
	grpcServer := grpc.NewServer()

	// 注册UserService服务
	pb.RegisterUserServiceServer(grpcServer, &handlers.Server{UserService: userService})

	// 监听指定端口
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// 启动gRPC服务器
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
