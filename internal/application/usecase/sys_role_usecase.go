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
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
)

type SysRoleService struct {
	repo   output.SysRoleRepository
	cache  *cache.FreeCacheClient
	logger *zap.Logger
}

func NewSysRoleService(repo output.SysRoleRepository, cache *cache.FreeCacheClient, logger *zap.Logger) input.SysRoleService {
	return &SysRoleService{repo: repo, cache: cache, logger: logger}
}

func (this *SysRoleService) QueryRolesByUserId(userId int64) ([]*model.SysRole, error) {
	roles := make([]*model.SysRole, 0)
	// 尝试从缓存中获取
	userBytes, err := this.cache.Get([]byte(fmt.Sprintf("UserRoles:%d", userId)))
	if err == nil {
		// 缓存命中
		err = json.Unmarshal(userBytes, &roles)
		if err == nil && len(roles) != 0 {
			// 缓存命中
			return roles, nil
		}
	}

	roles, err = this.repo.QueryRolesByUserId(userId)
	if err != nil {
		this.logger.Error("查询用户角色信息失败", zap.Error(err))
		return nil, err
	} else if len(roles) != 0 {
		// 序列化用户对象并存入缓存
		userBytes, err = json.Marshal(roles)
		if err == nil && len(roles) != 0 {
			this.cache.Set([]byte(fmt.Sprintf("UserRoles:%d", userId)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
		}
		return roles, nil
	}

	this.logger.Debug("查询用户角色信息失败", zap.Error(err))
	return nil, fmt.Errorf("查询用户角色信息失败", zap.Error(err))
}

func (this *SysRoleService) QueryRolePage(pageReq common.PageRequest, u *model.SysRoleRequest) ([]*model.SysRole, int64, error) {
	data, total, err := this.repo.QueryRolePage(pageReq, u)
	if err != nil {
		this.logger.Error("查询角色分页信息失败", zap.Error(err))
		return nil, 0, err
	}

	return data, total, nil
}

func (this *SysRoleService) QueryRoleList(u *model.SysRoleRequest) ([]*model.SysRole, error) {
	data, err := this.repo.QueryRoleList(u)
	if err != nil {
		this.logger.Error("QueryRoleList", zap.Error(err))
		return nil, err
	}

	return data, nil
}
