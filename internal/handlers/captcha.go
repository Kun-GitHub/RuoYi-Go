// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Author: K.
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package handler

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/response"
	"RuoYi-Go/pkg/captcha"
	"RuoYi-Go/pkg/redis"
	"fmt"
	"github.com/kataras/iris/v12"
	"strings"
	"time"
)

func CaptchaImage(ctx iris.Context) {
	// 获取当前时间并格式化为HTTP日期格式
	currentTime := time.Now().UTC().Format(time.DateTime)
	ctx.Header("Date", currentTime) // 设置Date头
	ctx.Header("Cache-Control", "no-store, no-cache, must-revalidate")
	ctx.Header("Cache-Control", "post-check=0, pre-check=0")
	ctx.Header("Pragma", "no-cache")
	ctx.ContentType("image/jpeg")

	id, b64s, a, err := captcha.GenerateCaptcha()
	if err != nil {
		ctx.JSON(response.Error(iris.StatusInternalServerError, "生成验证码失败"))
		return
	}
	ryredis.Redis.Set(fmt.Sprintf("%s:%d", common.CAPTCHA, id), a, time.Minute*5)

	user := captchaImage{
		Code:    response.SUCCESS,
		Uuid:    id,
		Img:     b64s[strings.Index(b64s, ",")+1:],
		Message: "操作成功",
	}
	// 使用 ctx.JSON 自动将user序列化为JSON并写入响应体
	ctx.JSON(user)
}

type captchaImage struct {
	Code           int    `json:"code"`
	Message        string `json:"msg"`
	Uuid           string `json:"uuid"`
	CaptchaEnabled bool   `json:"captchaEnabled,omitempty"`
	Img            string `json:"img"`
}
