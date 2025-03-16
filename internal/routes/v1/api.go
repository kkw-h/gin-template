package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kkw-h/gin-template/internal/controller/user"
)

func InitRoutesV1(r *gin.Engine) {
	// 初始化路由
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/user", user.AddUser)
}
