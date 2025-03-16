package main

import (
	"fmt"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/kkw-h/gin-template/internal/routes/v1"
	"github.com/kkw-h/gin-template/pkg/config"
	"github.com/kkw-h/gin-template/pkg/global"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	global.Config = cfg

	db := config.InitDatabase(cfg)
	global.DB = db

	// 初始化开发者模式日志（自动彩色输出、栈跟踪）
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(fmt.Errorf("实例化日志失败：%s", err.Error()))
	}
	// 关闭日志（重要！避免资源泄漏）
	defer logger.Sync()
	global.Logger = logger

	r := gin.Default()
	routes.InitRoutesV1(r)

	// 启动服务（支持优雅停机）
	if err := endless.ListenAndServe(cfg.Server.Port, r); err != nil {
		panic(err)
	}
}
