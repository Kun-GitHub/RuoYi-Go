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

type SysLogininforService struct {
	repo   output.SysLogininforRepository
	cache  *cache.FreeCacheClient
	logger *zap.Logger
}

func NewSysLogininforService(repo output.SysLogininforRepository, cache *cache.FreeCacheClient, logger *zap.Logger) input.SysLogininforService {
	return &SysLogininforService{repo: repo, cache: cache, logger: logger}
}

func (this *SysLogininforService) QueryLogininforByID(id int64) (*model.SysLogininfor, error) {
	structEntity := &model.SysLogininfor{}
	// 尝试从缓存中获取
	userBytes, err := this.cache.Get([]byte(fmt.Sprintf("logininfor_InfoID:%d", id)))
	if err == nil {
		// 缓存命中
		err = json.Unmarshal(userBytes, &structEntity)
		if err == nil && structEntity.InfoID != 0 {
			// 缓存命中
			return structEntity, nil
		}
	}

	structEntity, err = this.repo.QueryLogininforByID(id)
	if err != nil {
		this.logger.Error("查询部门信息失败", zap.Error(err))
		return nil, err
	} else if structEntity.InfoID != 0 {
		// 序列化用户对象并存入缓存
		userBytes, err = json.Marshal(structEntity)
		if err == nil {
			this.cache.Set([]byte(fmt.Sprintf("logininfor_InfoID:%d", structEntity.InfoID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
		}
		return structEntity, nil
	}

	this.logger.Debug("查询信息失败", zap.Error(err))
	return nil, fmt.Errorf("查询信息失败", zap.Error(err))
}

func (this *SysLogininforService) QueryLogininforList(post *model.SysLogininforRequest) ([]*model.SysLogininfor, error) {
	structEntity := make([]*model.SysLogininfor, 0)

	structEntity, err := this.repo.QueryLogininforList(post)
	if err != nil {
		this.logger.Error("查询信息失败", zap.Error(err))
		return nil, err
	} else {
		return structEntity, nil
	}
}

func (this *SysLogininforService) QueryLogininforPage(pageReq common.PageRequest, r *model.SysLogininforRequest) ([]*model.SysLogininfor, int64, error) {
	data, total, err := this.repo.QueryLogininforPage(pageReq, r)
	if err != nil {
		this.logger.Error("查询角色分页信息失败", zap.Error(err))
		return nil, 0, err
	}

	return data, total, nil
}

func (this *SysLogininforService) AddLogininfor(post *model.SysLogininfor) (*model.SysLogininfor, error) {
	data, err := this.repo.AddLogininfor(post)
	if err != nil {
		this.logger.Error("AddLogininfor", zap.Error(err))
		return nil, err
	}
	//if data != nil && data.InfoID != 0 {
	//	// 序列化用户对象并存入缓存
	//	userBytes, err := json.Marshal(data)
	//	if err == nil {
	//		this.cache.Set([]byte(fmt.Sprintf("LogininforID:%d", data.InfoID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
	//	}
	//}
	return data, nil
}

func (this *SysLogininforService) EditLogininfor(post *model.SysLogininfor) (*model.SysLogininfor, int64, error) {
	data, result, err := this.repo.EditLogininfor(post)
	if err != nil {
		this.logger.Error("EditLogininfor", zap.Error(err))
		return nil, 0, err
	}
	if data != nil && data.InfoID != 0 && result == 1 {
		// 序列化用户对象并存入缓存
		userBytes, err := json.Marshal(data)
		if err == nil {
			this.cache.Set([]byte(fmt.Sprintf("logininfor_InfoID:%d", data.InfoID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
		}
	}
	return data, result, nil
}

func (this *SysLogininforService) DeleteLogininforById(id int64) (int64, error) {
	result, err := this.repo.DeleteLogininforById(id)
	if err != nil {
		this.logger.Error("删除用户信息失败", zap.Error(err))
		return 0, err
	}
	if result > 0 {
		this.cache.Del(fmt.Sprintf("logininfor_InfoID:%d", id))
	}
	return result, nil
}
