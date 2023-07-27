package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/OrionLi/user-center-go/database"
	"github.com/OrionLi/user-center-go/handlers"
	"github.com/OrionLi/user-center-go/services"
)

func main() {
	database.InitDB()

	userService := services.NewUserService()
	userHandler := handlers.NewUserHandler(userService)

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.POST("/users", userHandler.CreateUser)
		v1.GET("/users/:id", userHandler.GetUserByID)
		// 其他路由...
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
