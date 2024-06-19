// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package filter

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"net/http"
)

// 创建全局的 validator 实例
var validate = validator.New()

// 通用的参数校验中间件
func ValidateRequest(ctx iris.Context, req interface{}) error {
	// 绑定请求数据到结构体
	if err := ctx.ReadJSON(req); err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid request payload"})
		return err
	}

	// 进行参数校验
	if err := validate.Struct(req); err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return err
	}
	return nil
}
