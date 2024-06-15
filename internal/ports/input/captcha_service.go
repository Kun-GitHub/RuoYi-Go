package input

import "RuoYi-Go/internal/domain/model"

// CaptchaService 输入端口接口
type CaptchaService interface {
	GenerateCaptchaImage() (model.CaptchaImage, error)
}
