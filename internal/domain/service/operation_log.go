// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package service

import "strings"

// BusinessType constants matching RuoYi Java
const (
	BusinessTypeOther   = 0
	BusinessTypeInsert  = 1
	BusinessTypeUpdate  = 2
	BusinessTypeDelete  = 3
	BusinessTypeGrant   = 4
	BusinessTypeExport  = 5
	BusinessTypeImport  = 6
	BusinessTypeForce   = 7
	BusinessTypeClean   = 8
	BusinessTypeGenCode = 9
)

// OperationLogResolver resolves operation log info from URL and method
type OperationLogResolver struct{}

func NewOperationLogResolver() *OperationLogResolver {
	return &OperationLogResolver{}
}

type OperationInfo struct {
	Title        string
	BusinessType int
}

// Resolve determines the operation title and business type from request path and method
func (r *OperationLogResolver) Resolve(path, method string) OperationInfo {
	title := "其他"
	bizType := BusinessTypeOther

	switch {
	case strings.HasPrefix(path, "/system/user"):
		title = "用户管理"
	case strings.HasPrefix(path, "/system/role"):
		title = "角色管理"
	case strings.HasPrefix(path, "/system/menu"):
		title = "菜单管理"
	case strings.HasPrefix(path, "/system/dept"):
		title = "部门管理"
	case strings.HasPrefix(path, "/system/post"):
		title = "岗位管理"
	case strings.HasPrefix(path, "/system/dict/type"):
		title = "字典类型"
	case strings.HasPrefix(path, "/system/dict/data"):
		title = "字典数据"
	case strings.HasPrefix(path, "/system/config"):
		title = "参数管理"
	case strings.HasPrefix(path, "/system/notice"):
		title = "通知公告"
	case strings.HasPrefix(path, "/system/user/profile"):
		title = "个人信息"
	case strings.HasPrefix(path, "/monitor/online"):
		title = "在线用户"
	case strings.HasPrefix(path, "/monitor/operlog"):
		title = "操作日志"
	case strings.HasPrefix(path, "/monitor/logininfor"):
		title = "登录日志"
	case strings.HasPrefix(path, "/monitor/job"):
		title = "定时任务"
	case strings.HasPrefix(path, "/monitor/cache"):
		title = "缓存监控"
	case strings.HasPrefix(path, "/monitor/server"):
		title = "服务器监控"
	case strings.HasPrefix(path, "/common/upload"):
		title = "文件管理"
	}

	switch method {
	case "POST":
		if strings.HasSuffix(path, "/export") {
			bizType = BusinessTypeExport
		} else if strings.Contains(path, "/import") {
			bizType = BusinessTypeImport
		} else {
			bizType = BusinessTypeInsert
		}
	case "PUT":
		bizType = BusinessTypeUpdate
	case "DELETE":
		if strings.HasSuffix(path, "/clean") {
			bizType = BusinessTypeClean
		} else {
			bizType = BusinessTypeDelete
		}
	}

	return OperationInfo{Title: title, BusinessType: bizType}
}
