// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Author: K.
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package middlewares

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/responses"
	"RuoYi-Go/pkg/config"
	"RuoYi-Go/pkg/jwt"
	ryredis "RuoYi-Go/pkg/redis"
	"github.com/kataras/iris/v12"
	"regexp"
	"strings"
)

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
	v, err := ryjwt.Valid(common.USER_ID, authorization[strings.Index(authorization, " ")+1:])
	if err != nil || v == "" {
		ctx.JSON(responses.Error(iris.StatusUnauthorized, "请重新登录"))
		return
	}
	v, err = ryredis.Redis.Get(common.TOKEN)
	if err != nil || v == "" {
		ctx.JSON(responses.Error(iris.StatusUnauthorized, "请重新登录"))
		return
	}

	ctx.Values().Set(common.USER_ID, v)
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
