package global

import (
	"github.com/kkw-h/gin-template/internal/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	Config *config.AppConfig
	Logger *zap.Logger
)
