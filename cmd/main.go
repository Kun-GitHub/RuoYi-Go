package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
	"go.uber.org/zap"
	"time"
	"RuoYi-Go/internal/shutdown"
	ws "RuoYi-Go/internal/websocket"
	"RuoYi-Go/pkg/config"
	"RuoYi-Go/pkg/i18n"
	"RuoYi-Go/pkg/logger"

	"context"
)

func main() {
	// 初始化配置
	conf, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	// 初始化日志
	log, err := logger.InitializeLogger(conf.Debug) // 假设配置中有Debug字段
	if err != nil {
		panic(err)
	}

	// 初始化国际化
	_, err = i18n.GetLocalizer(conf.Language) // 假设配置中指定了Language
	if err != nil {
		log.Error("Failed to get localizer", zap.Error(err))
	}

	//// 使用日志和国际化开始运行应用
	//app.Run(conf, localizer, log) // 将日志实例传递给app.Run

	app := iris.New()
	// 定义路由
	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Hello, Iris!")
	}) // 定义路由
	app.Get("/captchaImage", func(ctx iris.Context) {
		ctx.WriteString("Hello, Iris!")
	}) // 定义路由
	app.Post("/login", func(ctx iris.Context) {
		ctx.WriteString("Hello, Iris!")
	})

	app.Get("/ws", websocket.Handler(ws.InitWebsocket()))

	app.Run(iris.Addr(fmt.Sprintf(":%d", conf.Server.Port)))

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
			logger.Close()
		},
	)
}
