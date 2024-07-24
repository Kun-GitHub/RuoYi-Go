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

type SysNoticeService struct {
	repo   output.SysNoticeRepository
	cache  *cache.FreeCacheClient
	logger *zap.Logger
}

func NewSysNoticeService(repo output.SysNoticeRepository, cache *cache.FreeCacheClient, logger *zap.Logger) input.SysNoticeService {
	return &SysNoticeService{repo: repo, cache: cache, logger: logger}
}

func (this *SysNoticeService) QueryNoticeByID(id int64) (*model.SysNotice, error) {
	structEntity := &model.SysNotice{}
	// 尝试从缓存中获取
	userBytes, err := this.cache.Get([]byte(fmt.Sprintf("NoticeID:%d", id)))
	if err == nil {
		// 缓存命中
		err = json.Unmarshal(userBytes, &structEntity)
		if err == nil && structEntity.NoticeID != 0 {
			// 缓存命中
			return structEntity, nil
		}
	}

	structEntity, err = this.repo.QueryNoticeByID(id)
	if err != nil {
		this.logger.Error("查询部门信息失败", zap.Error(err))
		return nil, err
	} else if structEntity.NoticeID != 0 {
		// 序列化用户对象并存入缓存
		userBytes, err = json.Marshal(structEntity)
		if err == nil {
			this.cache.Set([]byte(fmt.Sprintf("NoticeID:%d", structEntity.NoticeID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
		}
		return structEntity, nil
	}

	this.logger.Debug("查询信息失败", zap.Error(err))
	return nil, fmt.Errorf("查询信息失败", zap.Error(err))
}

func (this *SysNoticeService) QueryNoticeList(post *model.SysNoticeRequest) ([]*model.SysNotice, error) {
	structEntity := make([]*model.SysNotice, 0)

	structEntity, err := this.repo.QueryNoticeList(post)
	if err != nil {
		this.logger.Error("查询信息失败", zap.Error(err))
		return nil, err
	} else {
		return structEntity, nil
	}
}

func (this *SysNoticeService) QueryNoticePage(pageReq common.PageRequest, r *model.SysNoticeRequest) ([]*model.SysNotice, int64, error) {
	data, total, err := this.repo.QueryNoticePage(pageReq, r)
	if err != nil {
		this.logger.Error("查询角色分页信息失败", zap.Error(err))
		return nil, 0, err
	}

	return data, total, nil
}

func (this *SysNoticeService) AddNotice(post *model.SysNotice) (*model.SysNotice, error) {
	data, err := this.repo.AddNotice(post)
	if err != nil {
		this.logger.Error("AddNotice", zap.Error(err))
		return nil, err
	}
	if data != nil && data.NoticeID != 0 {
		// 序列化用户对象并存入缓存
		userBytes, err := json.Marshal(data)
		if err == nil {
			this.cache.Set([]byte(fmt.Sprintf("NoticeID:%d", data.NoticeID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
		}
	}
	return data, nil
}

func (this *SysNoticeService) EditNotice(post *model.SysNotice) (*model.SysNotice, int64, error) {
	data, result, err := this.repo.EditNotice(post)
	if err != nil {
		this.logger.Error("EditNotice", zap.Error(err))
		return nil, 0, err
	}
	if data != nil && data.NoticeID != 0 && result == 1 {
		// 序列化用户对象并存入缓存
		userBytes, err := json.Marshal(data)
		if err == nil {
			this.cache.Set([]byte(fmt.Sprintf("NoticeID:%d", data.NoticeID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
		}
	}
	return data, result, nil
}

func (this *SysNoticeService) DeleteNoticeById(id int64) (int64, error) {
	result, err := this.repo.DeleteNoticeById(id)
	if err != nil {
		this.logger.Error("删除用户信息失败", zap.Error(err))
		return 0, err
	}
	if result == 1 {
		this.cache.Del(fmt.Sprintf("NoticeID:%d", id))
	}
	return result, nil
}
