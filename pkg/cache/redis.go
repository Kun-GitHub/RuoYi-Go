// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package cache

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap"

	"RuoYi-Go/config"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client

	mu sync.Mutex
}

func NewRedisClient(cfg config.AppConfig, logger *zap.Logger) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	// Ping Redis to check connection
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		logger.Error("failed to ping redis", zap.Error(err))
		return nil, err
	}

	return &RedisClient{client: client}, nil
}

// Get Redis `GET key` command. It returns redis.Nil error when key does not exist.
func (rs *RedisClient) Get(key string) (string, error) {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	return rs.client.Get(context.Background(), key).Result()
}

func (rs *RedisClient) Set(key string, value interface{}, expiration time.Duration) error {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	return rs.client.Set(context.Background(), key, value, expiration).Err()
}

func (rs *RedisClient) SetNotTime(key string, value interface{}) error {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	return rs.client.Set(context.Background(), key, value, time.Hour*1).Err()
}

func (rs *RedisClient) Del(key string) error {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	return rs.client.Del(context.Background(), key).Err()
}

func (rs *RedisClient) CloseRedis() error {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	if rs.client != nil {
		// 关闭连接
		return rs.client.Close()
	}
	return nil
}

func (rs *RedisClient) Keys(pattern string) ([]string, error) {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	return rs.client.Keys(context.Background(), pattern).Result()
}
