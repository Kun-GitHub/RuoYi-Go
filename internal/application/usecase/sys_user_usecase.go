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
	repo   output.SysUserRepository
	cache  *cache.FreeCacheClient
	logger *zap.Logger
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

func (this *SysUserService) QueryUserPage(pageReq common.PageRequest, userId int64, username string, phone string, status string, deptId int64) (*common.PageResponse, error) {
	data, err := this.repo.QueryUserPage(pageReq, userId, username, phone, status, deptId)
	if err != nil {
		this.logger.Error("查询用户分页信息失败", zap.Error(err))
		return nil, err
	}
	return data, nil
}
