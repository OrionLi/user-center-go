package models

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Account  string `json:"account"`
	Password string `json:"password"`
	// 其他字段...
}
