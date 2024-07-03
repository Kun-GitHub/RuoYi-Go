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

type SysUserHandler struct {
	service     input.SysUserService
	deptService input.SysDeptService
}

func NewSysUserHandler(service input.SysUserService, deptService input.SysDeptService) *SysUserHandler {
	return &SysUserHandler{service: service,
		deptService: deptService}
}

// GenerateCaptchaImage
func (h *SysUserHandler) UserPage(ctx iris.Context) {
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

func (h *SysUserHandler) DeptTree(ctx iris.Context) {
	data, err := h.deptService.QueryDeptList(nil)
	if err != nil {
		//h.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "DeptTree, error：%s", err.Error()))
		return
	}

	data = buildDeptTree(data)
	ctx.JSON(common.Success(data))
}

// buildDeptTree constructs a tree of departments from a flat list.
func buildDeptTree(depts []*model.SysDept) []*model.SysDept {
	// Create a map to store the department by its ID.
	idToDept := make(map[int64]*model.SysDept)
	for _, dept := range depts {
		idToDept[dept.DeptID] = dept
	}

	// Create a slice to hold the root departments.
	var rootDepts []*model.SysDept

	// Iterate over the departments and set up the parent-child relationships.
	for _, dept := range depts {
		if parentId, exists := idToDept[dept.ParentID]; exists {
			// If the parent department exists, add the current department as its child.
			dept.ID = dept.DeptID
			dept.Label = dept.DeptName

			parentId.Children = append(parentId.Children, dept)
		} else {
			dept.ID = dept.DeptID
			dept.Label = dept.DeptName
			// If the parent department does not exist, it's a root department.
			rootDepts = append(rootDepts, dept)
		}
	}

	// Return the root departments.
	return rootDepts
}
