// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package handler

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/domain/model"
	"RuoYi-Go/internal/filter"
	"RuoYi-Go/internal/ports/input"
	"fmt"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
)

// AuthHandler 认证处理器
// 负责处理HTTP认证相关的请求，包括登录、登出、获取用户信息等
type AuthHandler struct {
	service input.AuthService
	logger  *zap.Logger
}

// NewAuthHandler 创建认证处理器实例
//
// 参数:
//
//	service: 认证服务接口
//	logger: 日志记录器
//
// 返回值:
//
//	认证处理器
func NewAuthHandler(service input.AuthService, logger *zap.Logger) *AuthHandler {
	return &AuthHandler{service: service, logger: logger}
}

// Login 处理用户登录请求
// 接收登录表单数据，调用认证服务进行登录验证
//
// 参数:
//
//	ctx: Iris上下文对象
func (h *AuthHandler) Login(ctx iris.Context) {
	l := &model.LoginRequest{}
	// Attempt to read and bind the JSON request body to the 'user' variable
	if err := filter.ValidateRequest(ctx, l); err != nil {
		//ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid JSON, error:%s", err.Error()))
		return
	}

	resp, err := h.service.Login(l)
	if err != nil {
		//h.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "login failed, error：%s", err.Error()))
		return
	}

	ctx.JSON(resp)
}

// Logout 处理用户登出请求
// 清除用户的会话信息，实现安全退出
//
// 参数:
//
//	ctx: Iris上下文对象
func (h *AuthHandler) Logout(ctx iris.Context) {
	token := ctx.Values().Get(common.TOKEN)
	if token != nil {
		if err := h.service.Logout(fmt.Sprintf("%s:%s", common.TOKEN, token)); err != nil {
			ctx.JSON(common.Error(iris.StatusInternalServerError, "Logout failed"))
			return
		}
	}

	loginUser := ctx.Values().Get(common.LOGINUSER)
	// 类型断言
	_, ok := loginUser.(*model.UserInfoStruct)
	if ok {
		ctx.Values().Remove(common.LOGINUSER)
	}

	ctx.Values().Remove(common.TOKEN)
	ctx.Values().Remove(common.USER_ID)

	ctx.JSON(common.Success("Logout successful"))
}

// GetInfo 处理获取用户信息请求
// 返回当前登录用户的详细信息、权限和角色
//
// 参数:
//
//	ctx: Iris上下文对象
func (h *AuthHandler) GetInfo(ctx iris.Context) {
	user := ctx.Values().Get(common.LOGINUSER)
	// 类型断言
	loginUser, ok := user.(*model.UserInfoStruct)
	if !ok {
		ctx.JSON(common.Error(iris.StatusUnauthorized, "请重新登录"))
		return
	}

	info, p, roleNames, err := h.service.GetInfo(loginUser)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "getInfo failed, error：%s", err.Error()))
		return
	}

	infoSuccess := &model.GetInfoSuccess{
		Code:        common.SUCCESS,
		User:        info,
		Permissions: p,
		Roles:       roleNames,
		Message:     "操作成功",
	}

	// 使用 ctx.JSON 自动将user序列化为JSON并写入响应体
	ctx.JSON(infoSuccess)
}
