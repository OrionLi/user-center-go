package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primary_key"`
	Username  string `gorm:"not null"`
	Account   string `gorm:"not null;unique"`
	Avatar    string
	Gender    int
	Password  string `gorm:"not null"`
	Phone     string
	Email     string
	ExtraData string `gorm:"type:text"`
	Bio       string
	Role      int
	Status    int
	DeletedAt gorm.DeletedAt
	CreatedAt time.Time
	UpdatedAt time.Time
}
