package server

import "github.com/kataras/iris/v12"

var server *iris.Application

func InitServer(s *iris.Application) {
	server = s
}

func StartServer() {
	// 定义路由
	server.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Hello, Iris!")
	}) // 定义路由
	server.Get("/captchaImage", func(ctx iris.Context) {
		ctx.WriteString("Hello, Iris!")
	}) // 定义路由
	server.Post("/login", func(ctx iris.Context) {
		ctx.WriteString("Hello, Iris!")
	})
}
