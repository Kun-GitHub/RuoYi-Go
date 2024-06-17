// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package handler

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/ports/input"
	"github.com/kataras/iris/v12"
	"time"
)

type CaptchaHandler struct {
	service input.CaptchaService
}

func NewCaptchaHandler(service input.CaptchaService) *CaptchaHandler {
	return &CaptchaHandler{service: service}
}

// GenerateCaptchaImage
func (h *CaptchaHandler) GenerateCaptchaImage(ctx iris.Context) {
	c, err := h.service.GenerateCaptchaImage()
	if err != nil {
		ctx.JSON(common.Error(iris.StatusInternalServerError, "生成验证码失败"))
		return
	}
	// 获取当前时间并格式化为HTTP日期格式
	currentTime := time.Now().UTC().Format(time.DateTime)
	ctx.Header("Date", currentTime) // 设置Date头
	ctx.Header("Cache-Control", "no-store, no-cache, must-revalidate")
	ctx.Header("Cache-Control", "post-check=0, pre-check=0")
	ctx.Header("Pragma", "no-cache")
	ctx.ContentType("image/jpeg")

	ctx.JSON(c)
}
