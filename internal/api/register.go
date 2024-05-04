package api

import (
	"fmt"
	"ginweb/internal/db"
	"ginweb/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var r = gin.Default()

func GetRegisterHandler(c *gin.Context) {
	fmt.Println("Get Register Handler Success")
	c.JSON(http.StatusOK, gin.H{
		"message": "Get Registration success",
		"code":    1})
}

func PutRegisterHandler(c *gin.Context) {
	//name := c.PostForm("name")
	//password := c.PostForm("password")
	//email := c.PostForm("email")
	//repassword := c.PostForm("repassword")
}

func PostRegisterHandler(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// return success Page
	c.JSON(http.StatusOK, gin.H{
		"message": "Registration successful",
		"data":    user})

	fmt.Println(user)

	db.InsertUserIntoDB(&user)
}

func LoginUserHandler(c *gin.Context) {

}

// 现在要测试能不能将用户数据插入数据库
func GetUserHandler(uid int) (*models.User, error) {
	user, err := db.GetUserById(uid)

	if err != nil {
		fmt.Println("Get user id error ", err)
	}

	return &models.User{
		Uid:      user.Uid,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
