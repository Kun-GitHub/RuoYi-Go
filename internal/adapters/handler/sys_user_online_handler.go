package handler

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/ports/input"
	"strconv"

	"github.com/kataras/iris/v12"
)

type SysUserOnlineHandler struct {
	service input.SysUserOnlineService
}

func NewSysUserOnlineHandler(service input.SysUserOnlineService) *SysUserOnlineHandler {
	return &SysUserOnlineHandler{service: service}
}

func (h *SysUserOnlineHandler) List(ctx iris.Context) {
	pageNumStr := ctx.URLParamDefault("pageNum", "1")
	pageSizeStr := ctx.URLParamDefault("pageSize", "10")
	pageNum, _ := strconv.Atoi(pageNumStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	ipaddr := ctx.URLParam("ipaddr")
	userName := ctx.URLParam("userName")

	list, total, err := h.service.List(pageNum, pageSize, ipaddr, userName)
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

func (h *SysUserOnlineHandler) ForceLogout(ctx iris.Context) {
	tokenId := ctx.Params().GetString("tokenId")
	if tokenId == "" {
		ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid tokenId"))
		return
	}

	err := h.service.ForceLogout(tokenId)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "ForceLogout failed, error：%s", err.Error()))
		return
	}
	ctx.JSON(common.Success(nil))
}
