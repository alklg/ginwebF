package handlers

import (
	"ginweb/internal/api"
	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func GetRegisterHandler() {
	r.GET("/register", api.GetRegisterHandler)
}

func Register() {
	r.GET("/register", api.GetRegisterHandler)
	r.PUT("/register", api.PutRegisterHandler)
	r.POST("/register", api.PostRegisterHandler)
}
