package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"

	pb "user-center-go/proto/userpb" // 替换为你的 proto 文件的包路径

	"google.golang.org/grpc"
)

func main() {
	// 设置 gRPC 服务器的地址和端口
	address := "localhost:50051"

	// 创建与 gRPC 服务器的连接
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// 创建 gRPC 客户端
	client := pb.NewUserServiceClient(conn)

	// 循环获取用户输入的 ID
	for {
		// 获取用户输入的 ID
		var id int64
		fmt.Print("Enter the user ID (0 to exit): ")
		_, err = fmt.Scan(&id)
		if err != nil {
			log.Fatalf("Failed to read user ID: %v", err)
		}

		if id == 0 {
			// 输入 0 表示退出循环
			break
		}

		// 发起 gRPC 调用
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		// 调用 GetUserByID 方法
		getUserByIDRequest := &pb.GetUserRequest{
			Id: uint32(id),
		}
		getUserByIDResponse, err := client.GetUser(ctx, getUserByIDRequest)
		if err != nil {
			log.Fatalf("Failed to get user by ID: %v", err)
		}

		// 处理返回结果...
		log.Printf("User by ID: %v", getUserByIDResponse)
	}
}
