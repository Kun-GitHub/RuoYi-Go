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

type SysDictTypeService struct {
	repo   output.SysDictTypeRepository
	cache  *cache.FreeCacheClient
	logger *zap.Logger
}

func NewSysDictTypeService(repo output.SysDictTypeRepository, cache *cache.FreeCacheClient, logger *zap.Logger) input.SysDictTypeService {
	return &SysDictTypeService{repo: repo, cache: cache, logger: logger}
}

func (this *SysDictTypeService) QueryDictTypeByDictID(dictID int64) (*model.SysDictType, error) {
	structEntity := &model.SysDictType{}
	// 尝试从缓存中获取
	userBytes, err := this.cache.Get([]byte(fmt.Sprintf("DictID:%d", dictID)))
	if err == nil {
		// 缓存命中
		err = json.Unmarshal(userBytes, &structEntity)
		if err == nil && structEntity.DictID != 0 {
			// 缓存命中
			return structEntity, nil
		}
	}

	structEntity, err = this.repo.QueryDictTypeByDictID(dictID)
	if err != nil {
		this.logger.Error("查询部门信息失败", zap.Error(err))
		return nil, err
	} else if structEntity.DictID != 0 {
		// 序列化用户对象并存入缓存
		userBytes, err = json.Marshal(structEntity)
		if err == nil {
			this.cache.Set([]byte(fmt.Sprintf("DictID:%d", structEntity.DictID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
		}
		return structEntity, nil
	}

	this.logger.Debug("查询信息失败", zap.Error(err))
	return nil, fmt.Errorf("查询信息失败", zap.Error(err))
}

func (this *SysDictTypeService) QueryDictTypeList(post *model.SysDictTypeRequest) ([]*model.SysDictType, error) {
	structEntity := make([]*model.SysDictType, 0)

	structEntity, err := this.repo.QueryDictTypeList(post)
	if err != nil {
		this.logger.Error("查询信息失败", zap.Error(err))
		return nil, err
	} else {
		return structEntity, nil
	}
}

func (this *SysDictTypeService) QueryDictTypePage(pageReq common.PageRequest, r *model.SysDictTypeRequest) ([]*model.SysDictType, int64, error) {
	data, total, err := this.repo.QueryDictTypePage(pageReq, r)
	if err != nil {
		this.logger.Error("查询角色分页信息失败", zap.Error(err))
		return nil, 0, err
	}

	return data, total, nil
}

func (this *SysDictTypeService) AddDictType(post *model.SysDictType) (*model.SysDictType, error) {
	data, err := this.repo.AddDictType(post)
	if err != nil {
		this.logger.Error("AddDictType", zap.Error(err))
		return nil, err
	}
	if data != nil && data.DictID != 0 {
		// 序列化用户对象并存入缓存
		userBytes, err := json.Marshal(data)
		if err == nil {
			this.cache.Set([]byte(fmt.Sprintf("DictID:%d", data.DictID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
		}
	}
	return data, nil
}

func (this *SysDictTypeService) EditDictType(post *model.SysDictType) (*model.SysDictType, int64, error) {
	data, result, err := this.repo.EditDictType(post)
	if err != nil {
		this.logger.Error("AddDictType", zap.Error(err))
		return nil, 0, err
	}
	if data != nil && data.DictID != 0 && result == 1 {
		// 序列化用户对象并存入缓存
		userBytes, err := json.Marshal(data)
		if err == nil {
			this.cache.Set([]byte(fmt.Sprintf("DictID:%d", data.DictID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
		}
	}
	return data, result, nil
}

func (this *SysDictTypeService) DeleteDictTypeById(id int64) (int64, error) {
	result, err := this.repo.DeleteDictTypeById(id)
	if err != nil {
		this.logger.Error("删除用户信息失败", zap.Error(err))
		return 0, err
	}
	if result == 1 {
		this.cache.Del(fmt.Sprintf("DictTypeId:%d", id))
	}
	return result, nil
}

func (this *SysDictTypeService) CheckDictTypeUnique(id int64, name string) (int64, error) {
	result, err := this.repo.CheckDictTypeUnique(id, name)
	if err != nil {
		this.logger.Error("CheckDictTypeUnique", zap.Error(err))
		return -1, err
	}
	return result, nil
}
