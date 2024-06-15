package cache

import (
	"fmt"
	"github.com/coocood/freecache"
)

type FreeCacheClient struct {
	cache *freecache.Cache
}

func NewFreeCacheClient(size int) *FreeCacheClient {
	return &FreeCacheClient{cache: freecache.NewCache(size)}
}

func (c *FreeCacheClient) Set(key string, value interface{}, expireSeconds int) error {
	return c.cache.Set([]byte(key), []byte(fmt.Sprintf("%v", value)), expireSeconds)
}

func (c *FreeCacheClient) Get(key string) (interface{}, error) {
	data, err := c.cache.Get([]byte(key))
	if err != nil {
		return nil, err
	}
	return string(data), nil
}

func (c *FreeCacheClient) Del(key string) bool {
	return c.cache.Del([]byte(key))
}

func (c *FreeCacheClient) Clear() {
	c.cache.Clear()
}
