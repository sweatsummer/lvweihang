package model

type User struct {
	ID       int    `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Recharge int    `json:"recharge"`
}
