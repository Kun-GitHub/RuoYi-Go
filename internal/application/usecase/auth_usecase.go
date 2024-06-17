// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package usecase

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/domain/model"
	"RuoYi-Go/internal/ports/input"
	"RuoYi-Go/pkg/cache"
	ryjwt "RuoYi-Go/pkg/jwt"
	"fmt"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type AuthService struct {
	service input.SysUserService
	redis   *cache.RedisClient
	logger  *zap.Logger
}

func NewAuthService(service input.SysUserService, redis *cache.RedisClient, logger *zap.Logger) input.AuthService {
	return &AuthService{service: service, redis: redis, logger: logger}
}

func (this *AuthService) Login(l model.LoginRequest) (*model.LoginSuccess, error) {
	v, err := this.redis.Get(fmt.Sprintf("%s:%v", common.CAPTCHA, l.Uuid))
	if err != nil || v == "" {
		return nil, fmt.Errorf("验证码错误或已失效")
	}
	this.redis.Del(fmt.Sprintf("%s:%v", common.CAPTCHA, l.Uuid))

	if strings.EqualFold(v, l.Code) {
		sysUser := &model.SysUser{}
		sysUser, err = this.service.QueryUserByUserName(l.Username)
		if err != nil {
			return nil, err
		}
		if sysUser.UserID == 0 {
			return nil, fmt.Errorf("用户名或密码错误")
		}

		if err = bcrypt.CompareHashAndPassword([]byte(sysUser.Password), []byte(l.Password)); err != nil {
			return nil, fmt.Errorf("用户名或密码错误", zap.Error(err))
		}

		var token = ""
		token, err = ryjwt.Sign(common.USER_ID, fmt.Sprintf("%d", sysUser.UserID), 72)
		if err != nil {
			this.logger.Error("生成token失败", zap.Error(err))
			return nil, fmt.Errorf("生成token失败", zap.Error(err))
		} else {
			this.redis.Set(fmt.Sprintf("%s:%s", common.TOKEN, token), sysUser.UserID, 72*time.Hour)

			loginSuccess := &model.LoginSuccess{
				Code:    common.SUCCESS,
				Token:   token,
				Message: "操作成功",
			}
			return loginSuccess, nil
		}
	} else {
		return nil, fmt.Errorf("验证码错误或已失效")
	}
}

func (this *AuthService) Logout(token string) error {
	if err := this.redis.Del(token); err != nil {
		this.logger.Debug("redis del error", zap.Error(err))
		return err
	}
	return nil
}
