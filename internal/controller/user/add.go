package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kkw-h/gin-template/internal/models"
	"github.com/kkw-h/gin-template/pkg/global"
)

func AddUser(ctx *gin.Context) {
	var userMod models.User
	global.DB.First(&userMod)
	fmt.Println(userMod)

	ctx.JSON(200, gin.H{
		"message": "add user",
	})
}
