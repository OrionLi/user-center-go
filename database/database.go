package database

import (
	"fmt"
	"github.com/OrionLi/user-center-go/models"
	"gorm.io/gorm/schema"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	USER   = "root"
	PASS   = "123456"
	HOST   = "127.0.0.1"
	PORT   = "3306"
	DBNAME = "db1"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	DB = db

	migrateTables()
}

func migrateTables() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database tables: %v", err)
	}
}
