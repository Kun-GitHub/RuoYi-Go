// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package cache

import (
	"github.com/coocood/freecache"
)

type FreeCacheClient struct {
	cache *freecache.Cache
}

func NewFreeCacheClient(size int) *FreeCacheClient {
	return &FreeCacheClient{cache: freecache.NewCache(size)}
}

func (c *FreeCacheClient) Set(key, value []byte, expireSeconds int) (err error) {
	return c.cache.Set(key, value, expireSeconds)
}

func (c *FreeCacheClient) Get(key []byte) (value []byte, err error) {
	return c.cache.Get(key)
}

func (c *FreeCacheClient) Del(key string) bool {
	return c.cache.Del([]byte(key))
}

func (c *FreeCacheClient) Clear() {
	c.cache.Clear()
}
