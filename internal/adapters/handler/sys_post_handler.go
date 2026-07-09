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
	"RuoYi-Go/pkg/excel"
	"github.com/kataras/iris/v12"
	"os"
	"strconv"
	"strings"
	"time"
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
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "QueryPostPage, error：%s", err.Error()))
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
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "QueryPostByPostId, error：%s", err.Error()))
		return
	}

	ctx.JSON(common.Success(info))
}

func (this *SysPostHandler) AddPostInfo(ctx iris.Context) {
	post := &model.SysPost{}
	// Attempt to read and bind the JSON request body to the 'user' variable
	if err := filter.ValidateRequest(ctx, post); err != nil {
		//ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid JSON, error:%s", err.Error()))
		return
	}

	count, err := this.service.CheckPostNameUnique(-1, post.PostName)
	if err != nil || count != 0 {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "新增岗位失败，已存在相同名称岗位"))
		return
	}
	count, err = this.service.CheckPostCodeUnique(-1, post.PostCode)
	if err != nil || count != 0 {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "新增岗位失败，已存在相同编号岗位"))
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

	info, err := this.service.AddPost(post)
	if err != nil {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "AddPost, error：%s", err.Error()))
		return
	}

	ctx.JSON(common.Success(info))
}

func (this *SysPostHandler) EditPostInfo(ctx iris.Context) {

	post := &model.SysPost{}
	// Attempt to read and bind the JSON request body to the 'user' variable
	if err := filter.ValidateRequest(ctx, post); err != nil {
		//ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid JSON, error:%s", err.Error()))
		return
	}

	count, err := this.service.CheckPostNameUnique(post.PostID, post.PostName)
	if err != nil || count != 0 {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "修改岗位失败，已存在相同名称岗位"))
		return
	}
	count, err = this.service.CheckPostCodeUnique(post.PostID, post.PostCode)
	if err != nil || count != 0 {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "修改岗位失败，已存在相同编号岗位"))
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

	info, _, err := this.service.EditPost(post)
	if err != nil {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "EditPost, error：%s", err.Error()))
		return
	}

	ctx.JSON(common.Success(info))
}

func (this *SysPostHandler) DeletePostInfo(ctx iris.Context) {
	postIdStr := ctx.Params().GetString("postIds")
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

		_, err = this.service.DeletePostById(id)
		if err != nil {
			ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "DeletePostById error：%s", err.Error()))
			return
		}
	}

	ctx.JSON(common.Success(nil))
}

func (this *SysPostHandler) Export(ctx iris.Context) {
	status := ctx.URLParam("status")
	postCode := ctx.URLParam("postCode")
	postName := ctx.URLParam("postName")
	u := &model.SysPostRequest{
		Status:   status,
		PostCode: postCode,
		PostName: postName,
	}

	list, err := this.service.QueryPostList(u)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "Export error: %s", err.Error()))
		return
	}

	headers := []string{"岗位ID", "岗位编码", "岗位名称", "显示顺序", "状态", "创建时间"}
	rows := make([][]interface{}, len(list))
	for i, item := range list {
		createTime := ""
		if !item.CreateTime.IsZero() {
			createTime = item.CreateTime.Format("2006-01-02 15:04:05")
		}
		rows[i] = []interface{}{
			item.PostID,
			item.PostCode,
			item.PostName,
			item.PostSort,
			item.Status,
			createTime,
		}
	}

	filePath, err := excel.ExportExcel(headers, rows, "岗位数据")
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "ExportExcel error: %s", err.Error()))
		return
	}
	defer os.Remove(filePath)

	ctx.SendFile(filePath, "post.xlsx")
}

func (this *SysPostHandler) OptionSelect(ctx iris.Context) {
	data, err := this.service.SelectPostAll()
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "OptionSelect error: %s", err.Error()))
		return
	}
	ctx.JSON(common.Success(data))
}
