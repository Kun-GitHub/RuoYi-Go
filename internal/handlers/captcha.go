package handler

import (
	"RuoYi-Go/internal/response"
	"RuoYi-Go/pkg/captcha"
	ryredis "RuoYi-Go/pkg/redis"
	"fmt"
	"github.com/kataras/iris/v12"
	"strings"
	"time"
)

func CaptchaImage(ctx iris.Context) {
	// 获取当前时间并格式化为HTTP日期格式
	currentTime := time.Now().UTC().Format(time.RFC1123)
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

	//logger.Log.Info("验证码: %s", a)
	ryredis.Redis.Set(fmt.Sprintf("captcha:%d", id), a, time.Minute*5)

	user := response.CaptchaImage{
		Code:    200,
		Uuid:    id,
		Img:     b64s[strings.Index(b64s, ",")+1:],
		Message: "操作成功",
	}
	// 使用 ctx.JSON 自动将user序列化为JSON并写入响应体
	ctx.JSON(user)
}
