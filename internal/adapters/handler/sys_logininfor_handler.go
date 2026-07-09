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

type SysLogininforHandler struct {
	service input.SysLogininforService
}

func NewSysLogininforHandler(service input.SysLogininforService) *SysLogininforHandler {
	return &SysLogininforHandler{service: service}
}

// GenerateCaptchaImage
func (h *SysLogininforHandler) LogininforPage(ctx iris.Context) {
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
	ipaddr := ctx.URLParam("ipaddr")
	userName := ctx.URLParam("userName")
	u := &model.SysLogininforRequest{
		Status:    status,
		Ipaddr:    ipaddr,
		UserName:  userName,
		BeginTime: beginTime,
		EndTime:   endTime,
	}

	datas, total, err := h.service.QueryLogininforPage(l, u)
	if err != nil {
		//h.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "QueryLogininforPage, error：%s", err.Error()))
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

func (this *SysLogininforHandler) LogininforInfo(ctx iris.Context) {
	idStr := ctx.Params().GetString("infoId")
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

	info, err := this.service.QueryLogininforByID(id)
	if err != nil {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "QueryLogininforByID, error：%s", err.Error()))
		return
	}

	ctx.JSON(common.Success(info))
}

func (this *SysLogininforHandler) AddLogininforInfo(ctx iris.Context) {
	post := &model.SysLogininfor{}
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
	post.LoginTime = time.Now()
	post.UserName = loginUser.UserName

	info, err := this.service.AddLogininfor(post)
	if err != nil {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "AddLogininfor, error：%s", err.Error()))
		return
	}

	ctx.JSON(common.Success(info))
}

func (this *SysLogininforHandler) DeleteLogininforInfo(ctx iris.Context) {
	idStr := ctx.Params().GetString("infoIds")
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

		_, err = this.service.DeleteLogininforById(id)
		if err != nil {
			ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "DeleteLogininforById error：%s", err.Error()))
			return
		}
	}

	ctx.JSON(common.Success(nil))
}

func (this *SysLogininforHandler) Export(ctx iris.Context) {
	allParams := ctx.Request().URL.Query()
	beginTimeList, _ := allParams["params[beginTime]"]
	endTimeList, _ := allParams["params[endTime]"]
	beginTime := ""
	if len(beginTimeList) > 0 {
		beginTime = beginTimeList[0]
	}
	endTime := ""
	if len(endTimeList) > 0 {
		endTime = endTimeList[0]
	}

	status := ctx.URLParam("status")
	ipaddr := ctx.URLParam("ipaddr")
	userName := ctx.URLParam("userName")
	u := &model.SysLogininforRequest{
		Status:    status,
		Ipaddr:    ipaddr,
		UserName:  userName,
		BeginTime: beginTime,
		EndTime:   endTime,
	}

	list, err := this.service.QueryLogininforList(u)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "Export error: %s", err.Error()))
		return
	}

	headers := []string{"访问ID", "用户账号", "登录IP地址", "登录地点", "浏览器", "操作系统", "登录状态", "提示消息", "访问时间"}
	rows := make([][]interface{}, len(list))
	for i, item := range list {
		loginTime := ""
		if !item.LoginTime.IsZero() {
			loginTime = item.LoginTime.Format("2006-01-02 15:04:05")
		}
		rows[i] = []interface{}{
			item.InfoID,
			item.UserName,
			item.Ipaddr,
			item.LoginLocation,
			item.Browser,
			item.Os,
			item.Status,
			item.Msg,
			loginTime,
		}
	}

	filePath, err := excel.ExportExcel(headers, rows, "登录日志")
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "ExportExcel error: %s", err.Error()))
		return
	}
	defer os.Remove(filePath)

	ctx.SendFile(filePath, "logininfor.xlsx")
}

func (this *SysLogininforHandler) Unlock(ctx iris.Context) {
	_ = ctx.Params().GetString("userName")
	// In Java, this clears login lock for the user
	// For now, just return success
	ctx.JSON(common.Success(nil))
}
