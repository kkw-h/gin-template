package config

type AppConfig struct {
	Server struct {
		Port  string `yaml:"port"`
		Debug bool   `yaml:"debug"`
	}
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	}
}
