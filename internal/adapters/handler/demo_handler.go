package handler

import (
	"RuoYi-Go/internal/ports/input"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
)

type DemoHandler struct {
	service input.DemoService
	logger  *zap.Logger
}

func NewDemoHandler(service input.DemoService, logger *zap.Logger) *DemoHandler {
	return &DemoHandler{service: service, logger: logger}
}

func (h *DemoHandler) GetDemoByID(ctx iris.Context) {
	id, err := ctx.Params().GetUint("id")
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "invalid ID"})
		return
	}

	demo, err := h.service.GetDemoByID(id)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.JSON(demo)
}

func (h *DemoHandler) GenerateRandomCode(ctx iris.Context) {
	code, err := h.service.GenerateRandomCode()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": code})
}
