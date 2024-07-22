// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package handler

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/domain/model"
	"RuoYi-Go/internal/ports/input"
	"github.com/kataras/iris/v12"
	"strconv"
)

type SysPostHandler struct {
	service input.SysPostService
}

func NewSysPostHandler(service input.SysPostService) *SysPostHandler {
	return &SysPostHandler{service: service}
}

// GenerateCaptchaImage
func (h *SysPostHandler) PostPage(ctx iris.Context) {
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
	postCode := ctx.URLParam("postCode")
	postName := ctx.URLParam("postName")
	u := &model.SysPostRequest{
		Status:   status,
		PostCode: postCode,
		PostName: postName,
	}

	datas, total, err := h.service.QueryPostPage(l, u)
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

func (this *SysPostHandler) PostInfo(ctx iris.Context) {
	postIdStr := ctx.Params().GetString("postId")
	if postIdStr == "" {
		ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid postIdStr"))
		return
	}

	postId, err := strconv.ParseInt(postIdStr, 10, 64)
	if err != nil {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "ParseInt error：%s", err.Error()))
		return
	}

	info, err := this.service.QueryPostByPostId(postId)
	if err != nil {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "QueryUserInfoByUserId, error：%s", err.Error()))
		return
	}

	ctx.JSON(common.Success(info))
}
