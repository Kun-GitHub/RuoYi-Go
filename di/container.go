// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package di

import (
	"RuoYi-Go/config"
	"RuoYi-Go/internal/adapters/dao"
	"RuoYi-Go/internal/server"
	"RuoYi-Go/internal/websocket"
	"RuoYi-Go/pkg/cache"
	"RuoYi-Go/pkg/i18n"
	"RuoYi-Go/pkg/logger"
	"context"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go.uber.org/zap"
	"time"
)

type Container struct {
	appConfig config.AppConfig
	logger    *zap.Logger
	redis     *cache.RedisClient
	localizer *i18n.Localizer
	gormDB    *dao.DatabaseStruct
	app       *iris.Application
	freeCache *cache.FreeCacheClient
}

func NewContainer(c config.AppConfig) (*Container, error) {
	// NewZapLogger
	log := logger.NewZapLogger(c.LogLevel)

	// 初始化Redis
	redis, err := cache.NewRedisClient(c, log)
	if err != nil {
		log.Error("failed to connect to redis", zap.Error(err))
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	// 初始化国际化
	l := ryi18n.LoadLocalizer(c.Language) // 假设配置中指定了Language

	// 创建DatabaseStruct实例
	db, err := dao.OpenDB(c)
	if err != nil {
		log.Error("failed to initialize database", zap.Error(err))
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	freeCache := cache.NewFreeCacheClient(100 * 1024 * 1024)

	app := iris.New()
	ms := ryserver.ResolveServerMiddleware(db, redis, log, freeCache, c)
	app.Use(ms.MiddlewareHandler)

	//demoHandler := ryserver.ResolveDemoHandler(redis, cache, log)
	//app.Get("/demos/{id:uint}", demoHandler.GetDemoByID)
	//app.Get("/generate-code", demoHandler.GenerateRandomCode)

	captchaHandler := ryserver.ResolveCaptchaHandler(redis, log)
	app.Get("/captchaImage", captchaHandler.GenerateCaptchaImage)

	authHandler := ryserver.ResolveAuthHandler(db, redis, log, freeCache)
	app.Post("/login", authHandler.Login)
	app.Post("/logout", authHandler.Logout)
	app.Get("/getInfo", authHandler.GetInfo)

	sysMenuHandler := ryserver.ResolveSysMenuHandler(db, log, freeCache)
	app.Get("/getRouters", sysMenuHandler.GetRouters)

	pageSysUserHandler := ryserver.ResolvePageSysUserHandler(db, log, freeCache)
	app.Get("/system/user/list", ms.PermissionMiddleware("system:user:list"), pageSysUserHandler.UserPage)

	sysDictDataHandler := ryserver.ResolveSysDictDataHandler(db, log, freeCache)
	app.Get("/system/dict/data/type/{dictType:string}", sysDictDataHandler.DictType)

	ryws.StartWebSocket(app, log)

	err = app.Run(iris.Addr(fmt.Sprintf(":%d", c.Server.Port)))
	if err != nil {
		log.Error("failed to run http server", zap.Error(err))
		return nil, fmt.Errorf("failed to run http server: %w", err)
	}

	return &Container{
		appConfig: c,
		logger:    log,
		redis:     redis,
		localizer: l,
		gormDB:    db,
		app:       app,
		freeCache: freeCache,
	}, nil
}

func (c *Container) Close() {
	err := c.gormDB.CloseDB()
	if err != nil {
		c.logger.Error("Failed to close the database connection:", zap.Error(err))
	} else {
		c.logger.Info("database closed")
	}

	// 关闭Redis客户端
	if err := c.redis.CloseRedis(); err != nil {
		c.logger.Error("failed to close redis client", zap.Error(err))
	} else {
		c.logger.Info("Redis client closed")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// close all hosts
	if err := c.app.Shutdown(ctx); err != nil {
		c.logger.Error("failed to close all hosts", zap.Error(err))
	} else {
		c.logger.Info("all hosts closed")
	}

	if c.freeCache != nil {
		c.freeCache.Clear()
	}

	// 关闭日志
	c.logger.Sync()
}
