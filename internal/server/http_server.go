// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package ryserver

import (
	"RuoYi-Go/config"
	"RuoYi-Go/internal/adapters/handler"
	"RuoYi-Go/internal/adapters/persistence"
	"RuoYi-Go/internal/application/usecase"
	"RuoYi-Go/internal/middlewares"
	"RuoYi-Go/pkg/cache"
	rydb "RuoYi-Go/pkg/db"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
)

func StartServer(s *iris.Application) {
	s.Get("/getRouters", handler.GetRouters)

	//s.Get("/system/user/list", middlewares.PermissionMiddleware("system:user:list"), handler.UserList)
}

//	func ResolveDemoHandler(redis *cache.RedisClient, cache *freecache.Cache, logger *zap.Logger) *handler.DemoHandler {
//		demoRepo := persistence.NewDemoRepository()
//		demoService := usecase.NewDemoService(demoRepo, redis, cache, logger)
//		return handler.NewDemoHandler(demoService, logger)
//	}

func ResolveMiddlewareStruct(db *rydb.DatabaseStruct, redis *cache.RedisClient, logger *zap.Logger, cache *cache.FreeCacheClient, appConfig config.AppConfig) *middlewares.MiddlewareStruct {
	sysUserRepo := persistence.NewSysUserRepository(db)
	sysUserService := usecase.NewSysUserService(sysUserRepo, cache, logger)
	return middlewares.NewMiddlewareStruct(redis, logger, appConfig, sysUserService)
}

func ResolveCaptchaHandler(redis *cache.RedisClient, logger *zap.Logger) *handler.CaptchaHandler {
	demoService := usecase.NewCaptchaService(redis, logger)
	return handler.NewCaptchaHandler(demoService)
}

func ResolveAuthHandler(db *rydb.DatabaseStruct, redis *cache.RedisClient, logger *zap.Logger, cache *cache.FreeCacheClient) *handler.AuthHandler {
	sysUserRepo := persistence.NewSysUserRepository(db)
	sysUserService := usecase.NewSysUserService(sysUserRepo, cache, logger)

	sysRoleRepo := persistence.NewSysRoleRepository(db)
	sysRoleService := usecase.NewSysRoleService(sysRoleRepo, cache, logger)

	sysDeptRepo := persistence.NewSysDeptRepository(db)
	sysDeptService := usecase.NewSysDeptService(sysDeptRepo, cache, logger)

	authService := usecase.NewAuthService(sysUserService, sysRoleService, sysDeptService, redis, logger)
	return handler.NewAuthHandler(authService, logger)
}
