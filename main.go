package main

import (
	"fmt"
	"ginweb/internal/api"
	"ginweb/internal/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	db.Test()

	//var user = models.User{
	//	Username: "user1",
	//	Password: "123456",
	//	Email:    "thisis@email.com",
	//}
	port := ":8080"
	r := gin.Default()

	r.Static("/static", "./vueo")
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://119.45.192.148:8000"},    // 允许的前端应用地址
		AllowMethods: []string{"GET", "POST", "OPTIONS"},        // 允许的 HTTP 方法
		AllowHeaders: []string{"Authorization", "Content-Type"}, // 允许的请求头字段
	}))

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

	photoSolveGroup := r.Group("/photoSolve")
	{
		photoSolveGroup.GET("", api.GetPhotoSolve)
		photoSolveGroup.POST("", api.PostPhotoSolve)
	}

	err := r.Run(port)

	if err != nil {
		fmt.Println("run error ", err)
	}
}
