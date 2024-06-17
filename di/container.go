// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package di

import (
	"RuoYi-Go/config"
	"RuoYi-Go/internal/adapters/handler"
	"RuoYi-Go/internal/adapters/persistence"
	"RuoYi-Go/internal/application/usecase"
	"RuoYi-Go/internal/middlewares"
	ryws "RuoYi-Go/internal/websocket"
	"RuoYi-Go/pkg/cache"
	rydb "RuoYi-Go/pkg/db"
	ryi18n "RuoYi-Go/pkg/i18n"
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
	gormDB    *rydb.DatabaseStruct
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
	db, err := rydb.OpenDB(c)
	if err != nil {
		log.Error("failed to initialize database", zap.Error(err))
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	freeCache := cache.NewFreeCacheClient(100 * 1024 * 1024)

	app := iris.New()
	ms := resolveMiddlewareStruct(db, redis, log, freeCache, c)
	app.Use(ms.MiddlewareHandler)

	//demoHandler := resolveDemoHandler(redis, cache, log)
	//app.Get("/demos/{id:uint}", demoHandler.GetDemoByID)
	//app.Get("/generate-code", demoHandler.GenerateRandomCode)

	captchaHandler := resolveCaptchaHandler(redis, log)
	app.Get("/captchaImage", captchaHandler.GenerateCaptchaImage)

	authHandler := resolveAuthHandler(db, redis, log, freeCache)
	app.Post("/login", authHandler.Login)
	app.Post("/logout", authHandler.Logout)

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

//	func resolveDemoHandler(redis *cache.RedisClient, cache *freecache.Cache, logger *zap.Logger) *handler.DemoHandler {
//		demoRepo := persistence.NewDemoRepository()
//		demoService := usecase.NewDemoService(demoRepo, redis, cache, logger)
//		return handler.NewDemoHandler(demoService, logger)
//	}

func resolveMiddlewareStruct(db *rydb.DatabaseStruct, redis *cache.RedisClient, logger *zap.Logger, cache *cache.FreeCacheClient, appConfig config.AppConfig) *middlewares.MiddlewareStruct {
	sysUserRepo := persistence.NewSysUserRepository(db)
	sysUserService := usecase.NewSysUserService(sysUserRepo, cache, logger)
	return middlewares.NewMiddlewareStruct(redis, logger, appConfig, sysUserService)
}

func resolveCaptchaHandler(redis *cache.RedisClient, logger *zap.Logger) *handler.CaptchaHandler {
	demoService := usecase.NewCaptchaService(redis, logger)
	return handler.NewCaptchaHandler(demoService)
}

func resolveAuthHandler(db *rydb.DatabaseStruct, redis *cache.RedisClient, logger *zap.Logger, cache *cache.FreeCacheClient) *handler.AuthHandler {
	sysUserRepo := persistence.NewSysUserRepository(db)
	sysUserService := usecase.NewSysUserService(sysUserRepo, cache, logger)
	authService := usecase.NewAuthService(sysUserService, redis, logger)
	return handler.NewAuthHandler(authService, logger)
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

	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
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
