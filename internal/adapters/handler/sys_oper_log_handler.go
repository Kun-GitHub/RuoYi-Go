package handler

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/ports/input"
	"RuoYi-Go/pkg/excel"
	"os"
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

func (h *SysOperLogHandler) GetInfo(ctx iris.Context) {
	operIdStr := ctx.Params().GetString("operId")
	if operIdStr == "" {
		ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid operId"))
		return
	}
	operId, err := strconv.ParseInt(operIdStr, 10, 64)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "ParseInt error: %s", err.Error()))
		return
	}
	data, err := h.service.Get(operId)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "GetInfo error: %s", err.Error()))
		return
	}
	ctx.JSON(common.Success(data))
}

func (h *SysOperLogHandler) Export(ctx iris.Context) {
	operIp := ctx.URLParam("operIp")
	title := ctx.URLParam("title")
	operName := ctx.URLParam("operName")
	businessType := ctx.URLParam("businessType")
	status := ctx.URLParam("status")

	allParams := ctx.Request().URL.Query()
	beginTimeList, _ := allParams["params[beginTime]"]
	endTimeList, _ := allParams["params[endTime]"]
	var operTime []string
	if len(beginTimeList) > 0 && len(endTimeList) > 0 {
		operTime = []string{beginTimeList[0], endTimeList[0]}
	}

	list, _, err := h.service.List(1, 999999, operIp, title, operName, businessType, status, operTime)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "Export error: %s", err.Error()))
		return
	}

	headers := []string{"日志ID", "模块标题", "业务类型", "方法名称", "请求方式", "操作人员", "主机地址", "操作状态", "操作时间"}
	rows := make([][]interface{}, len(list))
	for i, item := range list {
		operTimeStr := ""
		if !item.OperTime.IsZero() {
			operTimeStr = item.OperTime.Format("2006-01-02 15:04:05")
		}
		rows[i] = []interface{}{
			item.OperID,
			item.Title,
			item.BusinessType,
			item.Method,
			item.RequestMethod,
			item.OperName,
			item.OperIP,
			item.Status,
			operTimeStr,
		}
	}

	filePath, err := excel.ExportExcel(headers, rows, "操作日志")
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "ExportExcel error: %s", err.Error()))
		return
	}
	defer os.Remove(filePath)

	ctx.SendFile(filePath, "operLog.xlsx")
}

func (h *SysOperLogHandler) Clean(ctx iris.Context) {
	err := h.service.Clean()
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "Clean failed, error：%s", err.Error()))
		return
	}
	ctx.JSON(common.Success(nil))
}
