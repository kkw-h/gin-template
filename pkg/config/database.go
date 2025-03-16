package config

import (
	"fmt"

	"github.com/kkw-h/gin-template/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase(cfg *config.AppConfig) *gorm.DB {
	mysqlCfg := mysql.Config{
		DSN: fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.Database.User,
			cfg.Database.Password,
			cfg.Database.Host,
			cfg.Database.Port,
			cfg.Database.DBName,
		),
		DefaultStringSize: 256,
	}
	db, err := gorm.Open(mysql.New(mysqlCfg), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("链接数据库失败: %s", err.Error()))
	}
	return db
}
