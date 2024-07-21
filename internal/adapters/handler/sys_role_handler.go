// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

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
	// 获取查询参数
	pageNumStr := ctx.URLParamDefault("pageNum", "1")
	pageSizeStr := ctx.URLParamDefault("pageSize", "10")

	// 使用 Query() 方法获取所有的查询参数
	allParams := ctx.Request().URL.Query()
	// 从 url.Values 结构体中获取参数
	beginTimeList, _ := allParams["params[beginTime]"]
	endTimeList, _ := allParams["params[endTime]"]
	// 假设我们只关心第一个值，我们可以这样获取：
	beginTime := ""
	if len(beginTimeList) > 0 {
		beginTime = beginTimeList[0]
	}
	endTime := ""
	if len(endTimeList) > 0 {
		endTime = endTimeList[0]
	}

	pageNum, _ := strconv.Atoi(pageNumStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)
	l := common.PageRequest{
		pageNum,
		pageSize,
	}

	status := ctx.URLParam("status")
	roleName := ctx.URLParam("roleName")
	roleKey := ctx.URLParam("roleKey")
	u := &model.SysRoleRequest{
		Status:    status,
		RoleName:  roleName,
		RoleKey:   roleKey,
		BeginTime: beginTime,
		EndTime:   endTime,
	}

	datas, total, err := h.service.QueryRolePage(l, u)
	if err != nil {
		//h.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "UserPage, error：%s", err.Error()))
		return
	}

	data := &common.PageResponse{
		Rows:    datas,
		Total:   total,
		Message: "操作成功",
		Code:    iris.StatusOK,
	}

	ctx.JSON(data)
}
