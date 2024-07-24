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

type SysConfigHandler struct {
	service input.SysConfigService
}

func NewSysConfigHandler(service input.SysConfigService) *SysConfigHandler {
	return &SysConfigHandler{service: service}
}

// GenerateCaptchaImage
func (h *SysConfigHandler) ConfigPage(ctx iris.Context) {
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

	configType := ctx.URLParam("configType")
	configName := ctx.URLParam("configName")
	configKey := ctx.URLParam("configKey")
	u := &model.SysConfigRequest{
		ConfigType: configType,
		ConfigName: configName,
		ConfigKey:  configKey,
		BeginTime:  beginTime,
		EndTime:    endTime,
	}

	datas, total, err := h.service.QueryConfigPage(l, u)
	if err != nil {
		//h.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "QueryConfigPage, error：%s", err.Error()))
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

func (this *SysConfigHandler) ConfigInfo(ctx iris.Context) {
	idStr := ctx.Params().GetString("configId")
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

	info, err := this.service.QueryConfigByID(id)
	if err != nil {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "QueryConfigByID, error：%s", err.Error()))
		return
	}

	ctx.JSON(common.Success(info))
}

func (this *SysConfigHandler) AddConfigInfo(ctx iris.Context) {
	post := &model.SysConfig{}
	// Attempt to read and bind the JSON request body to the 'user' variable
	if err := filter.ValidateRequest(ctx, post); err != nil {
		//ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid JSON, error:%s", err.Error()))
		return
	}

	count, err := this.service.CheckConfigNameUnique(-1, post.ConfigName)
	if err != nil || count != 0 {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "新增参数失败，已存在相同参数键名"))
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

	info, err := this.service.AddConfig(post)
	if err != nil {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "AddConfig, error：%s", err.Error()))
		return
	}

	ctx.JSON(common.Success(info))
}

func (this *SysConfigHandler) EditConfigInfo(ctx iris.Context) {

	post := &model.SysConfig{}
	// Attempt to read and bind the JSON request body to the 'user' variable
	if err := filter.ValidateRequest(ctx, post); err != nil {
		//ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid JSON, error:%s", err.Error()))
		return
	}

	count, err := this.service.CheckConfigNameUnique(post.ConfigID, post.ConfigName)
	if err != nil || count != 0 {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "修改参数失败，已存在相同参数键名"))
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

	info, _, err := this.service.EditConfig(post)
	if err != nil {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "EditConfig, error：%s", err.Error()))
		return
	}

	ctx.JSON(common.Success(info))
}

func (this *SysConfigHandler) DeleteConfigInfo(ctx iris.Context) {
	idStr := ctx.Params().GetString("configId")
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

		_, err = this.service.DeleteConfigById(id)
		if err != nil {
			ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "DeleteConfigById error：%s", err.Error()))
			return
		}
	}

	ctx.JSON(common.Success(nil))
}
