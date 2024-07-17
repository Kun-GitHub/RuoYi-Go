// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package usecase

import (
	"RuoYi-Go/internal/domain/model"
	"RuoYi-Go/internal/ports/input"
	"RuoYi-Go/internal/ports/output"
	"RuoYi-Go/pkg/cache"
	"go.uber.org/zap"
)

type SysDictDataService struct {
	repo   output.SysDictDataRepository
	cache  *cache.FreeCacheClient
	logger *zap.Logger
}

func NewSysDictDataService(repo output.SysDictDataRepository, cache *cache.FreeCacheClient, logger *zap.Logger) input.SysDictDataService {
	return &SysDictDataService{repo: repo, cache: cache, logger: logger}
}

func (this *SysDictDataService) QueryDictDatasByType(typeStr string) ([]*model.SysDictDatum, error) {
	structEntity := make([]*model.SysDictDatum, 0)

	structEntity, err := this.repo.QueryDictDatasByType(typeStr)
	if err != nil {
		this.logger.Error("查询信息失败", zap.Error(err))
		return nil, err
	} else {
		return structEntity, nil
	}
}
