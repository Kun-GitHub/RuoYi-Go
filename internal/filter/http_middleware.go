// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package filter

import (
	"RuoYi-Go/config"
	"RuoYi-Go/internal/adapters/dao"
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/domain/model"
	"RuoYi-Go/internal/ports/input"
	"RuoYi-Go/pkg/cache"
	ryjwt "RuoYi-Go/pkg/jwt"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
)

var loginUser = &model.UserInfoStruct{}

// ServerMiddleware 服务器中间件
// 负责HTTP请求的认证拦截、权限验证等前置处理
type ServerMiddleware struct {
	redis       *cache.RedisClient
	logger      *zap.Logger
	cfg         config.AppConfig
	service     input.SysUserService
	menuService input.SysMenuService
	db          *dao.DatabaseStruct
}

// NewServerMiddleware 创建服务器中间件实例
// 参数:
//   - db: 数据库结构体
//   - r: Redis客户端
//   - l: 日志记录器
//   - c: 应用配置
//   - s: 用户服务接口
//   - menuService: 菜单服务接口
//
// 返回值: 服务器
// 返回值: 服务器中间件
func NewServerMiddleware(db *dao.DatabaseStruct, r *cache.RedisClient, l *zap.Logger, c config.AppConfig, s input.SysUserService, menuService input.SysMenuService) *ServerMiddleware {
	return &ServerMiddleware{
		db:          db,
		redis:       r,
		logger:      l,
		cfg:         c,
		service:     s,
		menuService: menuService,
	}
}

// MiddlewareHandler 认证中间件处理器
// 验证请求的JWT令牌和Redis会话信息，设置用户上下文
// 参数:
//   - ctx: Iris上下文对象
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
	redis_val, err := this.redis.Get(fmt.Sprintf("%s:%s", common.TOKEN, token))
	if err != nil || redis_val == "" {
		ctx.JSON(common.Error(iris.StatusUnauthorized, "请重新登录"))
		return
	}

	var userOnline model.SysUserOnline
	if err := json.Unmarshal([]byte(redis_val), &userOnline); err != nil {
		// Fallback for backward compatibility or error
		// Try to parse as int if it was old format? No, let's just fail.
		ctx.JSON(common.Error(iris.StatusUnauthorized, "请重新登录"))
		return
	}

	// jwt_id is user_id string
	if fmt.Sprintf("%d", userOnline.UserID) != jwt_id {
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
	this.db.LoginUser(sysUser)
	if sysUser.UserID == common.ADMINID {
		loginUser.Admin = true
	}
	loginUser.Password = ""
	ctx.Values().Set(common.LOGINUSER, loginUser)

	// 继续执行下一个中间件或处理函数
	ctx.Next()

	ctx.Values().Reset()
	this.db.ClearUser()
}

// skipInterceptor 检查路径是否需要跳过拦截
// 根据配置的白名单判断是否跳过认证检查
// 参数:
//   - path: 请求路径
terceptList: 不拦截的路径列表
//
// 返回值: 
//   - notInterceptList: 不拦截的路径列表
// 返回值: bool 是否跳过拦截
func skipInterceptor(path string, notInterceptList []string) bool {
	for _, pattern := range notInterceptList {
		matched, _ := regexp.MatchString(pattern, path)
		if matched {
			return true
		}
	}
	return false
}

// hasPermission 检查用户是否具有指定权限
// 验证当前登录用户是否拥有指定的操作权限
ion: 权
// 参数:
//   - ctx: Iris上下文对象
//   - permission: 权限标识
// 返回值: bool 是否具有权限
func (this *ServerMiddleware) hasPermission(ctx iris.Context, permission string) bool {
	if loginUser == nil {
		return false
	} else if loginUser.Admin {
		return true
	}

	u := &model.SysMenuRequest{
		UserId: loginUser.UserID,
	}
	menus, err := this.menuService.QueryMenuList(u)
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

// PermissionMiddleware 权限检查中间件
ssion: 需要验证的权限标识
//
// 返回值:
// 创建指定权限的检查中间件，用于路由级别的权限控制
// 参数:
//   - permission: 需要验证的权限标识
// 返回值: iris.Handler 权限检查处理器
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
