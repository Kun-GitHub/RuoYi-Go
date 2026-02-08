// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package usecase

import (
	"RuoYi-Go/internal/common"
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
	structEntity, err := this.repo.QueryDictDatasByType(typeStr)
	if err != nil {
		this.logger.Error("QueryDictDatasByType failed", zap.Error(err))
		return nil, err
	}
	return structEntity, nil
}

func (this *SysDictDataService) Get(id uint) (*model.SysDictDatum, error) {
	return this.repo.Get(id)
}

func (this *SysDictDataService) List(page, size int, dictLabel, dictType, status string) ([]*model.SysDictDatum, int64, error) {
	return this.repo.List(page, size, dictLabel, dictType, status)
}

func (this *SysDictDataService) Create(data *model.SysDictDatum) error {
	return this.repo.Create(data)
}

func (this *SysDictDataService) Update(data *model.SysDictDatum) error {
	return this.repo.Update(data)
}

func (this *SysDictDataService) Delete(ids string) error {
	idList := common.SplitInt64(ids)
	return this.repo.Delete(idList)
}
