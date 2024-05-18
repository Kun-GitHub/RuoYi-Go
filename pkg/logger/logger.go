// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Author: K.
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package logger

import (
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	log *zap.Logger
}

var Log *Logger

// InitializeLogger 初始化zap日志实例，支持按日期和大小滚动日志文件
func InitializeLogger(debug bool) (*Logger, error) {
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

	// 根据debug标志创建logger实例
	var l *zap.Logger
	if debug {
		l = zap.New(core, zap.Development(), zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	} else {
		l = zap.New(core, zap.AddCaller())
	}

	Log = &Logger{
		log: l,
	}
	return Log, nil
}

// Close 关闭zap.Logger实例
func (this *Logger) Close() {
	if this != nil {
		_ = this.log.Sync()
	}
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (this *Logger) Info(format string, a ...any) {
	if this != nil {
		this.log.Info(fmt.Sprintf(format, a...))
	}
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (this *Logger) Error(msg string, fields ...zap.Field) {
	if this != nil {
		this.log.Error(msg, fields...)
	}
}
