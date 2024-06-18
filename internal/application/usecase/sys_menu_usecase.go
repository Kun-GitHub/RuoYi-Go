// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package usecase

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/domain/model"
	"RuoYi-Go/internal/ports/input"
	"RuoYi-Go/internal/ports/output"
	"RuoYi-Go/pkg/cache"
	"encoding/json"
	"fmt"
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
	// 尝试从缓存中获取
	userBytes, err := this.cache.Get([]byte(fmt.Sprintf("UserMenuID:%d", userId)))
	if err == nil {
		// 缓存命中
		err = json.Unmarshal(userBytes, &structEntity)
		if err == nil && userId != 0 {
			// 缓存命中
			return structEntity, nil
		}
	}

	structEntity, err = this.repo.QueryMenusByUserId(userId)
	if err != nil {
		this.logger.Error("查询用户菜单信息失败", zap.Error(err))
		return nil, err
	} else {
		// 序列化用户对象并存入缓存
		userBytes, err = json.Marshal(structEntity)
		if err == nil && userId != 0 {
			this.cache.Set([]byte(fmt.Sprintf("UserMenuID:%d", userId)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
			return structEntity, nil
		}
	}

	this.logger.Debug("查询用户菜单信息失败", zap.Error(err))
	return nil, fmt.Errorf("查询用户菜单信息失败", zap.Error(err))
}
