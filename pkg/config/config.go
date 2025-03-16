package config

import (
	"fmt"

	"github.com/kkw-h/gin-template/internal/config"
	"github.com/spf13/viper"
)

var (
	cfg *viper.Viper
	Err error
)

func InitConfig() (*config.AppConfig, error) {
	cfg = viper.New()
	cfg.SetConfigName("config")
	cfg.SetConfigType("yaml")
	cfg.AddConfigPath("./")

	envPrefix := "APP"
	cfg.SetEnvPrefix(envPrefix)
	cfg.AutomaticEnv()
	if err := cfg.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("加载配置文件失败: %w", err)
	}
	var appConfig config.AppConfig
	if err := cfg.Unmarshal(&appConfig); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %w", err)
	}
	return &appConfig, nil
}
