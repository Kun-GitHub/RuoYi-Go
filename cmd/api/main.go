// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package main

import (
	ryserver "RuoYi-Go/internal/server"
	"RuoYi-Go/internal/shutdown"
	"RuoYi-Go/internal/websocket"
	"RuoYi-Go/pkg/config"
	"RuoYi-Go/pkg/db"
	"RuoYi-Go/pkg/i18n"
	"RuoYi-Go/pkg/logger"
	"RuoYi-Go/pkg/redis"
	"os"

	"context"
	"fmt"
	"go.uber.org/zap"
	"time"

	"github.com/kataras/iris/v12"
)

func main() {
	// 初始化配置
	conf, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	// 初始化日志
	log := logger.GetLogger() // 假设配置中有Debug字段

	// 初始化国际化
	ryi18n.GetLocalizer(conf.Language) // 假设配置中指定了Language

	// 创建redisStruct实例
	ryredis.GetRedis()

	// 创建DatabaseStruct实例
	rydb.DB = &rydb.DatabaseStruct{}
	err = rydb.DB.OpenSqlite()
	if err != nil {
		log.Error("failed to initialize database,", zap.Error(err))
		os.Exit(0)
	}

	app := iris.New()
	ryserver.StartServer(app)
	rywebsocket.StartWebSocket(app)

	//log.Info("start server on:%d", conf.Server.Port)
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
			logger.GetLogger().Sync()
		},

		// 关闭 sqlService
		func() {
			err = rydb.DB.CloseSqlite()
			if err != nil {
				log.Error("Failed to close the database connection:", zap.Error(err))
			}
		},

		func() {
			// 完成操作后，关闭Redis连接
			err = ryredis.GetRedis().Close()
			if err != nil {
				log.Error("Failed to close Redis connection:", zap.Error(err))
			}
		},
	)
}
