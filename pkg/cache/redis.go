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

// RedisClient 是Redis客户端的封装结构体
// 提供线程安全的Redis操作接口
type RedisClient struct {
	client *redis.Client

	mu sync.Mutex
}

// NewRedisClient 创建Redis客户端实例
// 根据配置信息初始化Redis连接并进行连通性测试
//
// 参数:
//
//	cfg: 应用配置，包含Redis连接信息
//	logger: 日志记录器
//
// 返回值:
//
//	*RedisClient: Redis客户端实例
//	error: 连接错误信息
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

// Get 获取指定键的值
// Redis GET命令的封装，当键不存在时返回redis.Nil错误
//
// 参数:
//
//	key: 要获取的键名
//
// 返回值:
//
//	string: 键对应的值
//	error: 错误信息（键不存在时返回redis.Nil）
func (rs *RedisClient) Get(key string) (string, error) {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	return rs.client.Get(context.Background(), key).Result()
}

// Set 设置键值对
// Redis SET命令的封装，支持设置过期时间
//
// 参数:
//
//	key: 键名
//	value: 要存储的值
//	expiration: 过期时间
//
// 返回值:
//
//	error: 操作错误信息
func (rs *RedisClient) Set(key string, value interface{}, expiration time.Duration) error {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	return rs.client.Set(context.Background(), key, value, expiration).Err()
}

// SetNotTime 设置键值对（1小时过期）
// Redis SET命令的封装，默认设置1小时过期时间
//
// 参数:
//
//	key: 键名
//	value: 要存储的值
//
// 返回值:
//
//	error: 操作错误信息
func (rs *RedisClient) SetNotTime(key string, value interface{}) error {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	return rs.client.Set(context.Background(), key, value, time.Hour*1).Err()
}

// Del 删除指定键
// Redis DEL命令的封装
//
// 参数:
//
//	key: 要删除的键名
//
// 返回值:
//
//	error: 操作错误信息
func (rs *RedisClient) Del(key string) error {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	return rs.client.Del(context.Background(), key).Err()
}

// CloseRedis 关闭Redis连接
// 释放Redis客户端资源
//
// 返回值:
//
//	error: 关闭连接时的错误信息
func (rs *RedisClient) CloseRedis() error {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	if rs.client != nil {
		// 关闭连接
		return rs.client.Close()
	}
	return nil
}

// Keys 根据模式匹配获取键列表
// Redis KEYS命令的封装
//
// 参数:
//
//	pattern: 匹配模式
//
// 返回值:
//
//	[]string: 匹配的键列表
//	error: 操作错误信息
func (rs *RedisClient) Keys(pattern string) ([]string, error) {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	return rs.client.Keys(context.Background(), pattern).Result()
}
