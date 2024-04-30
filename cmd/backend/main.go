package main

import (
	"go.uber.org/zap"
	"time-machine/internal/app"
	"time-machine/internal/config"
	"time-machine/internal/i18n"
	"time-machine/internal/logger"
	"time-machine/internal/shutdown"
)

func main() {
	// 初始化配置
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	// 初始化日志
	log, err := logger.InitializeLogger(cfg.Debug) // 假设配置中有Debug字段
	if err != nil {
		panic(err)
	}

	// 初始化国际化
	localizer, err := i18n.GetLocalizer(cfg.Language) // 假设配置中指定了Language
	if err != nil {
		log.Fatal("Failed to get localizer", zap.Error(err))
		return
	}

	// 使用日志和国际化开始运行应用
	log.Info("Starting the application...")
	app.Run(cfg, localizer, log) // 将日志实例传递给app.Run

	// 系统关闭
	shutdown.NewHook().Close(
		// 关闭 logger
		func() {
			logger.Close(log)
		},
	)
}
