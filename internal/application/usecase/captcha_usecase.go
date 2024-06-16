package usecase

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/domain/model"
	"RuoYi-Go/internal/ports/input"
	"RuoYi-Go/pkg/cache"
	"RuoYi-Go/pkg/captcha"
	"fmt"
	"go.uber.org/zap"
	"strings"
	"time"
)

type CaptchaService struct {
	redis  *cache.RedisClient
	logger *zap.Logger
}

func NewCaptchaService(r *cache.RedisClient, l *zap.Logger) input.CaptchaService {
	return &CaptchaService{redis: r, logger: l}
}

func (this *CaptchaService) GenerateCaptchaImage() (model.CaptchaImage, error) {
	id, b64s, a, err := captcha.GenerateCaptcha()
	if err != nil {
		this.logger.Error("生成验证码失败", zap.Error(err))
		return model.CaptchaImage{}, err
	}
	this.redis.Set(fmt.Sprintf("%s:%v", common.CAPTCHA, id), a, time.Minute*5)

	c := model.CaptchaImage{
		Code:    common.SUCCESS,
		Uuid:    id,
		Img:     b64s[strings.Index(b64s, ",")+1:],
		Message: "操作成功",
	}
	return c, nil
}
