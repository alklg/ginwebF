package main

import (
	"fmt"
	"ginweb/internal/api"
	"ginweb/internal/db"
	"ginweb/internal/models"
	"ginweb/test"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	db.Test()
	test.TestUserLogin()

	var user = models.User{
		Username: "user1",
		Password: "123456",
		Email:    "thisis@email.com",
	}

	test.TestRegister(user)

	port := ":8080"
	r := gin.Default()

	r.Static("/static", "./vueo")
	r.Use(cors.Default())

	//r.GET("/register", api.GetRegisterHandler)
	// 1 admin(web)                                                2 merchant                                   3 user(android apple )
	registerGroup := r.Group("/register")
	{
		registerGroup.GET("", api.GetRegisterHandler)
		registerGroup.POST("", api.PostRegisterHandler)
		registerGroup.PUT("", api.PutRegisterHandler)
		registerGroup.DELETE("")
	}

	loginGroup := r.Group("/login")
	{
		loginGroup.GET("", api.GetLoginHandler)
		loginGroup.POST("", api.PostLoginHandler)
	}

	err := r.Run(port)

	if err != nil {
		fmt.Println("run error ", err)
	}
}
