package global

import (
	"github.com/kkw-h/gin-template/pkg/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	Config *config.AppConfig
	Logger *zap.Logger
)
