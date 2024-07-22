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
	"go.uber.org/zap"
)

type SysMenuService struct {
	repo   output.SysMenuRepository
	cache  *cache.FreeCacheClient
	logger *zap.Logger
}

func NewSysMenuService(repo output.SysMenuRepository, cache *cache.FreeCacheClient, logger *zap.Logger) input.SysMenuService {
	return &SysMenuService{repo: repo, cache: cache, logger: logger}
}

func (this *SysMenuService) QueryMenusByUserId(userId int64) ([]*model.SysMenu, error) {
	structEntity := make([]*model.SysMenu, 0)

	structEntity, err := this.repo.QueryMenusByUserId(userId)
	if err != nil {
		this.logger.Error("查询用户菜单信息失败", zap.Error(err))
		return nil, err
	} else {
		return structEntity, nil
	}
}
