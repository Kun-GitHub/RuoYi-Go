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

// SysDeptService 部门服务实现类
// 负责处理部门相关的业务逻辑，包括部门查询、增删改查、父子关系处理等
type SysDeptService struct {
	repo   output.SysDeptRepository
	cache  *cache.FreeCacheClient
	logger *zap.Logger
}

// NewSysDeptService 创建部门服务实例
// 参数:
//   - repo: 部门仓储接口
//   - cache: 缓存客户端
//   - logger: 日志记录器
//
// 返回值: 部门服务接口
func NewSysDeptService(repo output.SysDeptRepository, cache *cache.FreeCacheClient, logger *zap.Logger) input.SysDeptService {
	return &SysDeptService{repo: repo, cache: cache, logger: logger}
}

// QueryDeptList 查询部门列表
// 根据条件查询部门信息列表
// 参数:
//   - dept: 部门查询条件
//
// 返回值:
//   - []*model.SysDept: 部门列表
//   - error: 错误信息
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

// QueryDeptListExcludeById 查询排除指定ID的部门列表
// 查询所有部门，排除指定ID的部门
// 参数:
//   - id: 要排除的部门ID
//
// 返回值:
//   - []*model.SysDept: 部门列表
//   - error: 错误信息
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

// QueryDeptById 根据ID查询部门信息
// 支持缓存机制，优先从缓存获取，缓存未命中则查询数据库并更新缓存
// 参数:
//   - id: 部门ID
//
// 返回值:
//   - *model.SysDept: 部门信息
//   - error: 错误信息
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

// AddDept 添加新部门
// 创建新部门记录，并将部门信息缓存
// 参数:
//   - post: 部门信息
//
// 返回值:
//   - *model.SysDept: 创建的部门信息
//   - error: 错误信息
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

// EditDept 编辑部门信息
// 更新现有部门的信息，并同步更新缓存
// 参数:
//   - post: 更新后的部门信息
//
// 返回值:
//   - *model.SysDept: 更新后的部门信息
//   - int64: 影响的行数
//   - error: 错误信息
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

// DeleteDeptById 根据ID删除部门
// 逻辑删除部门，并清除相关缓存
// 参数:
//   - id: 部门ID
//
// 返回值:
//   - int64: 影响的行数
//   - error: 错误信息
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

// QueryChildIdListById 查询子部门ID列表
// 根据父部门ID查询所有子孙部门的ID列表
// 参数:
//   - id: 父部门ID
//
// 返回值:
//   - []int64: 子部门ID列表
//   - error: 错误信息
func (this *SysDeptService) QueryChildIdListById(id int64) ([]int64, error) {
	result, err := this.repo.QueryChildIdListById(id)
	if err != nil {
		this.logger.Error("查询用户信息失败", zap.Error(err))
		return nil, err
	}
	return result, nil
}
