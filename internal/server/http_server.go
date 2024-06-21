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
	"RuoYi-Go/internal/filter"
	"RuoYi-Go/pkg/cache"
	rydb "RuoYi-Go/pkg/db"
	"go.uber.org/zap"
)

//	func ResolveDemoHandler(redis *cache.RedisClient, cache *freecache.Cache, logger *zap.Logger) *handler.DemoHandler {
//		demoRepo := persistence.NewDemoRepository()
//		demoService := usecase.NewDemoService(demoRepo, redis, cache, logger)
//		return handler.NewDemoHandler(demoService, logger)
//	}

func ResolveServerMiddleware(db *rydb.DatabaseStruct, redis *cache.RedisClient, logger *zap.Logger, cache *cache.FreeCacheClient, appConfig config.AppConfig) *filter.ServerMiddleware {
	sysUserRepo := persistence.NewSysUserRepository(db)
	sysUserService := usecase.NewSysUserService(sysUserRepo, cache, logger)
	sysMenuRepo := persistence.NewSysMenuRepository(db)
	sysMenuService := usecase.NewSysMenuService(sysMenuRepo, cache, logger)
	return filter.NewServerMiddleware(redis, logger, appConfig, sysUserService, sysMenuService)
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

func ResolveSysMenuHandler(db *rydb.DatabaseStruct, logger *zap.Logger, cache *cache.FreeCacheClient) *handler.SysMenuHandler {
	sysMenuRepo := persistence.NewSysMenuRepository(db)
	sysMenuService := usecase.NewSysMenuService(sysMenuRepo, cache, logger)
	return handler.NewSysMenuHandler(sysMenuService)
}

func ResolveSysUserHandler(db *rydb.DatabaseStruct, logger *zap.Logger, cache *cache.FreeCacheClient) *handler.SysUserHandler {
	sysUserRepo := persistence.NewSysUserRepository(db)
	sysUserService := usecase.NewSysUserService(sysUserRepo, cache, logger)
	return handler.NewSysUserHandler(sysUserService)
}
