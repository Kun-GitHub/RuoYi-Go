// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Author: K.
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package handler

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/middlewares"
	"RuoYi-Go/internal/models"
	"RuoYi-Go/internal/responses"
	"RuoYi-Go/internal/services"
	"RuoYi-Go/pkg/jwt"
	"RuoYi-Go/pkg/redis"
	"fmt"
	"github.com/kataras/iris/v12"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type loginStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Code     string `json:"code"`
	Uuid     string `json:"uuid"`
}

type loginSuccess struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Token   string `json:"token"`
}

func Login(ctx iris.Context) {
	var l loginStruct
	// Attempt to read and bind the JSON request body to the 'user' variable
	if err := ctx.ReadJSON(&l); err != nil {
		ctx.JSON(responses.ErrorFormat(iris.StatusBadRequest, "Invalid JSON, error:%s", err.Error()))
		return
	}

	v, error := ryredis.Redis.Get(fmt.Sprintf("%s:%d", common.CAPTCHA, l.Uuid))
	if error != nil || v == "" {
		ctx.JSON(responses.Error(iris.StatusInternalServerError, "验证码错误或已失效"))
		return
	}
	ryredis.Redis.Del(fmt.Sprintf("%s:%d", common.CAPTCHA, l.Uuid))

	if v != "" && strings.EqualFold(v, l.Code) {
		sysUser := &models.SysUser{}

		if err := services.QueryUserByUserName(l.Username, sysUser); sysUser.UserID == 0 || err != nil {
			ctx.JSON(responses.Error(iris.StatusInternalServerError, "用户名或密码错误"))
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(sysUser.Password), []byte(l.Password)); err != nil {
			ctx.JSON(responses.Error(iris.StatusInternalServerError, "账号或密码错误"))
			return
		}

		token, error := ryjwt.Sign(common.USER_ID, fmt.Sprintf("%d", sysUser.UserID), 72)
		if error != nil {
			ctx.JSON(responses.Error(iris.StatusInternalServerError, "生成token失败"))
		} else {
			ryredis.Redis.Set(fmt.Sprintf("%s:%s", common.TOKEN, token), sysUser.UserID, 72*time.Hour)

			user := loginSuccess{
				Code:    responses.SUCCESS,
				Token:   token,
				Message: "操作成功",
			}
			// 使用 ctx.JSON 自动将user序列化为JSON并写入响应体
			ctx.JSON(user)
		}
	} else {
		ctx.JSON(responses.Error(iris.StatusInternalServerError, "验证码错误"))
		return
	}
}

type getInfoSuccess struct {
	Code        int             `json:"code"`
	Message     string          `json:"msg"`
	Permissions []string        `json:"permissions"`
	User        *models.SysUser `json:"user"`
	Roles       []string        `json:"roles"`
}

func GetInfo(ctx iris.Context) {
	loginUser := middlewares.GetLoginUser()
	if loginUser == nil || loginUser.UserID == 0 {
		ctx.JSON(responses.Error(iris.StatusUnauthorized, "请重新登录"))
		return
	}

	var p []string
	if loginUser.UserID == 1 {
		p = append(p, "*:*:*")
	} else {

	}

	user := getInfoSuccess{
		Code:        responses.SUCCESS,
		User:        loginUser,
		Permissions: p,
		Message:     "操作成功",
	}
	// 使用 ctx.JSON 自动将user序列化为JSON并写入响应体
	ctx.JSON(user)
}
