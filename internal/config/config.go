package config

import "github.com/spf13/viper"

type AppConfig struct {
	Debug    bool           `mapstructure:"debug"`    // 是否开启调试模式
	Language string         `mapstructure:"language"` // 应用语言
	Server   ServerConfig   `mapstructure:"server"`   // 服务器配置
	Database DatabaseConfig `mapstructure:"database"` // 数据库配置
	// 其他配置项...
}

// ServerConfig 和 DatabaseConfig 也是结构体，分别定义服务器和数据库的相关配置
type ServerConfig struct {
	Port int `mapstructure:"port"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

// InitConfig 函数中使用viper读取配置文件并映射到AppConfig结构体
func InitConfig() (*AppConfig, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg AppConfig
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
