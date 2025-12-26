// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package usecase

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/domain/model"
	"RuoYi-Go/internal/ports/input"
	"RuoYi-Go/internal/ports/output"
	"RuoYi-Go/pkg/cache"
	"RuoYi-Go/pkg/captcha"
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
)

type CaptchaService struct {
	repo   output.SysConfigRepository
	redis  *cache.RedisClient
	logger *zap.Logger
}

func NewCaptchaService(repo output.SysConfigRepository, r *cache.RedisClient, l *zap.Logger) input.CaptchaService {
	return &CaptchaService{repo: repo, redis: r, logger: l}
}

func (this *CaptchaService) GenerateCaptchaImage() (model.CaptchaImage, error) {
	// 查询验证码是否开启
	captchaEnabled := "true"
	result, err := this.repo.QueryConfigByKey("sys.account.captchaEnabled")
	if err == nil {
		captchaEnabled = result.ConfigValue
	}
	// 如果验证码未开启,直接返回
	if captchaEnabled != "true" {
		return model.CaptchaImage{
			Code:         common.SUCCESS,
			Message:      "操作成功",
			CaptchaEnabled: false,
		}, nil
	}
	
	id, b64s, a, err := captcha.GenerateCaptcha()
	if err != nil {
		this.logger.Error("生成验证码失败", zap.Error(err))
		return model.CaptchaImage{}, err
	}
	this.redis.Set(fmt.Sprintf("%s:%v", common.CAPTCHA, id), a, time.Minute*5)

	c := model.CaptchaImage{
		Code:    common.SUCCESS,
		Uuid:    id,
		CaptchaEnabled: true,
		Img:     b64s[strings.Index(b64s, ",")+1:],
		Message: "操作成功",
	}
	return c, nil
}
