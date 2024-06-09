// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package ryredis

import (
	"context"
	"fmt"
	"sync"
	"time"

	"RuoYi-Go/pkg/config"
	"github.com/redis/go-redis/v9"
)

type RedisStruct struct {
	rdb     *redis.Client
	options *redis.Options

	mu sync.Mutex
}

var Redis *RedisStruct

func (rs *RedisStruct) NewClient() error {
	rs.options = &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Conf.Redis.Host, config.Conf.Redis.Port),
		Password: config.Conf.Redis.Password, // no password set
		DB:       config.Conf.Redis.DB,       // use default DB
	}
	rs.rdb = redis.NewClient(rs.options)

	// 使用PING命令检查连接
	_, err := rs.rdb.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	return nil
}

// Get Redis `GET key` command. It returns redis.Nil error when key does not exist.
func (rs *RedisStruct) Get(key string) (string, error) {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	return rs.rdb.Get(context.Background(), key).Result()
}

func (rs *RedisStruct) Set(key string, value interface{}, expiration time.Duration) error {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	return rs.rdb.Set(context.Background(), key, value, expiration).Err()
}

func (rs *RedisStruct) Del(key string) error {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	return rs.rdb.Del(context.Background(), key).Err()
}

func (rs *RedisStruct) Close() error {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	if rs.rdb != nil {
		// 关闭连接
		return rs.rdb.Close()
	}
	return nil
}
