package di

import (
	"RuoYi-Go/config"
	"RuoYi-Go/pkg/cache"
	"RuoYi-Go/pkg/logger"
	"fmt"
	"go.uber.org/zap"
)

type Container struct {
	cfg    config.AppConfig
	logger *zap.Logger
	redis  *cache.RedisClient
}

func NewContainer(c config.AppConfig) (*Container, error) {
	// NewZapLogger
	log := logger.NewZapLogger(c.LogLevel)

	// 初始化Redis
	redisClient, err := cache.NewRedisClient(c, log)
	if err != nil {
		log.Error("failed to connect to cache", zap.Error(err))
		return nil, fmt.Errorf("failed to connect to cache: %w", err)
	}

	return &Container{
		cfg:    c,
		logger: log,
		redis:  redisClient,
	}, nil
}

func (c *Container) Close() {
	// 关闭Redis客户端
	if err := c.redis.CloseRedis(); err != nil {
		c.logger.Error("failed to close redis client", zap.Error(err))
	}
	// 关闭日志
	c.logger.Sync()
}
