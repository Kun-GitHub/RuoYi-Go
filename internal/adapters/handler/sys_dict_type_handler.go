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

type SysDictTypeHandler struct {
	service input.SysDictTypeService
}

func NewSysDictTypeHandler(service input.SysDictTypeService) *SysDictTypeHandler {
	return &SysDictTypeHandler{service: service}
}

// GenerateCaptchaImage
func (h *SysDictTypeHandler) DictTypePage(ctx iris.Context) {
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
	dictName := ctx.URLParam("dictName")
	dictType := ctx.URLParam("dictType")
	u := &model.SysDictTypeRequest{
		Status:    status,
		DictName:  dictName,
		DictType:  dictType,
		BeginTime: beginTime,
		EndTime:   endTime,
	}

	datas, total, err := h.service.QueryDictTypePage(l, u)
	if err != nil {
		//h.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "QueryDictTypePage, error：%s", err.Error()))
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

// GenerateCaptchaImage
func (h *SysDictTypeHandler) DictTypeList(ctx iris.Context) {

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

	status := ctx.URLParam("status")
	dictName := ctx.URLParam("dictName")
	dictType := ctx.URLParam("dictType")
	u := &model.SysDictTypeRequest{
		Status:    status,
		DictName:  dictName,
		DictType:  dictType,
		BeginTime: beginTime,
		EndTime:   endTime,
	}

	datas, err := h.service.QueryDictTypeList(u)
	if err != nil {
		//h.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "QueryDictTypePage, error：%s", err.Error()))
		return
	}

	data := &common.Response{
		Data:    datas,
		Message: "操作成功",
		Code:    iris.StatusOK,
	}

	ctx.JSON(data)
}

func (this *SysDictTypeHandler) DictTypeInfo(ctx iris.Context) {
	dictIdStr := ctx.Params().GetString("dictId")
	if dictIdStr == "" {
		ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid dictIdStr"))
		return
	}

	postId, err := strconv.ParseInt(dictIdStr, 10, 64)
	if err != nil {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "ParseInt error：%s", err.Error()))
		return
	}

	info, err := this.service.QueryDictTypeByDictID(postId)
	if err != nil {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "QueryDictTypeByDictID, error：%s", err.Error()))
		return
	}

	ctx.JSON(common.Success(info))
}

func (this *SysDictTypeHandler) AddDictTypeInfo(ctx iris.Context) {
	post := &model.SysDictType{}
	// Attempt to read and bind the JSON request body to the 'user' variable
	if err := filter.ValidateRequest(ctx, post); err != nil {
		//ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid JSON, error:%s", err.Error()))
		return
	}

	count, err := this.service.CheckDictTypeUnique(-1, post.DictType)
	if err != nil || count != 0 {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "新增岗位字典，已存在相同字典类型"))
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

	info, err := this.service.AddDictType(post)
	if err != nil {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "AddDictType, error：%s", err.Error()))
		return
	}

	ctx.JSON(common.Success(info))
}

func (this *SysDictTypeHandler) EditDictTypeInfo(ctx iris.Context) {

	post := &model.SysDictType{}
	// Attempt to read and bind the JSON request body to the 'user' variable
	if err := filter.ValidateRequest(ctx, post); err != nil {
		//ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid JSON, error:%s", err.Error()))
		return
	}

	count, err := this.service.CheckDictTypeUnique(post.DictID, post.DictType)
	if err != nil || count != 0 {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "修改岗位字典，已存在相同字典类型"))
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

	info, _, err := this.service.EditDictType(post)
	if err != nil {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "EditDictType, error：%s", err.Error()))
		return
	}

	ctx.JSON(common.Success(info))
}

func (this *SysDictTypeHandler) DeleteDictTypeInfo(ctx iris.Context) {
	postIdStr := ctx.Params().GetString("dictIds")
	if postIdStr == "" {
		ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid postIdStr"))
		return
	}

	parts := strings.Split(postIdStr, ",")
	for _, part := range parts {
		id, err := strconv.ParseInt(part, 10, 64)
		if err != nil {
			ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "ParseInt error：%s", err.Error()))
			return
		}

		_, err = this.service.DeleteDictTypeById(id)
		if err != nil {
			ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "DeleteDictTypeById error：%s", err.Error()))
			return
		}
	}

	ctx.JSON(common.Success(nil))
}

func (this *SysDictTypeHandler) RefreshCache(ctx iris.Context) {
	ctx.JSON(common.Success(nil))
}
