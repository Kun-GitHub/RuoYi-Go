package handler

import (
	"RuoYi-Go/internal/common"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
)

type MonitorHandler struct {
	logger *zap.Logger
}

func NewMonitorHandler(logger *zap.Logger) *MonitorHandler {
	return &MonitorHandler{logger: logger}
}

func (h *MonitorHandler) Server(ctx iris.Context) {
	ctx.JSON(common.Success(nil))
}

func (h *MonitorHandler) Cache(ctx iris.Context) {
	ctx.JSON(common.Success(nil))
}

func (h *MonitorHandler) CacheNames(ctx iris.Context) {
	ctx.JSON(common.Success(nil))
}
