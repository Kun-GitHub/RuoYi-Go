// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Author: K.
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package tests

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
	"time-machine/pkg/cache"
)

func TestRedis(t *testing.T) {
	// 创建redisStruct实例
	redisService := &cache.RedisStruct{
		Options: &redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		},
	}

	redisService.NewClient()

	// 在这里使用redisService.rdb执行Redis操作，例如：
	setErr := redisService.Set("key", "value", 0)
	if setErr != nil {
		fmt.Println("Failed to set key-value pair:", setErr)
	} else {
		val, getErr := redisService.Get("key")
		if getErr == nil {
			fmt.Println("Retrieved value:", val)
		} else {
			fmt.Println("Failed to get value:", getErr)
		}
	}

	defer func() {
		// 完成操作后，关闭Redis连接
		closeErr := redisService.Close()
		if closeErr != nil {
			fmt.Println("Failed to close Redis connection:", closeErr)
		}
	}()
}
