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

type SysDeptService struct {
	repo   output.SysDeptRepository
	cache  *cache.FreeCacheClient
	logger *zap.Logger
}

func NewSysDeptService(repo output.SysDeptRepository, cache *cache.FreeCacheClient, logger *zap.Logger) input.SysDeptService {
	return &SysDeptService{repo: repo, cache: cache, logger: logger}
}

func (this *SysDeptService) QueryRolesByDeptId(deptId int64) (*model.SysDept, error) {
	structEntity := &model.SysDept{}
	// 尝试从缓存中获取
	userBytes, err := this.cache.Get([]byte(fmt.Sprintf("DeptID:%d", deptId)))
	if err == nil {
		// 缓存命中
		err = json.Unmarshal(userBytes, &structEntity)
		if err == nil && structEntity.DeptID != 0 {
			// 缓存命中
			return structEntity, nil
		}
	}

	structEntity, err = this.repo.QueryRolesByDeptId(deptId)
	if err != nil {
		this.logger.Error("查询用户部门信息失败", zap.Error(err))
		return nil, err
	} else {
		// 序列化用户对象并存入缓存
		userBytes, err = json.Marshal(structEntity)
		if err == nil && structEntity.DeptID != 0 {
			this.cache.Set([]byte(fmt.Sprintf("DeptID:%d", structEntity.DeptID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
			return structEntity, nil
		}
	}

	this.logger.Debug("查询用户部门信息失败", zap.Error(err))
	return nil, fmt.Errorf("查询用户部门信息失败", zap.Error(err))
}

func (this *SysDeptService) QueryDeptList(dept *model.SysDept) ([]*model.SysDept, error) {
	structEntity := make([]*model.SysDept, 0)

	structEntity, err := this.repo.QueryDepts(dept)
	if err != nil {
		this.logger.Error("查询信息失败", zap.Error(err))
		return nil, err
	} else {
		return structEntity, nil
	}

	this.logger.Debug("查询信息失败", zap.Error(err))
	return nil, fmt.Errorf("查询信息失败", zap.Error(err))
}
