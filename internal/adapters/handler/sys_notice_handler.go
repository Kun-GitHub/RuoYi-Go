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

type SysNoticeHandler struct {
	service input.SysNoticeService
}

func NewSysNoticeHandler(service input.SysNoticeService) *SysNoticeHandler {
	return &SysNoticeHandler{service: service}
}

// GenerateCaptchaImage
func (h *SysNoticeHandler) NoticePage(ctx iris.Context) {
	// 获取查询参数
	pageNumStr := ctx.URLParamDefault("pageNum", "1")
	pageSizeStr := ctx.URLParamDefault("pageSize", "10")

	pageNum, _ := strconv.Atoi(pageNumStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)
	l := common.PageRequest{
		pageNum,
		pageSize,
	}

	noticeTitle := ctx.URLParam("noticeTitle")
	noticeType := ctx.URLParam("noticeType")
	createBy := ctx.URLParam("createBy")
	u := &model.SysNoticeRequest{
		NoticeTitle: noticeTitle,
		NoticeType:  noticeType,
		CreateBy:    createBy,
	}

	datas, total, err := h.service.QueryNoticePage(l, u)
	if err != nil {
		//h.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "QueryNoticePage, error：%s", err.Error()))
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

func (this *SysNoticeHandler) NoticeInfo(ctx iris.Context) {
	idStr := ctx.Params().GetString("noticeId")
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

	info, err := this.service.QueryNoticeByID(id)
	if err != nil {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "QueryNoticeByID, error：%s", err.Error()))
		return
	}

	ctx.JSON(common.Success(info))
}

func (this *SysNoticeHandler) AddNoticeInfo(ctx iris.Context) {
	post := &model.SysNotice{}
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

	info, err := this.service.AddNotice(post)
	if err != nil {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "AddNotice, error：%s", err.Error()))
		return
	}

	ctx.JSON(common.Success(info))
}

func (this *SysNoticeHandler) EditNoticeInfo(ctx iris.Context) {
	post := &model.SysNotice{}
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

	info, _, err := this.service.EditNotice(post)
	if err != nil {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "EditNotice, error：%s", err.Error()))
		return
	}

	ctx.JSON(common.Success(info))
}

func (this *SysNoticeHandler) DeleteNoticeInfo(ctx iris.Context) {
	idStr := ctx.Params().GetString("noticeId")
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

		_, err = this.service.DeleteNoticeById(id)
		if err != nil {
			ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "DeleteNoticeById error：%s", err.Error()))
			return
		}
	}

	ctx.JSON(common.Success(nil))
}
