package api

import (
	"fmt"
	"ginweb/internal/db"
	"ginweb/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetLoginHandler(c *gin.Context) {

	//token := jwt.New(jwt.SigningMethodHS256)
	//claims := token.Claims.(jwt.MapClaims)
	//
	//// examine Username
	//claims["username"] = "exampleUser"
}

func PostLoginHandler(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		fmt.Println("login err ", err)
	}

	// generate user object, include user.username, user.password
	fmt.Println("this is", user)

	//examine whether user info match with database
	code, err := db.ExamineUserLogin(user.Username, user.Password)
	fmt.Println("in login code = , err = ", code, err)

	if code == 2 {
		c.JSON(http.StatusOK, gin.H{
			"data": "userNotFound"})
	} else if code == 3 {
		c.JSON(http.StatusOK, gin.H{
			"data": "messageDoesn'tMatch"})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data":    user,
			"message": "test"})
	}

}
