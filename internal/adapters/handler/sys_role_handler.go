// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package handler

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/domain/model"
	"RuoYi-Go/internal/ports/input"
	"github.com/kataras/iris/v12"
	"strconv"
)

type SysRoleHandler struct {
	service input.SysRoleService
}

func NewSysRoleHandler(service input.SysRoleService) *SysRoleHandler {
	return &SysRoleHandler{service: service}
}

// GenerateCaptchaImage
func (h *SysRoleHandler) RolePage(ctx iris.Context) {
	pageNumStr := ctx.URLParam("pageNum")
	pageSizeStr := ctx.URLParam("pageSize")

	pageNum, _ := strconv.Atoi(pageNumStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	l := common.PageRequest{
		pageNum,
		pageSize,
	}

	status := ctx.URLParam("status")
	roleName := ctx.URLParam("roleName")
	roleKey := ctx.URLParam("roleKey")
	u := &model.SysRole{
		Status:   status,
		RoleName: roleName,
		RoleKey:  roleKey,
	}

	d, t, err := h.service.QueryRolePage(l, u)
	if err != nil {
		//h.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "UserPage, error：%s", err.Error()))
		return
	}

	data := &common.PageResponse{
		Rows:    d,
		Total:   t,
		Message: "操作成功",
		Code:    iris.StatusOK,
	}

	ctx.JSON(data)
}
