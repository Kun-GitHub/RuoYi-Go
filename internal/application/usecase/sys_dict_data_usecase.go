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
	"fmt"
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
		this.logger.Error("查询用户部门信息失败", zap.Error(err))
		return nil, err
	} else {
		//// 序列化用户对象并存入缓存
		//userBytes, err = json.Marshal(structEntity)
		//if err == nil && structEntity.DeptID != 0 {
		//	this.cache.Set([]byte(fmt.Sprintf("DeptID:%d", structEntity.DeptID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
		return structEntity, nil
		//}
	}

	this.logger.Debug("查询用户部门信息失败", zap.Error(err))
	return nil, fmt.Errorf("查询用户部门信息失败", zap.Error(err))
}
