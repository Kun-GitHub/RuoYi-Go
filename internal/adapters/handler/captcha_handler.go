package handler

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/ports/input"
	"github.com/kataras/iris/v12"
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

	ctx.JSON(c)
}
