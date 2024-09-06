// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package filter

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"io"
	"net/http"
)

// 创建全局的 validator 实例
var validate = validator.New()

// 通用的参数校验中间件
func ValidateRequest(ctx iris.Context, req interface{}) error {
	// 绑定请求数据到结构体
	bodyBytes, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "request body read failed" + err.Error()})
		return err
	}

	err = json.Unmarshal(bodyBytes, req)
	if err != nil {
		fixedJSON := fixJSON(string(bodyBytes))
		if err := json.Unmarshal([]byte(fixedJSON), req); err != nil {
			ctx.StatusCode(http.StatusBadRequest)
			ctx.JSON(iris.Map{"error": "Invalid request payload " + err.Error()})
			return err
		}
	}

	// 进行参数校验
	if err := validate.Struct(req); err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return err
	}
	return nil
}

// fixJSON 修正JSON字符串中的字段，确保它们被双引号包围
func fixJSON(jsonStr string) string {
	fixedStr := jsonStr
	temp := 0

	for i, char := range jsonStr {
		if char == ':' {
			if temp+i+1 < len(fixedStr) && fixedStr[temp+i+1] != '"' {
				fixedStr = fixedStr[:temp+i+1] + "\"" + fixedStr[temp+i+1:]
				temp++
			}
		} else if char == ',' {
			if temp+i-1 < len(fixedStr) && fixedStr[temp+i-1] != '"' {
				fixedStr = fixedStr[:temp+i] + "\"" + fixedStr[temp+i:]
				temp++
			}
		}
	}

	return fixedStr
}
