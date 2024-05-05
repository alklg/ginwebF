package test

import (
	"ginweb/internal/db"
	"ginweb/internal/models"
	"testing"
)

func TestRegister(t *testing.T) {
	var user models.User
	db.InsertUserIntoDB(&user)
}

//func TestToken(t *testing.T) {
//
//	reqBody := []byte {
//		{"username": "lkl"},
//			"password": "liu050202"},
//	}
//
//	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(reqBody))
//	var c *gin.Context
//	api.PostLoginHandler(c)
//}
