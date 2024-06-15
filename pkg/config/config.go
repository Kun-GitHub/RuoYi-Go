// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package config

import (
	"RuoYi-Go/config"
	"github.com/spf13/viper"
	"sync"
)

var (
	once sync.Once
	conf *config.AppConfig

	App = getConfig()
)

func getConfig() *config.AppConfig {
	once.Do(func() {
		conf = initConfig()
	})
	return conf
}

// InitConfig 函数中使用viper读取配置文件并映射到AppConfig结构体
func initConfig() *config.AppConfig {
	v := viper.New()
	v.SetConfigName("config")
	//v.SetConfigName("demo")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")

	if err := v.ReadInConfig(); err != nil {
		return nil
	}
	if err := v.Unmarshal(&conf); err != nil {
		return nil
	}
	return conf
}
