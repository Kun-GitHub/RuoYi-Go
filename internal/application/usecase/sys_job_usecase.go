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

type SysJobService struct {
	repo   output.SysJobRepository
	cache  *cache.FreeCacheClient
	logger *zap.Logger
}

func NewSysJobService(repo output.SysJobRepository, cache *cache.FreeCacheClient, logger *zap.Logger) input.SysJobService {
	return &SysJobService{repo: repo, cache: cache, logger: logger}
}

func (this *SysJobService) QueryJobByID(id int64) (*model.SysJob, error) {
	structEntity := &model.SysJob{}
	// 尝试从缓存中获取
	userBytes, err := this.cache.Get([]byte(fmt.Sprintf("JobID:%d", id)))
	if err == nil {
		// 缓存命中
		err = json.Unmarshal(userBytes, &structEntity)
		if err == nil && structEntity.JobID != 0 {
			// 缓存命中
			return structEntity, nil
		}
	}

	structEntity, err = this.repo.QueryJobByID(id)
	if err != nil {
		this.logger.Error("查询部门信息失败", zap.Error(err))
		return nil, err
	} else if structEntity.JobID != 0 {
		// 序列化用户对象并存入缓存
		userBytes, err = json.Marshal(structEntity)
		if err == nil {
			this.cache.Set([]byte(fmt.Sprintf("JobID:%d", structEntity.JobID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
		}
		return structEntity, nil
	}

	this.logger.Debug("查询信息失败", zap.Error(err))
	return nil, fmt.Errorf("查询信息失败", zap.Error(err))
}

func (this *SysJobService) QueryJobList(post *model.SysJobRequest) ([]*model.SysJob, error) {
	structEntity := make([]*model.SysJob, 0)

	structEntity, err := this.repo.QueryJobList(post)
	if err != nil {
		this.logger.Error("查询信息失败", zap.Error(err))
		return nil, err
	} else {
		return structEntity, nil
	}
}

func (this *SysJobService) QueryJobPage(pageReq common.PageRequest, r *model.SysJobRequest) ([]*model.SysJob, int64, error) {
	data, total, err := this.repo.QueryJobPage(pageReq, r)
	if err != nil {
		this.logger.Error("查询角色分页信息失败", zap.Error(err))
		return nil, 0, err
	}

	return data, total, nil
}

func (this *SysJobService) AddJob(post *model.SysJob) (*model.SysJob, error) {
	data, err := this.repo.AddJob(post)
	if err != nil {
		this.logger.Error("AddJob", zap.Error(err))
		return nil, err
	}
	//if data != nil && data.InfoID != 0 {
	//	// 序列化用户对象并存入缓存
	//	userBytes, err := json.Marshal(data)
	//	if err == nil {
	//		this.cache.Set([]byte(fmt.Sprintf("JobID:%d", data.InfoID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
	//	}
	//}
	return data, nil
}

func (this *SysJobService) EditJob(post *model.SysJob) (*model.SysJob, int64, error) {
	data, result, err := this.repo.EditJob(post)
	if err != nil {
		this.logger.Error("EditJob", zap.Error(err))
		return nil, 0, err
	}
	if data != nil && data.JobID != 0 && result == 1 {
		// 序列化用户对象并存入缓存
		userBytes, err := json.Marshal(data)
		if err == nil {
			this.cache.Set([]byte(fmt.Sprintf("JobID:%d", data.JobID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
		}
	}
	return data, result, nil
}

func (this *SysJobService) DeleteJobById(id int64) (int64, error) {
	result, err := this.repo.DeleteJobById(id)
	if err != nil {
		this.logger.Error("删除用户信息失败", zap.Error(err))
		return 0, err
	}
	if result > 0 {
		this.cache.Del(fmt.Sprintf("JobID:%d", id))
	}
	return result, nil
}

func (this *SysJobService) ChangeJobStatus(user *model.ChangeJobStatusRequest) (int64, error) {
	result, err := this.repo.ChangeJobStatus(user)
	if err != nil {
		this.logger.Error("修改用户状态失败", zap.Error(err))
		return 0, err
	}
	if result > 0 {
		structEntity, err := this.repo.QueryJobByID(user.JobID)
		if err != nil {
			this.logger.Error("查询用户信息失败", zap.Error(err))
			return 0, err
		} else if structEntity.JobID != 0 {
			// 序列化用户对象并存入缓存
			userBytes, err := json.Marshal(structEntity)
			if err == nil {
				this.cache.Set([]byte(fmt.Sprintf("RoleID:%d", structEntity.JobID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
			}
			return result, nil
		}
	}
	return result, nil
}
