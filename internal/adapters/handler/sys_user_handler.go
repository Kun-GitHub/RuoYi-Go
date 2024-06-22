// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package handler

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/ports/input"
	"github.com/kataras/iris/v12"
	"strconv"
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

	pageNumStr := ctx.URLParam("pageNum")
	pageSizeStr := ctx.URLParam("pageSize")

	pageNum, _ := strconv.Atoi(pageNumStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	l := common.PageRequest{
		pageNum,
		pageSize,
	}
	data, err := h.service.QueryUserPage(l, 0, "", "", "", 0)
	if err != nil {
		//h.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "UserPage, error：%s", err.Error()))
		return
	}

	data = &common.PageResponse{
		Rows:    data.Rows,
		Total:   data.Total,
		Message: "操作成功",
		Code:    iris.StatusOK,
	}

	ctx.JSON(data)
}
