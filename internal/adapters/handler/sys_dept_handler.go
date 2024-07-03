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
)

type SysDeptHandler struct {
	service input.SysDeptService
}

func NewSysDeptHandler(service input.SysDeptService) *SysDeptHandler {
	return &SysDeptHandler{service: service}
}

// GenerateCaptchaImage
func (h *SysDeptHandler) DeptList(ctx iris.Context) {
	deptName := ctx.URLParam("deptName")
	status := ctx.URLParam("status")

	sysDept := &model.SysDept{
		DeptName: deptName,
		Status:   status,
	}

	s, err := h.service.QueryDeptList(sysDept)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "DeptList failed, error：%s", err.Error()))
		return
	}
	ctx.JSON(common.Success(s))
}
