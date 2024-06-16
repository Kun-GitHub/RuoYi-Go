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
