// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package filter

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/domain/model"
	"RuoYi-Go/internal/ports/input"
	"time"

	"github.com/kataras/iris/v12"
)

type OperationLog struct {
	service input.SysOperLogService
}

func NewOperationLog(service input.SysOperLogService) *OperationLog {
	return &OperationLog{service: service}
}

type LogConfig struct {
	Title        string
	BusinessType int
}

func (ol *OperationLog) Log(config LogConfig) iris.Handler {
	return func(ctx iris.Context) {
		startTime := time.Now()

		loginUser, _ := ctx.Values().Get(common.LOGINUSER).(*model.UserInfoStruct)

		ctx.Next()

		costTime := time.Since(startTime).Milliseconds()

		if loginUser != nil {
			status := int32(0)
			if ctx.GetStatusCode() >= 400 {
				status = 1
			}

			operLog := &model.SysOperLog{
				Title:         config.Title,
				BusinessType:  int32(config.BusinessType),
				Method:        ctx.HandlerName(),
				RequestMethod: ctx.Method(),
				OperatorType:  1,
				OperName:      loginUser.UserName,
				OperURL:       ctx.RequestPath(false),
				OperIP:        ctx.RemoteAddr(),
				Status:        status,
				OperTime:      startTime,
				CostTime:      costTime,
			}

			ol.service.Create(operLog)
		}
	}
}
