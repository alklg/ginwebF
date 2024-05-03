package models

type User struct {
	Uid      int    `json:"uid"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
