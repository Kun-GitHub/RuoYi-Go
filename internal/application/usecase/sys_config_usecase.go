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

type SysConfigService struct {
	repo   output.SysConfigRepository
	cache  *cache.FreeCacheClient
	logger *zap.Logger
}

func NewSysConfigService(repo output.SysConfigRepository, cache *cache.FreeCacheClient, logger *zap.Logger) input.SysConfigService {
	return &SysConfigService{repo: repo, cache: cache, logger: logger}
}

func (this *SysConfigService) QueryConfigByID(id int64) (*model.SysConfig, error) {
	structEntity := &model.SysConfig{}
	// 尝试从缓存中获取
	userBytes, err := this.cache.Get([]byte(fmt.Sprintf("ConfigID:%d", id)))
	if err == nil {
		// 缓存命中
		err = json.Unmarshal(userBytes, &structEntity)
		if err == nil && structEntity.ConfigID != 0 {
			// 缓存命中
			return structEntity, nil
		}
	}

	structEntity, err = this.repo.QueryConfigByID(id)
	if err != nil {
		this.logger.Error("查询部门信息失败", zap.Error(err))
		return nil, err
	} else if structEntity.ConfigID != 0 {
		// 序列化用户对象并存入缓存
		userBytes, err = json.Marshal(structEntity)
		if err == nil {
			this.cache.Set([]byte(fmt.Sprintf("ConfigID:%d", structEntity.ConfigID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
		}
		return structEntity, nil
	}

	this.logger.Debug("查询信息失败", zap.Error(err))
	return nil, fmt.Errorf("查询信息失败", zap.Error(err))
}

func (this *SysConfigService) QueryConfigList(post *model.SysConfigRequest) ([]*model.SysConfig, error) {
	structEntity := make([]*model.SysConfig, 0)

	structEntity, err := this.repo.QueryConfigList(post)
	if err != nil {
		this.logger.Error("查询信息失败", zap.Error(err))
		return nil, err
	} else {
		return structEntity, nil
	}
}

func (this *SysConfigService) QueryConfigPage(pageReq common.PageRequest, r *model.SysConfigRequest) ([]*model.SysConfig, int64, error) {
	data, total, err := this.repo.QueryConfigPage(pageReq, r)
	if err != nil {
		this.logger.Error("查询角色分页信息失败", zap.Error(err))
		return nil, 0, err
	}

	return data, total, nil
}

func (this *SysConfigService) AddConfig(post *model.SysConfig) (*model.SysConfig, error) {
	data, err := this.repo.AddConfig(post)
	if err != nil {
		this.logger.Error("AddConfig", zap.Error(err))
		return nil, err
	}
	if data != nil && data.ConfigID != 0 {
		// 序列化用户对象并存入缓存
		userBytes, err := json.Marshal(data)
		if err == nil {
			this.cache.Set([]byte(fmt.Sprintf("ConfigID:%d", data.ConfigID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
		}
	}
	return data, nil
}

func (this *SysConfigService) EditConfig(post *model.SysConfig) (*model.SysConfig, int64, error) {
	data, result, err := this.repo.EditConfig(post)
	if err != nil {
		this.logger.Error("AddConfig", zap.Error(err))
		return nil, 0, err
	}
	if data != nil && data.ConfigID != 0 && result == 1 {
		// 序列化用户对象并存入缓存
		userBytes, err := json.Marshal(data)
		if err == nil {
			this.cache.Set([]byte(fmt.Sprintf("ConfigID:%d", data.ConfigID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
		}
	}
	return data, result, nil
}

func (this *SysConfigService) DeleteConfigById(id int64) (int64, error) {
	result, err := this.repo.DeleteConfigById(id)
	if err != nil {
		this.logger.Error("删除用户信息失败", zap.Error(err))
		return 0, err
	}
	if result == 1 {
		this.cache.Del(fmt.Sprintf("ConfigId:%d", id))
	}
	return result, nil
}

func (this *SysConfigService) CheckConfigNameUnique(id int64, name string) (int64, error) {
	result, err := this.repo.CheckConfigNameUnique(id, name)
	if err != nil {
		this.logger.Error("CheckConfigNameUnique", zap.Error(err))
		return -1, err
	}
	return result, nil
}
