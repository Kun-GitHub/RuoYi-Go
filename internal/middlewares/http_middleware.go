// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package middlewares

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/models"
	"RuoYi-Go/internal/responses"
	"RuoYi-Go/internal/services"
	"RuoYi-Go/pkg/config"
	"RuoYi-Go/pkg/jwt"
	ryredis "RuoYi-Go/pkg/redis"
	"fmt"
	"github.com/kataras/iris/v12"
	"regexp"
	"strings"
)

type LoginUserStruct struct {
	models.SysUser
	Admin bool `json:"admin"`
}

var loginUser = &LoginUserStruct{}

func MiddlewareHandler(ctx iris.Context) {
	uri := ctx.Request().RequestURI

	// 检查当前请求路径是否在跳过列表中
	if skipInterceptor(uri, config.Conf.Server.NotIntercept) {
		// 如果是，则直接调用Next，跳过此中间件的其余部分
		ctx.Next()
		return
	}

	authorization := ctx.GetHeader(common.AUTHORIZATION)
	if authorization == "" {
		ctx.JSON(responses.Error(iris.StatusUnauthorized, "请重新登录"))
		return
	}
	token := authorization[strings.Index(authorization, " ")+1:]

	jwt_id, err := ryjwt.Valid(common.USER_ID, token)
	if err != nil || jwt_id == "" {
		ctx.JSON(responses.Error(iris.StatusUnauthorized, "请重新登录"))
		return
	}
	redis_id, err := ryredis.Redis.Get(fmt.Sprintf("%s:%s", common.TOKEN, token))
	if err != nil || redis_id == "" || jwt_id != redis_id {
		ctx.JSON(responses.Error(iris.StatusUnauthorized, "请重新登录"))
		return
	}

	sysUser := &models.SysUser{}
	ctx.Values().Set(common.USER_ID, jwt_id)
	if err := services.QueryUserByUserId(jwt_id, sysUser); err != nil {
		ctx.JSON(responses.Error(iris.StatusUnauthorized, "请重新登录"))
		return
	}

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