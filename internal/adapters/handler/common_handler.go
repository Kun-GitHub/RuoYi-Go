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

func (h *CommonHandler) Download(ctx iris.Context) {
	fileName := ctx.URLParam("fileName")
	_ = ctx.URLParam("delete")
	// Service method to handle download
	filePath, err := h.service.Download(fileName)
	if err != nil {
		ctx.StatusCode(http.StatusNotFound)
		return
	}
	ctx.SendFile(filePath, fileName)
}

func (h *CommonHandler) UploadFiles(ctx iris.Context) {
	// Multiple file upload
	err := ctx.Request().ParseMultipartForm(32 << 20)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Upload failed: %s", err.Error()))
		return
	}

	files := ctx.Request().MultipartForm.File["files"]
	var results []map[string]interface{}
	for _, header := range files {
		url, fileName, err := h.service.UploadFile(header)
		if err != nil {
			continue
		}
		results = append(results, map[string]interface{}{
			"url":              url,
			"fileName":         fileName,
			"newFileName":      fileName,
			"originalFilename": header.Filename,
		})
	}
	ctx.JSON(common.Success(results))
}
