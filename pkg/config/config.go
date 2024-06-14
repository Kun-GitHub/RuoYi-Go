// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package config

import "github.com/spf13/viper"

var Conf *AppConfig

// InitConfig 函数中使用viper读取配置文件并映射到AppConfig结构体
func InitConfig() (*AppConfig, error) {
	v := viper.New()
	v.SetConfigName("config")
	//v.SetConfigName("demo")
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
