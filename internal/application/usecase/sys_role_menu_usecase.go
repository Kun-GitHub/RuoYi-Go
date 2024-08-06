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

type SysRoleMenuService struct {
	repo   output.SysRoleMenuRepository
	cache  *cache.FreeCacheClient
	logger *zap.Logger
}

func NewSysRoleMenuService(repo output.SysRoleMenuRepository, cache *cache.FreeCacheClient, logger *zap.Logger) input.SysRoleMenuService {
	return &SysRoleMenuService{repo: repo, cache: cache, logger: logger}
}

func (this *SysRoleMenuService) AddRoleMenu(post *model.SysRoleMenu) (*model.SysRoleMenu, error) {
	data, err := this.repo.AddRoleMenu(post)
	if err != nil {
		this.logger.Error("AddRoleMenu", zap.Error(err))
		return nil, err
	}
	return data, nil
}

func (this *SysRoleMenuService) DeleteRoleMenuByRoleId(roleId int64) (int64, error) {
	result, err := this.repo.DeleteRoleMenuByRoleId(roleId)
	if err != nil {
		this.logger.Error("删除用户信息失败", zap.Error(err))
		return 0, err
	}
	if result == 1 {
		//this.cache.Del(fmt.Sprintf("UserRoles:%d", userId))
	}
	return result, nil
}
