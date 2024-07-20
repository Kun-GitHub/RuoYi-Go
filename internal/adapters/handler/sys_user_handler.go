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
	"strconv"
	"strings"
)

type SysUserHandler struct {
	service     input.SysUserService
	deptService input.SysDeptService
	roleService input.SysRoleService
}

func NewSysUserHandler(service input.SysUserService, deptService input.SysDeptService, roleService input.SysRoleService) *SysUserHandler {
	return &SysUserHandler{service: service,
		deptService: deptService, roleService: roleService}
}

// GenerateCaptchaImage
func (this *SysUserHandler) UserPage(ctx iris.Context) {
	// 获取查询参数
	pageNumStr := ctx.URLParamDefault("pageNum", "1")
	pageSizeStr := ctx.URLParamDefault("pageSize", "10")

	pageNum, _ := strconv.Atoi(pageNumStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	l := common.PageRequest{
		pageNum,
		pageSize,
	}

	status := ctx.URLParam("status")
	deptIdStr := ctx.URLParam("deptId")
	var deptId int64
	userName := ctx.URLParam("userName")
	phonenumber := ctx.URLParam("phonenumber")
	//params := ctx.URLParam("params")
	//fmt.Println(params)

	if deptIdStr != "" {
		deptId, _ = strconv.ParseInt(deptIdStr, 10, 64)
	}

	u := &model.SysUser{
		Status:      status,
		DeptID:      deptId,
		UserName:    userName,
		Phonenumber: phonenumber,
	}

	d, t, err := this.service.QueryUserPage(l, u)
	if err != nil {
		//this.logger.Debug("login failed", zap.Error(err))
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

func (this *SysUserHandler) DeptTree(ctx iris.Context) {
	data, err := this.deptService.QueryDeptList(nil)
	if err != nil {
		//this.logger.Debug("login failed", zap.Error(err))
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

func (this *SysUserHandler) UserInfo(ctx iris.Context) {
	userIdStr := ctx.Params().GetString("userId")
	if userIdStr == "" {
		ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid userIdStr"))
		return
	}

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "ParseInt error：%s", err.Error()))
		return
	}

	user, err := this.service.QueryUserInfoByUserId(userId)
	if err != nil {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "QueryUserInfoByUserId, error：%s", err.Error()))
		return
	}

	userInfo := &model.UserInfoStruct{}
	userInfo.SysUser = user

	if user.UserID == common.ADMINID {
		userInfo.Admin = true
	}

	roles, err := this.roleService.QueryRolesByUserId(user.UserID)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "QueryRolesByUserId, error：%s", err.Error()))
		return
	}
	userInfo.Roles = roles

	dept, err := this.deptService.QueryRolesByDeptId(user.DeptID)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "QueryRolesByDeptId, error：%s", err.Error()))
		return
	}
	userInfo.Dept = dept

	ctx.JSON(common.Success(userInfo))
}

func (this *SysUserHandler) ChangeUserStatus(ctx iris.Context) {
	var u model.ChangeUserStatusRequest
	// Attempt to read and bind the JSON request body to the 'user' variable
	if err := filter.ValidateRequest(ctx, &u); err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid JSON, error:%s", err.Error()))
		return
	}

	_, err := this.service.ChangeUserStatus(u)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "ChangeUserStatus error：%s", err.Error()))
		return
	}
	ctx.JSON(common.Success(nil))
}

func (this *SysUserHandler) ResetUserPwd(ctx iris.Context) {
	var u model.ResetUserPwdRequest
	// Attempt to read and bind the JSON request body to the 'user' variable
	if err := filter.ValidateRequest(ctx, &u); err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid JSON, error:%s", err.Error()))
		return
	}

	_, err := this.service.ResetUserPwd(u)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "ResetUserPwd error：%s", err.Error()))
		return
	}
	ctx.JSON(common.Success(nil))
}

func (this *SysUserHandler) DeleteUser(ctx iris.Context) {
	userIdStr := ctx.Params().GetString("userId")
	if userIdStr == "" {
		ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid userIdStr"))
		return
	}

	parts := strings.Split(userIdStr, ",")
	for _, part := range parts {
		userId, err := strconv.ParseInt(part, 10, 64)
		if err != nil {
			ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "ParseInt error：%s", err.Error()))
			return
		}

		_, err = this.service.DeleteUserByUserId(userId)
		if err != nil {
			ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "DeleteUserByUserId error：%s", err.Error()))
			return
		}
	}

	ctx.JSON(common.Success(nil))
}