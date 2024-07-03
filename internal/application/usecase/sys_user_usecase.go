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

type SysUserService struct {
	repo     output.SysUserRepository
	roleRepo output.SysRoleRepository
	deptRepo output.SysDeptRepository
	cache    *cache.FreeCacheClient
	logger   *zap.Logger
}

func NewPageSysUserService(repo output.SysUserRepository, roleRepo output.SysRoleRepository, deptRepo output.SysDeptRepository, cache *cache.FreeCacheClient, logger *zap.Logger) input.SysUserService {
	return &SysUserService{repo: repo, roleRepo: roleRepo, deptRepo: deptRepo, cache: cache, logger: logger}
}

func NewSysUserService(repo output.SysUserRepository, cache *cache.FreeCacheClient, logger *zap.Logger) input.SysUserService {
	return &SysUserService{repo: repo, cache: cache, logger: logger}
}

func (this *SysUserService) QueryUserInfoByUserName(username string) (*model.SysUser, error) {
	structEntity := &model.SysUser{}
	// 尝试从缓存中获取
	userBytes, err := this.cache.Get([]byte(fmt.Sprintf("UserName:%s", username)))
	if err == nil {
		// 缓存命中
		err = json.Unmarshal(userBytes, &structEntity)
		if err == nil && structEntity.UserID != 0 {
			// 缓存命中
			return structEntity, nil
		}
	}

	structEntity, err = this.repo.QueryUserInfoByUserName(username)
	if err != nil {
		this.logger.Error("查询用户信息失败", zap.Error(err))
		return nil, err
	} else {
		// 序列化用户对象并存入缓存
		userBytes, err = json.Marshal(structEntity)
		if err == nil && structEntity.UserID != 0 {
			this.cache.Set([]byte(fmt.Sprintf("UserName:%s", username)), userBytes, common.EXPIRESECONDS)          // 第三个参数是过期时间，0表示永不过期
			this.cache.Set([]byte(fmt.Sprintf("UserID:%d", structEntity.UserID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
			return structEntity, nil
		}
	}

	this.logger.Debug("查询用户信息失败", zap.Error(err))
	return nil, fmt.Errorf("查询用户信息失败", zap.Error(err))
}

func (this *SysUserService) QueryUserInfoByUserId(userId string) (*model.SysUser, error) {
	structEntity := &model.SysUser{}
	// 尝试从缓存中获取
	userBytes, err := this.cache.Get([]byte(fmt.Sprintf("UserID:%v", userId)))
	if err == nil {
		// 缓存命中
		err = json.Unmarshal(userBytes, &structEntity)
		if err == nil && structEntity.UserID != 0 {
			// 缓存命中
			return structEntity, nil
		}
	}

	structEntity, err = this.repo.QueryUserInfoByUserId(userId)
	if err != nil {
		this.logger.Error("查询用户信息失败", zap.Error(err))
		return nil, err
	} else {
		// 序列化用户对象并存入缓存
		userBytes, err = json.Marshal(structEntity)
		if err == nil && structEntity.UserID != 0 {
			this.cache.Set([]byte(fmt.Sprintf("UserID:%d", structEntity.UserID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
			return structEntity, nil
		}
	}

	this.logger.Debug("查询用户信息失败", zap.Error(err))
	return nil, fmt.Errorf("查询用户信息失败", zap.Error(err))
}

func (this *SysUserService) QueryUserPage(pageReq common.PageRequest, u *model.SysUser) (*common.PageResponse, error) {
	data, total, err := this.repo.QueryUserPage(pageReq, u)
	if err != nil {
		this.logger.Error("查询用户分页信息失败", zap.Error(err))
		return nil, err
	}

	userList := make([]*model.UserInfoStruct, 0)
	for _, user := range data {
		userInfo := &model.UserInfoStruct{}
		userInfo.SysUser = user
		if user.UserID == common.ADMINID {
			userInfo.Admin = true
		}

		roles, err := this.roleRepo.QueryRolesByUserId(user.UserID)
		if err != nil {
			this.logger.Error("QueryRolesByUserId error,", zap.Error(err))
			return nil, fmt.Errorf("getInfo error", zap.Error(err))
		}
		userInfo.Roles = roles

		dept, err := this.deptRepo.QueryRolesByDeptId(user.DeptID)
		if err != nil {
			this.logger.Error("QueryRolesByDeptId error,", zap.Error(err))
			return nil, fmt.Errorf("getInfo error", zap.Error(err))
		}
		userInfo.Dept = dept
		userList = append(userList, userInfo)
	}

	return &common.PageResponse{
		Total: total,
		Rows:  userList,
	}, nil
}

func (this *SysUserService) QueryUserList(user *model.SysUser) ([]*model.UserInfoStruct, error) {
	return nil, nil
}
