// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package logger

import (
	"RuoYi-Go/pkg/config"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogStruct struct {
}

var (
	once sync.Once
	this *zap.Logger
)

func GetLogger() *zap.Logger {
	once.Do(func() {
		logger := LogStruct{}
		this = logger.InitializeLogger()
	})
	return this
}

// InitializeLogger 初始化zap日志实例，支持按日期和大小滚动日志文件
func (this *LogStruct) InitializeLogger() *zap.Logger {
	// lumberjack配置
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "./logs/app.log", // 日志文件路径
		MaxSize:    100,              // 单个日志文件最大大小（单位：MB）
		MaxBackups: 3,                // 保留旧文件的最大数量
		MaxAge:     28,               // 旧文件保留最大天数
		Compress:   true,             // 是否压缩旧文件
	}

	// 自定义zap的encoder配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// 构建zap.Core
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // 使用JSON格式编码日志
		zapcore.NewMultiWriteSyncer( // 同时写入多个地方：控制台和文件
			zapcore.AddSync(os.Stdout),        // 输出到控制台
			zapcore.AddSync(lumberjackLogger), // 输出到文件，使用lumberjack进行日志分割管理
		),
		zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zap.InfoLevel // 设置日志级别
		}),
	)

	var zaplogger *zap.Logger
	// 根据debug标志创建logger实例
	if config.Conf.Debug {
		zaplogger = zap.New(core, zap.Development(), zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	} else {
		zaplogger = zap.New(core, zap.AddCaller())
	}

	return zaplogger
}
