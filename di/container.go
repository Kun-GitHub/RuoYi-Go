// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package di

import (
	"RuoYi-Go/config"
	"RuoYi-Go/internal/adapters/dao"
	"RuoYi-Go/internal/jobs"
	"RuoYi-Go/internal/server"
	"RuoYi-Go/internal/websocket"
	"RuoYi-Go/pkg/cache"
	"RuoYi-Go/pkg/i18n"
	"RuoYi-Go/pkg/logger"
	"RuoYi-Go/pkg/task"
	"context"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go.uber.org/zap"
	"time"
)

type Container struct {
	appConfig config.AppConfig
	logger    *zap.Logger
	redis     *cache.RedisClient
	localizer *i18n.Localizer
	gormDB    *dao.DatabaseStruct
	app       *iris.Application
	freeCache *cache.FreeCacheClient
	jobs      *task.TaskManager
}

func NewContainer(c config.AppConfig) (*Container, error) {
	// NewZapLogger
	log := logger.NewZapLogger(c)

	// 初始化Redis
	redis, err := cache.NewRedisClient(c, log)
	if err != nil {
		log.Error("failed to connect to redis", zap.Error(err))
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	// 初始化国际化
	l := ryi18n.LoadLocalizer(c.App.Language) // 假设配置中指定了Language

	// 创建DatabaseStruct实例
	db, err := dao.OpenDB(c)
	if err != nil {
		log.Error("failed to initialize database", zap.Error(err))
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	freeCache := cache.NewFreeCacheClient(100 * 1024 * 1024)

	app := iris.New()
	ms := ryserver.ResolveServerMiddleware(db, redis, log, freeCache, c)
	app.Use(ms.MiddlewareHandler)

	//demoHandler := ryserver.ResolveDemoHandler(redis, cache, log)
	//app.Get("/demos/{id:uint}", demoHandler.GetDemoByID)
	//app.Get("/generate-code", demoHandler.GenerateRandomCode)

	captchaHandler := ryserver.ResolveCaptchaHandler(redis, log)
	app.Get("/captchaImage", captchaHandler.GenerateCaptchaImage)

	authHandler := ryserver.ResolveAuthHandler(db, redis, log, freeCache)
	app.Post("/login", authHandler.Login)
	app.Post("/logout", authHandler.Logout)
	app.Get("/getInfo", authHandler.GetInfo)

	sysMenuHandler := ryserver.ResolveSysMenuHandler(db, log, freeCache)
	app.Get("/getRouters", sysMenuHandler.GetRouters)
	app.Get("/system/menu/list", ms.PermissionMiddleware("system:menu:list"), sysMenuHandler.MenuList)
	app.Get("/system/menu/treeselect", sysMenuHandler.TreeSelect)
	app.Get("/system/menu/roleMenuTreeselect/{roleId:uint}", sysMenuHandler.RoleMenuTreeselect)
	app.Get("/system/menu/{menuId:uint}", ms.PermissionMiddleware("system:menu:query"), sysMenuHandler.MenuInfo)
	app.Post("/system/menu", ms.PermissionMiddleware("system:menu:add"), sysMenuHandler.AddMenuInfo)
	app.Put("/system/menu", ms.PermissionMiddleware("system:menu:edit"), sysMenuHandler.EditMenuInfo)
	app.Delete("/system/menu/*menuIds", ms.PermissionMiddleware("system:menu:remove"), sysMenuHandler.DeleteMenuInfo)

	pageSysUserHandler := ryserver.ResolvePageSysUserHandler(db, log, freeCache)
	app.Get("/system/user/", ms.PermissionMiddleware("system:user:query"), pageSysUserHandler.UserInfo)
	app.Get("/system/user/list", ms.PermissionMiddleware("system:user:list"), pageSysUserHandler.UserPage)
	app.Get("/system/user/deptTree", ms.PermissionMiddleware("system:user:list"), pageSysUserHandler.DeptTree)
	app.Get("/system/user/{userId:uint}", ms.PermissionMiddleware("system:user:query"), pageSysUserHandler.UserInfo)
	app.Put("/system/user/changeStatus", ms.PermissionMiddleware("system:user:edit"), pageSysUserHandler.ChangeUserStatus)
	app.Put("/system/user/resetPwd", ms.PermissionMiddleware("system:user:resetPwd"), pageSysUserHandler.ResetUserPwd)
	app.Delete("/system/user/*userIds", ms.PermissionMiddleware("system:user:remove"), pageSysUserHandler.DeleteUser)
	app.Get("/system/user/authRole/{userId:uint}", ms.PermissionMiddleware("system:user:query"), pageSysUserHandler.AuthRoleByUserId)
	app.Put("/system/user/authRole", ms.PermissionMiddleware("system:user:edit"), pageSysUserHandler.AuthRole)
	app.Post("/system/user", ms.PermissionMiddleware("system:user:add"), pageSysUserHandler.AddUser)
	app.Put("/system/user", ms.PermissionMiddleware("system:user:edit"), pageSysUserHandler.EditUser)

	sysDictDataHandler := ryserver.ResolveSysDictDataHandler(db, log, freeCache)
	app.Get("/system/dict/data/type/{dictType:string}", sysDictDataHandler.DictType)

	sysDeptHandler := ryserver.ResolveSysDeptHandler(db, log, freeCache)
	app.Get("/system/dept/list", ms.PermissionMiddleware("system:dept:list"), sysDeptHandler.DeptList)
	app.Get("/system/dept/{deptId:uint}", ms.PermissionMiddleware("system:dept:query"), sysDeptHandler.DeptInfo)
	app.Get("/system/dept/list/exclude/{deptId:uint}", ms.PermissionMiddleware("system:dept:query"), sysDeptHandler.DeptListExcludeById)
	app.Post("/system/dept", ms.PermissionMiddleware("system:dept:add"), sysDeptHandler.AddDeptInfo)
	app.Put("/system/dept", ms.PermissionMiddleware("system:dept:edit"), sysDeptHandler.EditDeptInfo)
	app.Delete("/system/dept/*deptIds", ms.PermissionMiddleware("system:dept:remove"), sysDeptHandler.DeleteDeptInfo)

	sysRoleHandler := ryserver.ResolveSysRoleHandler(db, log, freeCache)
	app.Get("/system/role/list", ms.PermissionMiddleware("system:role:list"), sysRoleHandler.RolePage)
	app.Get("/system/role/{roleId:uint}", ms.PermissionMiddleware("system:role:query"), sysRoleHandler.RoleInfo)
	app.Post("/system/role", ms.PermissionMiddleware("system:role:add"), sysRoleHandler.AddRoleInfo)
	app.Put("/system/role", ms.PermissionMiddleware("system:role:edit"), sysRoleHandler.EditRoleInfo)
	app.Delete("/system/role/*roleIds", ms.PermissionMiddleware("system:role:remove"), sysRoleHandler.DeleteRoleInfo)
	app.Put("/system/role/changeStatus", ms.PermissionMiddleware("system:role:edit"), sysRoleHandler.ChangeRoleStatus)

	sysPostHandler := ryserver.ResolveSysPostHandler(db, log, freeCache)
	app.Get("/system/post/list", ms.PermissionMiddleware("system:post:list"), sysPostHandler.PostPage)
	app.Get("/system/post/{postId:uint}", ms.PermissionMiddleware("system:post:query"), sysPostHandler.PostInfo)
	app.Post("/system/post", ms.PermissionMiddleware("system:post:add"), sysPostHandler.AddPostInfo)
	app.Put("/system/post", ms.PermissionMiddleware("system:post:edit"), sysPostHandler.EditPostInfo)
	app.Delete("/system/post/*postIds", ms.PermissionMiddleware("system:post:remove"), sysPostHandler.DeletePostInfo)

	sysDictTypeHandler := ryserver.ResolveSysDictTypeHandler(db, log, freeCache)
	app.Get("/system/dict/type/list", ms.PermissionMiddleware("system:dict:type:list"), sysDictTypeHandler.DictTypePage)
	app.Get("/system/dict/type/{dictId:uint}", ms.PermissionMiddleware("system:dict:type:query"), sysDictTypeHandler.DictTypeInfo)
	app.Post("/system/dict/type", ms.PermissionMiddleware("system:dict:type:add"), sysDictTypeHandler.AddDictTypeInfo)
	app.Put("/system/dict/type", ms.PermissionMiddleware("system:dict:type:edit"), sysDictTypeHandler.EditDictTypeInfo)
	app.Delete("/system/dict/type/refreshCache", ms.PermissionMiddleware("system:dict:type:remove"), sysDictTypeHandler.RefreshCache)
	app.Delete("/system/dict/type/*dictIds", ms.PermissionMiddleware("system:dict:type:remove"), sysDictTypeHandler.DeleteDictTypeInfo)

	sysConfigHandler := ryserver.ResolveSysConfigHandler(db, log, freeCache)
	app.Get("/system/config/list", ms.PermissionMiddleware("system:config:list"), sysConfigHandler.ConfigPage)
	app.Get("/system/config/{configId:uint}", ms.PermissionMiddleware("system:config:query"), sysConfigHandler.ConfigInfo)
	app.Post("/system/config", ms.PermissionMiddleware("system:config:add"), sysConfigHandler.AddConfigInfo)
	app.Put("/system/config", ms.PermissionMiddleware("system:config:edit"), sysConfigHandler.EditConfigInfo)
	app.Delete("/system/config/*configIds", ms.PermissionMiddleware("system:config:remove"), sysConfigHandler.DeleteConfigInfo)
	app.Get("/system/config/configKey/{configKey:string}", sysConfigHandler.ConfigInfoByKey)

	sysNoticeHandler := ryserver.ResolveSysNoticeHandler(db, log, freeCache)
	app.Get("/system/notice/list", ms.PermissionMiddleware("system:notice:list"), sysNoticeHandler.NoticePage)
	app.Get("/system/notice/{noticeId:uint}", ms.PermissionMiddleware("system:notice:query"), sysNoticeHandler.NoticeInfo)
	app.Post("/system/notice", ms.PermissionMiddleware("system:notice:add"), sysNoticeHandler.AddNoticeInfo)
	app.Put("/system/notice", ms.PermissionMiddleware("system:notice:edit"), sysNoticeHandler.EditNoticeInfo)
	app.Delete("/system/notice/*noticeIds", ms.PermissionMiddleware("system:notice:remove"), sysNoticeHandler.DeleteNoticeInfo)

	sysLogininforHandler := ryserver.ResolveSysLogininforHandler(db, log, freeCache)
	app.Get("/monitor/logininfor/list", ms.PermissionMiddleware("monitor:logininfor:list"), sysLogininforHandler.LogininforPage)
	app.Get("/monitor/logininfor/{infoId:uint}", ms.PermissionMiddleware("monitor:logininfor:query"), sysLogininforHandler.LogininforInfo)
	app.Post("/monitor/logininfor", ms.PermissionMiddleware("monitor:logininfor:add"), sysLogininforHandler.AddLogininforInfo)
	app.Delete("/monitor/logininfor/*infoIds", ms.PermissionMiddleware("monitor:logininfor:remove"), sysLogininforHandler.DeleteLogininforInfo)

	monitorHandler := ryserver.ResolveMonitorHandler(db, redis, log, freeCache)
	app.Get("/monitor/server", ms.PermissionMiddleware("monitor:server:list"), monitorHandler.Server)
	app.Get("/monitor/cache", ms.PermissionMiddleware("monitor:cache:list"), monitorHandler.Cache)
	app.Get("/monitor/cache/getNames", ms.PermissionMiddleware("monitor:cache:list"), monitorHandler.CacheNames)

	sysJobHandler := ryserver.ResolveSysJobHandler(db, log, freeCache)
	app.Get("/monitor/job/list", ms.PermissionMiddleware("monitor:job:list"), sysJobHandler.JobPage)
	app.Get("/monitor/job/{jobId:uint}", ms.PermissionMiddleware("monitor:job:query"), sysJobHandler.JobInfo)
	app.Post("/monitor/job", ms.PermissionMiddleware("monitor:job:add"), sysJobHandler.AddJobInfo)
	app.Put("/monitor/job", ms.PermissionMiddleware("monitor:job:edit"), sysJobHandler.EditJobInfo)
	app.Delete("/monitor/job/*jobIds", ms.PermissionMiddleware("monitor:job:remove"), sysJobHandler.DeleteJobInfo)

	task := task.NewTaskManager(log)
	task.RegisterTask("ryTask.ryNoParams", jobs.NewTaskDemo(log))

	return &Container{
		appConfig: c,
		logger:    log,
		redis:     redis,
		localizer: l,
		gormDB:    db,
		app:       app,
		freeCache: freeCache,
		jobs:      task,
	}, nil
}

func (c *Container) InitJob() {
	sysJobHandler := ryserver.ResolveSysJobHandler(c.gormDB, c.logger, c.freeCache)
	data, err := sysJobHandler.JobList(nil)
	if err == nil {
		for _, item := range data {
			if item.Status == "0" && item.MisfirePolicy == "1" {
				c.jobs.StartTask(item.InvokeTarget, item.CronExpression)
			}
		}
	}
}

func (c *Container) StartServer() error {
	ryws.StartWebSocket(c.app, c.logger)

	c.logger.Info("http server started", zap.Int("port", c.appConfig.Server.Port))
	err := c.app.Run(iris.Addr(fmt.Sprintf(":%d", c.appConfig.Server.Port)))
	if err != nil {
		c.logger.Error("failed to run http server", zap.Error(err))
		return fmt.Errorf("failed to run http server: %w", err)
	}
	return nil
}

func (c *Container) Close() {
	err := c.gormDB.CloseDB()
	if err != nil {
		c.logger.Error("Failed to close the database connection:", zap.Error(err))
	} else {
		c.logger.Info("database closed")
	}

	// 关闭Redis客户端
	if err := c.redis.CloseRedis(); err != nil {
		c.logger.Error("failed to close redis client", zap.Error(err))
	} else {
		c.logger.Info("Redis client closed")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// close all hosts
	if err := c.app.Shutdown(ctx); err != nil {
		c.logger.Error("failed to close all hosts", zap.Error(err))
	} else {
		c.logger.Info("all hosts closed")
	}

	if c.freeCache != nil {
		c.freeCache.Clear()
	}

	// 关闭日志
	c.logger.Sync()
}
