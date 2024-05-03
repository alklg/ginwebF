package db

import (
	"fmt"
	"ginweb/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var user = models.User{}

func TestLink() {
	dsn := "root:liu050202@tcp(127.0.0.1:3306)/ginweb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("link mysql error ", err)
	}

	db.First(&user)

	fmt.Println(user)
}
