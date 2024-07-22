// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package usecase

import (
	"RuoYi-Go/internal/domain/model"
	"RuoYi-Go/internal/ports/input"
	"RuoYi-Go/internal/ports/output"
	"RuoYi-Go/pkg/cache"
	"encoding/json"
	"fmt"
	"github.com/coocood/freecache"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"time"
)

type DemoService struct {
	repo   output.DemoRepository
	redis  *cache.RedisClient
	cache  *freecache.Cache
	logger *zap.Logger
}

func NewDemoService(repo output.DemoRepository, redis *cache.RedisClient, cache *freecache.Cache, logger *zap.Logger) input.DemoService {
	return &DemoService{repo: repo, redis: redis, cache: cache, logger: logger}
}

func (s *DemoService) GetDemoByID(id uint) (*model.Demo, error) {
	s.logger.Info("getting demo by ID", zap.Uint("id", id))

	cacheKey := fmt.Sprintf("demo_%d", id)
	if data, err := s.cache.Get([]byte(cacheKey)); err == nil {
		var demo model.Demo
		if err := json.Unmarshal(data, &demo); err == nil {
			return &demo, nil
		}
	}

	demo, err := s.repo.GetDemoByID(id)
	if err != nil {
		return nil, err
	}

	data, _ := json.Marshal(demo)
	s.cache.Set([]byte(cacheKey), data, 300) // Cache for 5 minutes

	return demo, nil
}

func (s *DemoService) GenerateRandomCode() (string, error) {
	code := uuid.New().String()
	err := s.redis.Set("random_code", code, time.Minute*5)
	if err != nil {
		s.logger.Error("failed to set code in redis", zap.Error(err))
		return "", err
	}

	s.logger.Info("generated random code", zap.String("code", code))
	return code, nil
}
