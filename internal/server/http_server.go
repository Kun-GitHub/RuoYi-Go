package ryserver

import (
	"RuoYi-Go/internal/handlers"
	"github.com/kataras/iris/v12"
)

func StartServer(s *iris.Application) {
	s.Get("/captchaImage", handler.CaptchaImage)
	s.Post("/login", handler.Login)
}
