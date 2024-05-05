package db

import (
	"errors"
	"fmt"
	"ginweb/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "root:liu050202@tcp(127.0.0.1:3306)/ginweb?charset=utf8mb4&parseTime=True&loc=Local"
var db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

func Test() {
	TestLink()
}

func GetUserById(uid int) (models.User, error) {
	var user models.User

	return user, nil
}

// ExamineUserLogin examine if the user exist by username and password.
// this function will return two values, one is error while
// another is code, indicates different situations will happen
// code: 1, link to database failed
// code: 2, user not find
// code: 3, user's message does not match with database
// code: 4, success find and match
func ExamineUserLogin(username, password string) (code int, err error) {

	var user models.User

	// query from table admins find user name = lkl
	// db.Table("admins").Where("name = ?", username).First(&user)

	result := db.Table("users").Where("username = ? AND password = ?", username, password).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// not find user
			fmt.Println("user not find")
			err = errors.New("userNotFound")
			return 2, err
		} else {
			fmt.Println("Error: ", result.Error)
			return 3, result.Error
		}

	}
	fmt.Println("From db find ", user)

	return 4, nil
}

// InsertUserIntoDB receive user message from frontend and curd to database
func InsertUserIntoDB(user *models.User) {
	db.Table("users").Select("Username", "Password", "Email").Create(&user)
	var user2 models.User
	db.Table("users").Where("username = ? AND password = ?", user.Username, user.Password, user.Email).First(&user2)
	fmt.Println(user2)
}

func ParseUserMessage(user *models.User) int {
	var user2 models.User
	db.Table("users").Where("username = ? AND password = ?", user.Username, user.Password).First(&user2)
	user.Uid = user2.Uid

	return user.Uid
}
