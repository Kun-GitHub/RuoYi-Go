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

type SysDeptService struct {
	repo   output.SysDeptRepository
	cache  *cache.FreeCacheClient
	logger *zap.Logger
}

func NewSysDeptService(repo output.SysDeptRepository, cache *cache.FreeCacheClient, logger *zap.Logger) input.SysDeptService {
	return &SysDeptService{repo: repo, cache: cache, logger: logger}
}

func (this *SysDeptService) QueryDeptList(dept *model.SysDept) ([]*model.SysDept, error) {
	structEntity := make([]*model.SysDept, 0)

	structEntity, err := this.repo.QueryDeptList(dept)
	if err != nil {
		this.logger.Error("查询信息失败", zap.Error(err))
		return nil, err
	} else {
		return structEntity, nil
	}
}

func (this *SysDeptService) QueryDeptListExcludeById(id int64) ([]*model.SysDept, error) {
	structEntity := make([]*model.SysDept, 0)

	structEntity, err := this.repo.QueryDeptListExcludeById(id)
	if err != nil {
		this.logger.Error("查询信息失败", zap.Error(err))
		return nil, err
	} else {
		return structEntity, nil
	}
}

func (this *SysDeptService) QueryDeptById(id int64) (*model.SysDept, error) {
	structEntity := &model.SysDept{}
	// 尝试从缓存中获取
	userBytes, err := this.cache.Get([]byte(fmt.Sprintf("DeptID:%d", id)))
	if err == nil {
		// 缓存命中
		err = json.Unmarshal(userBytes, &structEntity)
		if err == nil && structEntity.DeptID != 0 {
			// 缓存命中
			return structEntity, nil
		}
	}

	structEntity, err = this.repo.QueryDeptById(id)
	if err != nil {
		this.logger.Error("查询部门信息失败", zap.Error(err))
		return nil, err
	} else if structEntity.DeptID != 0 {
		// 序列化用户对象并存入缓存
		userBytes, err = json.Marshal(structEntity)
		if err == nil {
			this.cache.Set([]byte(fmt.Sprintf("DeptID:%d", structEntity.DeptID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
		}
		return structEntity, nil
	}

	this.logger.Debug("查询信息失败", zap.Error(err))
	return nil, fmt.Errorf("查询信息失败", zap.Error(err))
}

func (this *SysDeptService) AddDept(post *model.SysDept) (*model.SysDept, error) {
	data, err := this.repo.AddDept(post)
	if err != nil {
		this.logger.Error("AddDept", zap.Error(err))
		return nil, err
	}
	if data != nil && data.DeptID != 0 {
		// 序列化用户对象并存入缓存
		userBytes, err := json.Marshal(data)
		if err == nil {
			this.cache.Set([]byte(fmt.Sprintf("DeptID:%d", data.DeptID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
		}
	}
	return data, nil
}

func (this *SysDeptService) EditDept(post *model.SysDept) (*model.SysDept, int64, error) {
	data, result, err := this.repo.EditDept(post)
	if err != nil {
		this.logger.Error("EditDept", zap.Error(err))
		return nil, 0, err
	}
	if data != nil && data.DeptID != 0 && result == 1 {
		// 序列化用户对象并存入缓存
		userBytes, err := json.Marshal(data)
		if err == nil {
			this.cache.Set([]byte(fmt.Sprintf("DeptID:%d", data.DeptID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
		}
	}
	return data, result, nil
}

func (this *SysDeptService) DeleteDeptById(id int64) (int64, error) {
	result, err := this.repo.DeleteDeptById(id)
	if err != nil {
		this.logger.Error("删除用户信息失败", zap.Error(err))
		return 0, err
	}
	if result > 0 {
		this.cache.Del(fmt.Sprintf("DeptID:%d", id))
	}
	return result, nil
}
