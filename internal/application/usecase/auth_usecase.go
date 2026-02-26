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
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

// AuthService 认证服务实现类
// 负责处理用户登录认证、登出、获取用户信息等核心认证逻辑
type AuthService struct {
	service       input.SysUserService
	roleService   input.SysRoleService
	deptService   input.SysDeptService
	loginService  input.SysLogininforService
	menuService   input.SysMenuService
	redis         *cache.RedisClient
	logger        *zap.Logger
	configService input.SysConfigService
}

// NewAuthService 创建认证服务实例
// 参数:
//   - service: 用户服务接口
//   - roleService: 角色服务接口
//   - deptService: 部门服务接口
//   - configService: 配置服务接口
//   - loginService: 登录信息服务接口
//   - menuService: 菜单服务接口
//   - redis: Redis客户端
//   - logger: 日志记录器
//
// 返回值: 认证服
// 返回值: 认证服务接口
func NewAuthService(service input.SysUserService, roleService input.SysRoleService, deptService input.SysDeptService, configService input.SysConfigService, loginService input.SysLogininforService, menuService input.SysMenuService, redis *cache.RedisClient, logger *zap.Logger) input.AuthService {
	return &AuthService{service: service, roleService: roleService, deptService: deptService, configService: configService, loginService: loginService, menuService: menuService, redis: redis, logger: logger}
}

// Login 用户登录认证
// 处理用户登录请求，包括验证码校验、用户信息验证、密码比对、JWT令牌生成等流程
// 参数:
请求参数，包含用户名、密码、验证码等信息
//
// 返回值:

//   - l: 登录请求参数，包含用户名、密码、验证码等信息
// 返回值:
//   - *model.LoginSuccess: 登录成功响应信息
//   - error: 错误信息
func (this *AuthService) Login(l *model.LoginRequest) (*model.LoginSuccess, error) {
	// 查询验证码是否开启
	captchaEnabled := "true"
	config, err := this.configService.QueryConfigByKey("sys.account.captchaEnabled")
	if err == nil {
		captchaEnabled = config.ConfigValue
	}
	// 如果验证码已开启
	if captchaEnabled == "true" {
		v, err := this.redis.Get(fmt.Sprintf("%s:%v", common.CAPTCHA, l.Uuid))
		if err != nil || v == "" {
			return nil, fmt.Errorf("验证码错误或已失效")
		}
		this.redis.Del(fmt.Sprintf("%s:%v", common.CAPTCHA, l.Uuid))

		if !strings.EqualFold(v, l.Code) {
			return nil, fmt.Errorf("验证码错误或已失效")
		}
	}

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
		// Store Online User Info in Redis
		userOnline := &model.SysUserOnline{
			TokenID:  token, // Or uuid if we want separate token id
			DeptName: "",    // Need to fill this if available
			UserName: sysUser.UserName,
			Ipaddr:   sysUser.LoginIP, // Assuming this was set somewhere? actually DB has it.
			// But Login ip should come from request context...
			// However request context is not passed here efficiently.
			// For now let's use what we have or empty.
			// SysUser from DB has LoginIP which is "last login ip".
			LoginLocation: "",
			Browser:       "",
			Os:            "",
			LoginTime:     time.Now(),
			UserID:        sysUser.UserID,
		}

		// Fill DeptName if possible
		if sysUser.DeptID != 0 {
			dept, err := this.deptService.QueryDeptById(sysUser.DeptID)
			if err == nil && dept != nil {
				userOnline.DeptName = dept.DeptName
			}
		}

		userOnlineJson, _ := json.Marshal(userOnline)
		this.redis.Set(fmt.Sprintf("%s:%s", common.TOKEN, token), string(userOnlineJson), 72*time.Hour)
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
}

 参数:
//   - to
// Logout 用户登出
// 清除用户的Redis会话信息，实现安全退出
// 参数:
//   - token: 用户认证令牌
// 返回值:
//   - error: 错误信息
func (this *AuthService) Logout(token string) error {
	if err := this.redis.Del(token); err != nil {
		this.logger.Debug("redis del error", zap.Error(err))
		return err
	}
	return nil
}


// GetInfo 获取用户详细信息
// 根据登录用户信息获取完整的用户资料，包括角色、权限、部门等信息
// 参数:
//   - loginUser: 已登录的用户基本信息
// 返回值:
//   - *model.UserInfoStruct: 完整的用户信息
//   - []string: 用户权限列表
//   - []string: 用户角色名称列表
//   - error: 错误信息
func (this *AuthService) GetInfo(loginUser *model.UserInfoStruct) (*model.UserInfoStruct, []string, []string, error) {
	var p []string

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

	if loginUser.UserID == common.ADMINID {
		p = append(p, "*:*:*")
	} else {
		// 通过用户id查询权限，并赋值给p（用逗号隔开）
		menus, err := this.menuService.QueryMenuList(&model.SysMenuRequest{
			UserId: loginUser.UserID,
		})
		if err != nil {
			this.logger.Error("查询用户菜单权限失败", zap.Error(err))
			return nil, p, roleNames, err
		}
		for _, menu := range menus {
			if menu.Perms != "" {
				p = append(p, menu.Perms)
			}
		}
	}

	return loginUser, p, roleNames, nil
}
