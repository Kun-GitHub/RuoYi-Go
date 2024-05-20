// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Author: K.
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package handler

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/model"
	"RuoYi-Go/internal/response"
	rydb "RuoYi-Go/pkg/db"
	ryjwt "RuoYi-Go/pkg/jwt"
	ryredis "RuoYi-Go/pkg/redis"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/kataras/iris/v12"
	"strings"
)

func Login(ctx iris.Context) {
	var l loginStruct
	// Attempt to read and bind the JSON request body to the 'user' variable
	if err := ctx.ReadJSON(&l); err != nil {
		ctx.JSON(response.ErrorFormat(iris.StatusBadRequest, "Invalid JSON, error:%s", err.Error()))
		return
	}

	v, error := ryredis.Redis.Get(fmt.Sprintf("%s:%d", common.CAPTCHA, l.Uuid))
	if error != nil || v == "" {
		ctx.JSON(response.Error(iris.StatusInternalServerError, "验证码错误或已失效"))
		return
	}

	if v != "" && strings.EqualFold(v, l.Code) {
		sysUser := &model.SysUser{}

		if err := rydb.DB.FindColumns(model.TableNameSysUser, sysUser, "login_name = ? and status = '0'", l.Username); err != nil {
			ctx.JSON(response.Error(iris.StatusInternalServerError, "用户名或密码错误"))
			return
		}

		if sysUser.UserID == 0 || sysUser.Password != encryptPassword(sysUser.LoginName, l.Password, sysUser.Salt) {
			ctx.JSON(response.Error(iris.StatusInternalServerError, "账号或密码错误"))
			return
		}

		token, error := ryjwt.Sign(common.USER_ID, fmt.Sprintf("%d", sysUser.UserID), 72)
		if error != nil {
			ctx.JSON(response.Error(iris.StatusInternalServerError, "生成token失败"))
		} else {
			user := loginSuccess{
				Code:    response.SUCCESS,
				Token:   token,
				Message: "操作成功",
			}
			// 使用 ctx.JSON 自动将user序列化为JSON并写入响应体
			ctx.JSON(user)
		}
	} else {
		ctx.JSON(response.Error(iris.StatusInternalServerError, "验证码错误"))
		return
	}
}

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

func encryptPassword(loginName, password, salt string) string {
	// 结合loginName、password和salt
	data := []byte(loginName + password + salt)

	// 计算MD5哈希
	hasher := md5.New()
	hasher.Write(data)
	md5Bytes := hasher.Sum(nil)

	// 将哈希结果转换为十六进制字符串
	md5Hex := hex.EncodeToString(md5Bytes)

	return md5Hex
}

func GetInfo(ctx iris.Context) {
	v := ctx.Value(common.USER_ID)

	fmt.Println(v)
}
