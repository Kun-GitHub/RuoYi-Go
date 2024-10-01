// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

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
	service      input.SysUserService
	roleService  input.SysRoleService
	deptService  input.SysDeptService
	loginService input.SysLogininforService
	redis        *cache.RedisClient
	logger       *zap.Logger
}

func NewAuthService(service input.SysUserService, roleService input.SysRoleService, deptService input.SysDeptService, loginService input.SysLogininforService, redis *cache.RedisClient, logger *zap.Logger) input.AuthService {
	return &AuthService{service: service, roleService: roleService, deptService: deptService, loginService: loginService, redis: redis, logger: logger}
}

func (this *AuthService) Login(l *model.LoginRequest) (*model.LoginSuccess, error) {
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
			this.loginService.AddLogininfor(&model.SysLogininfor{
				Status:    "1",
				UserName:  l.Username,
				LoginTime: time.Now(),
			})

			return nil, fmt.Errorf("用户名或密码错误")
		}

		if err = bcrypt.CompareHashAndPassword([]byte(sysUser.Password), []byte(l.Password)); err != nil {
			this.loginService.AddLogininfor(&model.SysLogininfor{
				Status:    "1",
				UserName:  l.Username,
				LoginTime: time.Now(),
			})

			return nil, fmt.Errorf("用户名或密码错误", zap.Error(err))
		}

		var token = ""
		token, err = ryjwt.Sign(common.USER_ID, fmt.Sprintf("%d", sysUser.UserID), 72)
		if err != nil {
			this.loginService.AddLogininfor(&model.SysLogininfor{
				Status:    "1",
				UserName:  l.Username,
				LoginTime: time.Now(),
			})

			this.logger.Error("生成token失败", zap.Error(err))
			return nil, fmt.Errorf("生成token失败", zap.Error(err))
		} else {
			this.redis.Set(fmt.Sprintf("%s:%s", common.TOKEN, token), sysUser.UserID, 72*time.Hour)
			this.service.UserLogin(sysUser)

			this.loginService.AddLogininfor(&model.SysLogininfor{
				Status:    "0",
				UserName:  sysUser.UserName,
				LoginTime: time.Now(),
			})

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

func (this *AuthService) GetInfo(loginUser *model.UserInfoStruct) (*model.UserInfoStruct, []string, []string, error) {
	var p []string
	if loginUser.UserID == common.ADMINID {
		p = append(p, "*:*:*")
	} else {
	}

	roles, err := this.roleService.QueryRolesByUserId(loginUser.UserID)
	if err != nil {
		this.logger.Error("QueryRolesByUserId error,", zap.Error(err))
		return nil, p, nil, fmt.Errorf("getInfo error", zap.Error(err))
	}

	loginUser.Roles = roles
	var roleNames []string
	for _, role := range roles {
		roleNames = append(roleNames, role.RoleKey)
	}

	dept, err := this.deptService.QueryDeptById(loginUser.DeptID)
	if err != nil {
		this.logger.Error("QueryDeptById error,", zap.Error(err))
		return nil, p, roleNames, fmt.Errorf("getInfo error", zap.Error(err))
	}
	loginUser.Dept = dept

	return loginUser, p, roleNames, nil
}
