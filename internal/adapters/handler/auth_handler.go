// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

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

type AuthHandler struct {
	service input.AuthService
	logger  *zap.Logger
}

func NewAuthHandler(service input.AuthService, logger *zap.Logger) *AuthHandler {
	return &AuthHandler{service: service, logger: logger}
}

func (h *AuthHandler) Login(ctx iris.Context) {
	var l model.LoginRequest
	// Attempt to read and bind the JSON request body to the 'user' variable
	if err := filter.ValidateRequest(ctx, &l); err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid JSON, error:%s", err.Error()))
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

func (h *AuthHandler) GetInfo(ctx iris.Context) {
	user := ctx.Values().Get(common.LOGINUSER)
	// 类型断言
	loginUser, ok := user.(*model.UserInfoStruct)
	if !ok {
		ctx.JSON(common.Error(iris.StatusUnauthorized, "请重新登录"))
		return
	}

	info, err := h.service.GetInfo(loginUser)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "getInfo failed, error：%s", err.Error()))
		return
	}
	// 使用 ctx.JSON 自动将user序列化为JSON并写入响应体
	ctx.JSON(info)
}
