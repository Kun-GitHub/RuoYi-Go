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
	"github.com/kataras/iris/v12"
)

type SysUserHandler struct {
	service input.SysUserService
}

func NewSysUserHandler(service input.SysUserService) *SysUserHandler {
	return &SysUserHandler{service: service}
}

// GenerateCaptchaImage
func (h *SysUserHandler) UserPage(ctx iris.Context) {
	//user := ctx.Values().Get(common.LOGINUSER)
	//// 类型断言
	//loginUser, ok := user.(*model.LoginUserStruct)
	//if !ok {
	//	ctx.JSON(common.Error(iris.StatusUnauthorized, "请重新登录"))
	//	return
	//}

	var l model.UserList
	// Attempt to read and bind the JSON request body to the 'user' variable
	if err := filter.ValidateRequest(ctx, &l); err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid JSON, error:%s", err.Error()))
		return
	}

	//h.service.QueryUserInfoByUserId()

	ctx.JSON(common.Success(""))
}
