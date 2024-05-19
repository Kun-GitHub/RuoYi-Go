package ryserver

import (
	"RuoYi-Go/internal/handlers"
	"github.com/kataras/iris/v12"
)

func StartServer(s *iris.Application) {
	// 定义路由
	s.Get("/", func(ctx iris.Context) {
		ctx.WriteString("")
	})

	s.Get("/captchaImage", handler.CaptchaImage)
	s.Post("/login", handler.Login)
}
