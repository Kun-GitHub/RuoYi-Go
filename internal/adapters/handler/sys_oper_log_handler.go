package handler

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/ports/input"
	"strconv"

	"github.com/kataras/iris/v12"
)

type SysOperLogHandler struct {
	service input.SysOperLogService
}

func NewSysOperLogHandler(service input.SysOperLogService) *SysOperLogHandler {
	return &SysOperLogHandler{service: service}
}

func (h *SysOperLogHandler) List(ctx iris.Context) {
	pageNumStr := ctx.URLParamDefault("pageNum", "1")
	pageSizeStr := ctx.URLParamDefault("pageSize", "10")
	pageNum, _ := strconv.Atoi(pageNumStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	operIp := ctx.URLParam("operIp")
	title := ctx.URLParam("title")
	operName := ctx.URLParam("operName")
	businessType := ctx.URLParam("businessType")
	status := ctx.URLParam("status")

	// Handle operTime usually passed as params[beginTime] and params[endTime]
	// But let's check SysUserHandler again. It parses params[beginTime].
	allParams := ctx.Request().URL.Query()
	beginTimeList, _ := allParams["params[beginTime]"]
	endTimeList, _ := allParams["params[endTime]"]
	var operTime []string
	if len(beginTimeList) > 0 && len(endTimeList) > 0 {
		operTime = []string{beginTimeList[0], endTimeList[0]}
	}

	list, total, err := h.service.List(pageNum, pageSize, operIp, title, operName, businessType, status, operTime)
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

func (h *SysOperLogHandler) Delete(ctx iris.Context) {
	operIds := ctx.Params().GetString("operIds")
	if operIds == "" {
		ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid operIds"))
		return
	}

	err := h.service.Delete(operIds)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "Delete failed, error：%s", err.Error()))
		return
	}
	ctx.JSON(common.Success(nil))
}

func (h *SysOperLogHandler) Clean(ctx iris.Context) {
	err := h.service.Clean()
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "Clean failed, error：%s", err.Error()))
		return
	}
	ctx.JSON(common.Success(nil))
}
