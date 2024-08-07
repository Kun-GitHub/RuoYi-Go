// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package common

import "fmt"

const SUCCESS = 200

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"` // 使用omitempty确保没有数据时此字段不会出现在JSON中
}

func Success(data interface{}) *Response {
	return &Response{Code: SUCCESS, Message: "操作成功", Data: data}
}

func Error(code int, message string) Response {
	return Response{Code: code, Message: message}
}

func ErrorFormat(code int, format string, a ...any) Response {
	return Response{Code: code, Message: fmt.Sprintf(format, a...)}
}

type PageResponse struct {
	Total   int64  `json:"total"`
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Rows    any    `json:"rows"`
}
