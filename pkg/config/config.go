package config

import "github.com/spf13/viper"

var Conf *AppConfig

// InitConfig 函数中使用viper读取配置文件并映射到AppConfig结构体
func InitConfig() (*AppConfig, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := v.Unmarshal(&Conf); err != nil {
		return nil, err
	}

	return Conf, nil
}
