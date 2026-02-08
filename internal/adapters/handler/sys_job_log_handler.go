package handler

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/ports/input"
	"strconv"

	"github.com/kataras/iris/v12"
)

type SysJobLogHandler struct {
	service input.SysJobLogService
}

func NewSysJobLogHandler(service input.SysJobLogService) *SysJobLogHandler {
	return &SysJobLogHandler{service: service}
}

func (h *SysJobLogHandler) List(ctx iris.Context) {
	pageNumStr := ctx.URLParamDefault("pageNum", "1")
	pageSizeStr := ctx.URLParamDefault("pageSize", "10")
	pageNum, _ := strconv.Atoi(pageNumStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	jobName := ctx.URLParam("jobName")
	jobGroup := ctx.URLParam("jobGroup")
	status := ctx.URLParam("status")

	allParams := ctx.Request().URL.Query()
	beginTimeList, _ := allParams["params[beginTime]"]
	endTimeList, _ := allParams["params[endTime]"]
	var createTime []string
	if len(beginTimeList) > 0 && len(endTimeList) > 0 {
		createTime = []string{beginTimeList[0], endTimeList[0]}
	}

	list, total, err := h.service.List(pageNum, pageSize, jobName, jobGroup, status, createTime)
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

func (h *SysJobLogHandler) Delete(ctx iris.Context) {
	jobLogIds := ctx.Params().GetString("jobLogIds")
	if jobLogIds == "" {
		ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid jobLogIds"))
		return
	}

	err := h.service.Delete(jobLogIds)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "Delete failed, error：%s", err.Error()))
		return
	}
	ctx.JSON(common.Success(nil))
}

func (h *SysJobLogHandler) Clean(ctx iris.Context) {
	err := h.service.Clean()
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "Clean failed, error：%s", err.Error()))
		return
	}
	ctx.JSON(common.Success(nil))
}
