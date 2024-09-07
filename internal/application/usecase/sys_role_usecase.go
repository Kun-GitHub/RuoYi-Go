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

	return roles, nil
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

func (this *SysRoleService) QueryRoleByID(id int64) (*model.SysRole, error) {
	structEntity := &model.SysRole{}
	// 尝试从缓存中获取
	userBytes, err := this.cache.Get([]byte(fmt.Sprintf("RoleID:%d", id)))
	if err == nil {
		// 缓存命中
		err = json.Unmarshal(userBytes, &structEntity)
		if err == nil && structEntity.RoleID != 0 {
			// 缓存命中
			return structEntity, nil
		}
	}

	structEntity, err = this.repo.QueryRoleByID(id)
	if err != nil {
		this.logger.Error("QueryRoleByID", zap.Error(err))
		return nil, err
	} else if structEntity.RoleID != 0 {
		if structEntity.RoleID == common.ADMINID {
			structEntity.Admin = true
		}

		// 序列化用户对象并存入缓存
		userBytes, err = json.Marshal(structEntity)
		if err == nil {
			this.cache.Set([]byte(fmt.Sprintf("RoleID:%d", structEntity.RoleID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
		}
		return structEntity, nil
	}

	this.logger.Debug("查询信息失败", zap.Error(err))
	return nil, fmt.Errorf("查询信息失败", zap.Error(err))
}

func (this *SysRoleService) AddRole(post *model.SysRole) (*model.SysRole, error) {
	data, err := this.repo.AddRole(post)
	if err != nil {
		this.logger.Error("AddRole", zap.Error(err))
		return nil, err
	}
	if data != nil && data.RoleID != 0 {
		// 序列化用户对象并存入缓存
		userBytes, err := json.Marshal(data)
		if err == nil {
			this.cache.Set([]byte(fmt.Sprintf("RoleID:%d", data.RoleID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
		}
	}
	return data, nil
}

func (this *SysRoleService) EditRole(post *model.SysRole) (*model.SysRole, int64, error) {
	data, result, err := this.repo.EditRole(post)
	if err != nil {
		this.logger.Error("EditRole", zap.Error(err))
		return nil, 0, err
	}
	if data != nil && data.RoleID != 0 && result == 1 {
		// 序列化用户对象并存入缓存
		userBytes, err := json.Marshal(data)
		if err == nil {
			this.cache.Set([]byte(fmt.Sprintf("RoleID:%d", data.RoleID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
		}
	}
	return data, result, nil
}

func (this *SysRoleService) DeleteRoleById(id int64) (int64, error) {
	result, err := this.repo.DeleteRoleById(id)
	if err != nil {
		this.logger.Error("删除用户信息失败", zap.Error(err))
		return 0, err
	}
	if result > 0 {
		this.cache.Del(fmt.Sprintf("RoleID:%d", id))
	}
	return result, nil
}

func (this *SysRoleService) ChangeRoleStatus(user *model.ChangeRoleStatusRequest) (int64, error) {
	result, err := this.repo.ChangeRoleStatus(user)
	if err != nil {
		this.logger.Error("修改用户状态失败", zap.Error(err))
		return 0, err
	}
	if result == 1 {
		structEntity, err := this.repo.QueryRoleByID(user.RoleId)
		if err != nil {
			this.logger.Error("查询用户信息失败", zap.Error(err))
			return 0, err
		} else if structEntity.RoleID != 0 {
			// 序列化用户对象并存入缓存
			userBytes, err := json.Marshal(structEntity)
			if err == nil {
				this.cache.Set([]byte(fmt.Sprintf("RoleID:%d", structEntity.RoleID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
			}
			return result, nil
		}
	}
	return result, nil
}
