package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
	"go.uber.org/zap"
	"time"
	"time-machine/internal/app"
	"time-machine/internal/config"
	"time-machine/internal/i18n"
	"time-machine/internal/logger"
	"time-machine/internal/shutdown"
	ws "time-machine/internal/websocket"

	"context"
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

	app := iris.New()
	// 定义路由
	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Hello, Iris!")
	})

	app.Get("/msg", websocket.Handler(ws.InitWebsocket()))

	app.Run(iris.Addr(":8080"))

	iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		// close all hosts
		app.Shutdown(ctx)
	})

	// 系统关闭
	shutdown.NewHook().Close(
		// 关闭 logger
		func() {
			logger.Close(log)
		},
	)
}
