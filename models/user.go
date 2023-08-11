package models

type User struct {
	ID       uint   `gorm:"primary_key"`
	Username string `gorm:"not null"`
	Account  string `gorm:"not null;unique"`
}
