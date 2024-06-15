package di

import (
	"RuoYi-Go/config"
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
	cfg       config.AppConfig
	logger    *zap.Logger
	redis     *cache.RedisClient
	localizer *i18n.Localizer
	db        *rydb.DatabaseStruct
	app       *iris.Application
}

func NewContainer(c config.AppConfig) (*Container, error) {
	// NewZapLogger
	log := logger.NewZapLogger(c.LogLevel)

	// 初始化Redis
	redis, err := cache.NewRedisClient(c, log)
	if err != nil {
		log.Error("failed to connect to cache", zap.Error(err))
		return nil, fmt.Errorf("failed to connect to cache: %w", err)
	}

	// 初始化国际化
	l := ryi18n.LoadLocalizer(c.Language) // 假设配置中指定了Language

	// 创建DatabaseStruct实例
	db, err := rydb.OpenSqlite(c)
	if err != nil {
		log.Error("failed to initialize database", zap.Error(err))
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	app := iris.New()
	ryws.StartWebSocket(app, log)

	err = app.Run(iris.Addr(fmt.Sprintf(":%d", c.Server.Port)))
	if err != nil {
		log.Error("failed to run http server", zap.Error(err))
		return nil, fmt.Errorf("failed to run http server: %w", err)
	}

	return &Container{
		cfg:       c,
		logger:    log,
		redis:     redis,
		localizer: l,
		db:        db,
		app:       app,
	}, nil
}

func (c *Container) Close() {
	err := c.db.CloseSqlite()
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

	// 关闭日志
	c.logger.Sync()
}
