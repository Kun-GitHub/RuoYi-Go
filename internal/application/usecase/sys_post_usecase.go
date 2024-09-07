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

type SysPostService struct {
	repo   output.SysPostRepository
	cache  *cache.FreeCacheClient
	logger *zap.Logger
}

func NewSysPostService(repo output.SysPostRepository, cache *cache.FreeCacheClient, logger *zap.Logger) input.SysPostService {
	return &SysPostService{repo: repo, cache: cache, logger: logger}
}

func (this *SysPostService) QueryPostByUserId(userId int64) ([]*model.SysPost, error) {
	//TODO 看这里是否要加缓存
	data, err := this.repo.QueryPostByUserId(userId)
	if err != nil {
		this.logger.Error("QueryPostByUserId", zap.Error(err))
		return nil, err
	}

	return data, nil
}

func (this *SysPostService) QueryPostByPostId(postId int64) (*model.SysPost, error) {
	structEntity := &model.SysPost{}
	// 尝试从缓存中获取
	userBytes, err := this.cache.Get([]byte(fmt.Sprintf("PostId:%d", postId)))
	if err == nil {
		// 缓存命中
		err = json.Unmarshal(userBytes, &structEntity)
		if err == nil && structEntity.PostID != 0 {
			// 缓存命中
			return structEntity, nil
		}
	}

	structEntity, err = this.repo.QueryPostByPostId(postId)
	if err != nil {
		this.logger.Error("查询部门信息失败", zap.Error(err))
		return nil, err
	} else if structEntity.PostID != 0 {
		// 序列化用户对象并存入缓存
		userBytes, err = json.Marshal(structEntity)
		if err == nil {
			this.cache.Set([]byte(fmt.Sprintf("PostId:%d", structEntity.PostID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
		}
		return structEntity, nil
	}

	this.logger.Debug("查询信息失败", zap.Error(err))
	return nil, fmt.Errorf("查询信息失败", zap.Error(err))
}

func (this *SysPostService) QueryPostList(post *model.SysPostRequest) ([]*model.SysPost, error) {
	structEntity := make([]*model.SysPost, 0)

	structEntity, err := this.repo.QueryPostList(post)
	if err != nil {
		this.logger.Error("查询信息失败", zap.Error(err))
		return nil, err
	} else {
		return structEntity, nil
	}
}

func (this *SysPostService) QueryPostPage(pageReq common.PageRequest, r *model.SysPostRequest) ([]*model.SysPost, int64, error) {
	data, total, err := this.repo.QueryPostPage(pageReq, r)
	if err != nil {
		this.logger.Error("查询角色分页信息失败", zap.Error(err))
		return nil, 0, err
	}
	return data, total, nil
}

func (this *SysPostService) AddPost(post *model.SysPost) (*model.SysPost, error) {
	data, err := this.repo.AddPost(post)
	if err != nil {
		this.logger.Error("AddPost", zap.Error(err))
		return nil, err
	}
	if data != nil && data.PostID != 0 {
		// 序列化用户对象并存入缓存
		userBytes, err := json.Marshal(data)
		if err == nil {
			this.cache.Set([]byte(fmt.Sprintf("PostId:%d", data.PostID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
		}
	}
	return data, nil
}

func (this *SysPostService) EditPost(post *model.SysPost) (*model.SysPost, int64, error) {
	data, result, err := this.repo.EditPost(post)
	if err != nil {
		this.logger.Error("AddPost", zap.Error(err))
		return nil, 0, err
	}
	if data != nil && data.PostID != 0 && result == 1 {
		// 序列化用户对象并存入缓存
		userBytes, err := json.Marshal(data)
		if err == nil {
			this.cache.Set([]byte(fmt.Sprintf("PostId:%d", data.PostID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
		}
	}
	return data, result, nil
}

func (this *SysPostService) DeletePostById(id int64) (int64, error) {
	result, err := this.repo.DeletePostById(id)
	if err != nil {
		this.logger.Error("删除用户信息失败", zap.Error(err))
		return 0, err
	}
	if result > 0 {
		this.cache.Del(fmt.Sprintf("PostId:%d", id))
	}
	return result, nil
}

func (this *SysPostService) CheckPostNameUnique(id int64, name string) (int64, error) {
	result, err := this.repo.CheckPostNameUnique(id, name)
	if err != nil {
		this.logger.Error("CheckPostNameUnique", zap.Error(err))
		return -1, err
	}
	return result, nil
}

func (this *SysPostService) CheckPostCodeUnique(id int64, code string) (int64, error) {
	result, err := this.repo.CheckPostCodeUnique(id, code)
	if err != nil {
		this.logger.Error("CheckPostCodeUnique", zap.Error(err))
		return -1, err
	}
	return result, nil
}
