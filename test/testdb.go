package test

import (
	"fmt"
	"ginweb/internal/db"
)

func TestUserLogin() {
	_, err := db.ExamineUserLogin("lkl", "liu050202")
	if err != nil {
		fmt.Println("err ", err)
		return
	}
}
