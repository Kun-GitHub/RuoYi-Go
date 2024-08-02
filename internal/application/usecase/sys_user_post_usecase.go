// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package usecase

import (
	"RuoYi-Go/internal/domain/model"
	"RuoYi-Go/internal/ports/input"
	"RuoYi-Go/internal/ports/output"
	"RuoYi-Go/pkg/cache"
	"go.uber.org/zap"
)

type SysUserPostService struct {
	repo   output.SysUserPostRepository
	cache  *cache.FreeCacheClient
	logger *zap.Logger
}

func NewSysUserPostService(repo output.SysUserPostRepository, cache *cache.FreeCacheClient, logger *zap.Logger) input.SysUserPostService {
	return &SysUserPostService{repo: repo, cache: cache, logger: logger}
}

func (this *SysUserPostService) AddUserPost(post *model.SysUserPost) (*model.SysUserPost, error) {
	data, err := this.repo.AddUserPost(post)
	if err != nil {
		this.logger.Error("AddUserPost", zap.Error(err))
		return nil, err
	}
	return data, nil
}

func (this *SysUserPostService) DeleteUserPostByUserId(userId int64) (int64, error) {
	result, err := this.repo.DeleteUserPostByUserId(userId)
	if err != nil {
		this.logger.Error("删除用户信息失败", zap.Error(err))
		return 0, err
	}
	//if result == 1 {
	//	this.cache.Del(fmt.Sprintf("UserPosts:%d", userId))
	//}
	return result, nil
}
