package test

import (
	"fmt"
	"ginweb/internal/db"
	"testing"
)

func TestUserLogin(t *testing.T) {

	_, err := db.ExamineUserLogin("lkl", "liu050202")
	if err != nil {
		fmt.Println("err ", err)
		return
	}
}
