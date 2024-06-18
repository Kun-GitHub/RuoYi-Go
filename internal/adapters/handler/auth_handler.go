// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package handler

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/domain/model"
	"RuoYi-Go/internal/middlewares"
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
	if err := ctx.ReadJSON(&l); err != nil {
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

	loginUser := middlewares.GetLoginUser()
	if loginUser != nil {
		middlewares.ClearLoginUser()
	}

	ctx.Values().Set(common.TOKEN, nil)
	ctx.Values().Set(common.USER_ID, nil)

	ctx.JSON(common.Success("Logout successful"))
}

func (h *AuthHandler) GetInfo(ctx iris.Context) {
	loginUser := middlewares.GetLoginUser()
	if loginUser == nil || loginUser.UserID == 0 {
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
