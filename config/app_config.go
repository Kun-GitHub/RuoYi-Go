// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package config

import "go.uber.org/zap/zapcore"

type AppConfig struct {
	Language string         `mapstructure:"language"` // 应用语言
	Log      LogConfig      `mapstructure:"logInfo"`  // 应用语言
	Server   ServerConfig   `mapstructure:"server"`   // 服务器配置
	Database DatabaseConfig `mapstructure:"database"` // 数据库配置
	Redis    RedisConfig    `mapstructure:"redis"`    // redis配置
}

// ServerConfig 和 DatabaseConfig 也是结构体，分别定义服务器和数据库的相关配置
type ServerConfig struct {
	Port         int      `mapstructure:"port"`
	NotIntercept []string `mapstructure:"notIntercept"`
}

type LogConfig struct {
	LogLevel zapcore.Level `mapstructure:"logLevel"`
	LogPath  string        `mapstructure:"logPath"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	DBtype   string `mapstructure:"dbtype"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	Password string `mapstructure:"password"`
}
