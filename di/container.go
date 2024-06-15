package di

import (
	"RuoYi-Go/config"
	"RuoYi-Go/pkg/cache"
	rydb "RuoYi-Go/pkg/db"
	ryi18n "RuoYi-Go/pkg/i18n"
	"RuoYi-Go/pkg/logger"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go.uber.org/zap"
)

type Container struct {
	cfg       config.AppConfig
	logger    *zap.Logger
	redis     *cache.RedisClient
	localizer *i18n.Localizer
	db        *rydb.DatabaseStruct
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
	l := ryi18n.GetLocalizer(c.Language) // 假设配置中指定了Language

	// 创建DatabaseStruct实例
	db, err := rydb.OpenSqlite(c)
	if err != nil {
		log.Error("failed to initialize database", zap.Error(err))
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	return &Container{
		cfg:       c,
		logger:    log,
		redis:     redis,
		localizer: l,
		db:        db,
	}, nil
}

func (c *Container) Close() {
	err := c.db.CloseSqlite()
	if err != nil {
		c.logger.Error("Failed to close the database connection:", zap.Error(err))
	}

	// 关闭Redis客户端
	if err := c.redis.CloseRedis(); err != nil {
		c.logger.Error("failed to close redis client", zap.Error(err))
	}
	// 关闭日志
	c.logger.Sync()
}
