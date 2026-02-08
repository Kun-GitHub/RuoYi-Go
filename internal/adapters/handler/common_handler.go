package handler

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/ports/input"
	"net/http"

	"github.com/kataras/iris/v12"
)

type CommonHandler struct {
	service input.CommonService
}

func NewCommonHandler(service input.CommonService) *CommonHandler {
	return &CommonHandler{service: service}
}

func (h *CommonHandler) UploadFile(ctx iris.Context) {
	file, header, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Upload file failed: %s", err.Error()))
		return
	}
	defer file.Close()

	url, fileName, err := h.service.UploadFile(header)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "Upload failed: %s", err.Error()))
		return
	}

	ctx.JSON(common.Success(map[string]interface{}{
		"url":              url,
		"fileName":         fileName,
		"newFileName":      fileName, // RuoYi Vue expects newFileName
		"originalFilename": header.Filename,
	}))
}

func (h *CommonHandler) GetResource(ctx iris.Context) {
	resource := ctx.URLParam("resource")
	if resource == "" {
		ctx.StatusCode(http.StatusNotFound)
		return
	}

	filePath, err := h.service.GetResource(resource)
	if err != nil {
		ctx.StatusCode(http.StatusNotFound)
		return
	}

	// Serve the file
	ctx.SendFile(filePath, "")
}
