// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package ryserver

import (
	"RuoYi-Go/internal/adapters/handler"
	"RuoYi-Go/internal/middlewares"
	"github.com/kataras/iris/v12"
)

//var HttServer *iris.Application

func StartServer(s *iris.Application) {
	//HttServer = s

	//s.Use(middlewares.MiddlewareHandler)

	s.Get("/getInfo", handler.GetInfo)
	s.Get("/getRouters", handler.GetRouters)

	s.Post("/login", handler.Login)
	s.Post("/logout", handler.Logout)

	s.Get("/system/user/list", middlewares.PermissionMiddleware("system:user:list"), handler.UserList)

}
