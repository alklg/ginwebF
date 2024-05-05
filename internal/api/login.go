package api

import (
	"crypto/rand"
	"fmt"
	"ginweb/internal/db"
	"ginweb/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"time"
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

	// above will generate a signingKey
	// fmt.Println("this is signingKey: ", signingKey)

	err := c.ShouldBindJSON(&user)

	if err != nil {
		log.Println("login err ", err)
	}

	// generate user object, include user.username, user.password
	log.Println("this is", user)

	//examine whether user info match with database
	code, err := db.ExamineUserLogin(user.Username, user.Password)
	fmt.Printf("this is code: %d err: %v", code, err)

	if err != nil {
		log.Printf("in login code = %d, err = %v\n", code, err)

		c.JSON(http.StatusOK, gin.H{
			"data":   "Error",
			"code":   code,
			"source": "PostLoginHandler"})
	}

	if code == 4 {
		// user message
		// send user message

		uid := db.ParseUserMessage(&user)

		// generate jwt
		signingKey := make([]byte, 32)

		_, err := rand.Read(signingKey)
		if err != nil {
			log.Println("Error generating random key:", err)
			return
		}

		Claims := jwt.MapClaims{
			"iss": "FurudoErika",
			"sub": uid,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
		signedToken, err := token.SignedString(signingKey)

		if err != nil {
			log.Println(" generate token err = %v ", err)
			return
		}

		c.Header("Authorization", "Bearer "+signedToken)
		c.JSON(http.StatusOK, gin.H{
			"data":   signedToken,
			"code":   "200",
			"status": "success",
			"source": "PostLoginHandler",
		})
	}
}
