// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package middlewares

import (
	"RuoYi-Go/config"
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/domain/model"
	"RuoYi-Go/internal/ports/input"
	"RuoYi-Go/pkg/cache"
	"RuoYi-Go/pkg/jwt"
	"fmt"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
	"regexp"
	"strings"
)

type LoginUserStruct struct {
	model.SysUser
	Admin bool             `json:"admin"`
	Roles []*model.SysRole `json:"roles"`
}

var loginUser = &LoginUserStruct{}

type MiddlewareStruct struct {
	redis   *cache.RedisClient
	logger  *zap.Logger
	cfg     config.AppConfig
	service input.SysUserService
}

func NewMiddlewareStruct(r *cache.RedisClient, l *zap.Logger, c config.AppConfig, s input.SysUserService) *MiddlewareStruct {
	return &MiddlewareStruct{
		redis:   r,
		logger:  l,
		cfg:     c,
		service: s,
	}
}

func (this *MiddlewareStruct) MiddlewareHandler(ctx iris.Context) {
	uri := ctx.Request().RequestURI
	// 检查当前请求路径是否在跳过列表中
	if skipInterceptor(uri, this.cfg.Server.NotIntercept) {
		// 如果是，则直接调用Next，跳过此中间件的其余部分
		ctx.Next()
		return
	}

	authorization := ctx.GetHeader(common.AUTHORIZATION)
	if authorization == "" {
		ctx.JSON(common.Error(iris.StatusUnauthorized, "请重新登录"))
		return
	}
	token := authorization
	if strings.Index(authorization, " ") > 0 {
		token = authorization[strings.Index(authorization, " ")+1:]
	}

	jwt_id, err := ryjwt.Valid(common.USER_ID, token)
	if err != nil || jwt_id == "" {
		ctx.JSON(common.Error(iris.StatusUnauthorized, "请重新登录"))
		return
	}
	redis_id, err := this.redis.Get(fmt.Sprintf("%s:%s", common.TOKEN, token))
	if err != nil || redis_id == "" || jwt_id != redis_id {
		ctx.JSON(common.Error(iris.StatusUnauthorized, "请重新登录"))
		return
	}

	sysUser, err := this.service.QueryUserByUserId(jwt_id)
	if err != nil {
		ctx.JSON(common.Error(iris.StatusUnauthorized, "请重新登录"))
		return
	}

	ctx.Values().Set(common.USER_ID, jwt_id)
	ctx.Values().Set(common.TOKEN, token)

	loginUser.SysUser = *sysUser
	if sysUser.UserID == 1 {
		loginUser.Admin = true
	}

	// 继续执行下一个中间件或处理函数
	ctx.Next()
}

func skipInterceptor(path string, notInterceptList []string) bool {
	for _, pattern := range notInterceptList {
		matched, _ := regexp.MatchString(pattern, path)
		if matched {
			return true
		}
	}
	return false
}

func GetLoginUser() *LoginUserStruct {
	return loginUser
}

func ClearLoginUser() {
	loginUser = nil
}

//// 定义一个权限检查函数
//func hasPermission(ctx iris.Context, permission string) bool {
//	if loginUser == nil {
//		return false
//	} else if loginUser.Admin {
//		return true
//	}
//
//	menus, err := service2.QueryMenusByUserId(loginUser.UserID)
//	if err != nil {
//		//logger.Log.Error("QueryMenusByUserId error,", zap.Error(err))
//		return false
//	}
//	for _, menu := range menus {
//		if menu.Perms == permission {
//			return true
//		}
//	}
//	return false
//}
//
//// 定义一个权限检查的中间件
//func PermissionMiddleware(permission string) iris.Handler {
//	return func(ctx iris.Context) {
//		if !hasPermission(ctx, permission) {
//			ctx.StatusCode(http.StatusForbidden)
//			ctx.StopExecution()
//			return
//		}
//		ctx.Next()
//	}
//}
