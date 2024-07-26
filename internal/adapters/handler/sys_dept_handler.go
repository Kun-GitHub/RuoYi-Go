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
	"github.com/kataras/iris/v12"
	"strconv"
	"strings"
	"time"
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

// GenerateCaptchaImage
func (h *SysDeptHandler) DeptListExcludeById(ctx iris.Context) {
	idStr := ctx.Params().GetString("deptId")
	if idStr == "" {
		ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid idStr"))
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "ParseInt error：%s", err.Error()))
		return
	}

	s, err := h.service.QueryDeptListExcludeById(id)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "DeptList failed, error：%s", err.Error()))
		return
	}

	ctx.JSON(common.Success(s))
}

func (this *SysDeptHandler) DeptInfo(ctx iris.Context) {
	idStr := ctx.Params().GetString("deptId")
	if idStr == "" {
		ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid idStr"))
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "ParseInt error：%s", err.Error()))
		return
	}

	info, err := this.service.QueryDeptById(id)
	if err != nil {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "QueryDeptByID, error：%s", err.Error()))
		return
	}

	ctx.JSON(common.Success(info))
}

func (this *SysDeptHandler) AddDeptInfo(ctx iris.Context) {
	post := &model.SysDept{}
	// Attempt to read and bind the JSON request body to the 'user' variable
	if err := filter.ValidateRequest(ctx, post); err != nil {
		//ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid JSON, error:%s", err.Error()))
		return
	}

	user := ctx.Values().Get(common.LOGINUSER)
	// 类型断言
	loginUser, ok := user.(*model.UserInfoStruct)
	if !ok {
		ctx.JSON(common.Error(iris.StatusUnauthorized, "请重新登录"))
		return
	}
	post.CreateTime = time.Now()
	post.CreateBy = loginUser.UserName
	post.UpdateTime = time.Now()
	post.UpdateBy = loginUser.UserName

	info, err := this.service.AddDept(post)
	if err != nil {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "AddDept, error：%s", err.Error()))
		return
	}

	ctx.JSON(common.Success(info))
}

func (this *SysDeptHandler) EditDeptInfo(ctx iris.Context) {
	post := &model.SysDept{}
	// Attempt to read and bind the JSON request body to the 'user' variable
	if err := filter.ValidateRequest(ctx, post); err != nil {
		//ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid JSON, error:%s", err.Error()))
		return
	}

	user := ctx.Values().Get(common.LOGINUSER)
	// 类型断言
	loginUser, ok := user.(*model.UserInfoStruct)
	if !ok {
		ctx.JSON(common.Error(iris.StatusUnauthorized, "请重新登录"))
		return
	}
	post.UpdateTime = time.Now()
	post.UpdateBy = loginUser.UserName

	info, _, err := this.service.EditDept(post)
	if err != nil {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "EditDept, error：%s", err.Error()))
		return
	}

	ctx.JSON(common.Success(info))
}

func (this *SysDeptHandler) DeleteDeptInfo(ctx iris.Context) {
	idStr := ctx.Params().GetString("deptIds")
	if idStr == "" {
		ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid idStr"))
		return
	}

	parts := strings.Split(idStr, ",")
	for _, part := range parts {
		id, err := strconv.ParseInt(part, 10, 64)
		if err != nil {
			ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "ParseInt error：%s", err.Error()))
			return
		}

		_, err = this.service.DeleteDeptById(id)
		if err != nil {
			ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "DeleteDeptById error：%s", err.Error()))
			return
		}
	}

	ctx.JSON(common.Success(nil))
}
