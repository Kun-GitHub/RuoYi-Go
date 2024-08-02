// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package filter

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
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var loginUser = &model.UserInfoStruct{}

type ServerMiddleware struct {
	redis       *cache.RedisClient
	logger      *zap.Logger
	cfg         config.AppConfig
	service     input.SysUserService
	menuService input.SysMenuService
}

func NewServerMiddleware(r *cache.RedisClient, l *zap.Logger, c config.AppConfig, s input.SysUserService, menuService input.SysMenuService) *ServerMiddleware {
	return &ServerMiddleware{
		redis:       r,
		logger:      l,
		cfg:         c,
		service:     s,
		menuService: menuService,
	}
}

func (this *ServerMiddleware) MiddlewareHandler(ctx iris.Context) {
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

	jwtId, err := strconv.ParseInt(jwt_id, 10, 64)
	if err != nil {
		ctx.JSON(common.Error(iris.StatusUnauthorized, "请重新登录"))
		return
	}

	sysUser, err := this.service.QueryUserByUserId(jwtId)
	if err != nil {
		ctx.JSON(common.Error(iris.StatusUnauthorized, "请重新登录"))
		return
	}

	ctx.Values().Set(common.USER_ID, jwt_id)
	ctx.Values().Set(common.TOKEN, token)

	if loginUser == nil {
		loginUser = &model.UserInfoStruct{}
	}
	loginUser.SysUser = sysUser
	if sysUser.UserID == common.ADMINID {
		loginUser.Admin = true
	}
	loginUser.Password = ""
	ctx.Values().Set(common.LOGINUSER, loginUser)

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

// 定义一个权限检查函数
func (this *ServerMiddleware) hasPermission(ctx iris.Context, permission string) bool {
	if loginUser == nil {
		return false
	} else if loginUser.Admin {
		return true
	}

	menus, err := this.menuService.QueryMenusByUserId(loginUser.UserID)
	if err != nil {
		return false
	}
	for _, menu := range menus {
		if menu.Perms == permission {
			return true
		}
	}
	return false
}

// 定义一个权限检查的中间件
func (this *ServerMiddleware) PermissionMiddleware(permission string) iris.Handler {
	return func(ctx iris.Context) {
		if !this.hasPermission(ctx, permission) {
			ctx.StatusCode(http.StatusForbidden)
			ctx.StopExecution()
			return
		}
		ctx.Next()
	}
}
