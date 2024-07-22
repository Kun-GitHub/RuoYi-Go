// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package config

import (
	"RuoYi-Go/config"
	"github.com/spf13/viper"
)

func LoadConfig() (config.AppConfig, error) {
	var config config.AppConfig

	viper.SetConfigName("config")
	//viper.SetConfigName("demo")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	//viper.AutomaticEnv() //将环境变量与配置绑定
	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}
	return config, nil
}
