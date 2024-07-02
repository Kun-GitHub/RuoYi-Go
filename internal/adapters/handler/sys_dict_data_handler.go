// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package handler

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/ports/input"
	"github.com/kataras/iris/v12"
)

type SysDictDataHandler struct {
	service input.SysDictDataService
}

func NewSysDictDataHandler(service input.SysDictDataService) *SysDictDataHandler {
	return &SysDictDataHandler{service: service}
}

// GenerateCaptchaImage
func (h *SysDictDataHandler) DictType(ctx iris.Context) {
	dictType := ctx.Params().GetString("dictType")
	if dictType == "" {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "invalid ID"})
		return
	}

	s, _ := h.service.QueryDictDatasByType(dictType)
	ctx.JSON(common.Success(s))
}
