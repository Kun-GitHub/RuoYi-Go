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
	"strconv"
	"time"

	"github.com/kataras/iris/v12"
)

type SysDictDataHandler struct {
	service input.SysDictDataService
}

func NewSysDictDataHandler(service input.SysDictDataService) *SysDictDataHandler {
	return &SysDictDataHandler{service: service}
}

// DictType 查询字典数据
func (h *SysDictDataHandler) DictType(ctx iris.Context) {
	dictType := ctx.Params().GetString("dictType")
	if dictType == "" {
		ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid dictType"))
		return
	}

	s, err := h.service.QueryDictDatasByType(dictType)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "DictType failed, error：%s", err.Error()))
		return
	}
	ctx.JSON(common.Success(s))
}

func (h *SysDictDataHandler) List(ctx iris.Context) {
	pageNumStr := ctx.URLParamDefault("pageNum", "1")
	pageSizeStr := ctx.URLParamDefault("pageSize", "10")
	pageNum, _ := strconv.Atoi(pageNumStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	dictLabel := ctx.URLParam("dictLabel")
	dictType := ctx.URLParam("dictType")
	status := ctx.URLParam("status")

	list, total, err := h.service.List(pageNum, pageSize, dictLabel, dictType, status)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "List failed, error：%s", err.Error()))
		return
	}

	ctx.JSON(&common.PageResponse{
		Rows:    list,
		Total:   total,
		Message: "操作成功",
		Code:    iris.StatusOK,
	})
}

func (h *SysDictDataHandler) Get(ctx iris.Context) {
	dictCodeStr := ctx.Params().GetString("dictCode")
	dictCode, err := strconv.ParseUint(dictCodeStr, 10, 64)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid dictCode"))
		return
	}

	data, err := h.service.Get(uint(dictCode))
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "Get failed, error：%s", err.Error()))
		return
	}
	ctx.JSON(common.Success(data))
}

func (h *SysDictDataHandler) Add(ctx iris.Context) {
	var data model.SysDictDatum
	if err := filter.ValidateRequest(ctx, &data); err != nil {
		// ValidateRequest already handles error response if needed, possibly?
		// Actually looking at SysUserHandler, it seems to handle error return itself if ValidateRequest returns error?
		// No, SysUserHandler just returns if ValidateRequest returns error.
		return
	}

	user := ctx.Values().Get(common.LOGINUSER)
	if loginUser, ok := user.(*model.UserInfoStruct); ok {
		data.CreateBy = loginUser.UserName
		data.UpdateBy = loginUser.UserName
	}
	data.CreateTime = time.Now()
	data.UpdateTime = time.Now()

	err := h.service.Create(&data)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "Add failed, error：%s", err.Error()))
		return
	}
	ctx.JSON(common.Success(nil))
}

func (h *SysDictDataHandler) Edit(ctx iris.Context) {
	var data model.SysDictDatum
	if err := filter.ValidateRequest(ctx, &data); err != nil {
		return
	}

	user := ctx.Values().Get(common.LOGINUSER)
	if loginUser, ok := user.(*model.UserInfoStruct); ok {
		data.UpdateBy = loginUser.UserName
	}
	data.UpdateTime = time.Now()

	err := h.service.Update(&data)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "Edit failed, error：%s", err.Error()))
		return
	}
	ctx.JSON(common.Success(nil))
}

func (h *SysDictDataHandler) Delete(ctx iris.Context) {
	dictCodes := ctx.Params().GetString("dictCodes")
	if dictCodes == "" {
		ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid dictCodes"))
		return
	}

	err := h.service.Delete(dictCodes)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "Delete failed, error：%s", err.Error()))
		return
	}
	ctx.JSON(common.Success(nil))
}
