// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysLogininfor = "sys_logininfor"

// SysLogininfor mapped from table <sys_logininfor>
type SysLogininfor struct {
	InfoID        int64     `gorm:"column:info_id;primaryKey;comment:访问ID" json:"infoId"`    // 访问ID
	UserName      string    `gorm:"column:user_name;comment:用户账号" json:"userName"`           // 用户账号
	Ipaddr        string    `gorm:"column:ipaddr;comment:登录IP地址" json:"ipaddr"`              // 登录IP地址
	LoginLocation string    `gorm:"column:login_location;comment:登录地点" json:"loginLocation"` // 登录地点
	Browser       string    `gorm:"column:browser;comment:浏览器类型" json:"browser"`             // 浏览器类型
	Os            string    `gorm:"column:os;comment:操作系统" json:"os"`                        // 操作系统
	Status        string    `gorm:"column:status;comment:登录状态（0成功 1失败）" json:"status"`       // 登录状态（0成功 1失败）
	Msg           string    `gorm:"column:msg;comment:提示消息" json:"msg"`                      // 提示消息
	LoginTime     time.Time `gorm:"column:login_time;comment:访问时间" json:"loginTime"`         // 访问时间
}

// TableName SysLogininfor's table name
func (*SysLogininfor) TableName() string {
	return TableNameSysLogininfor
}

// SysLogininforRequest mapped from table <SysLogininforRequest>
type SysLogininforRequest struct {
	Status    string `json:"status"`
	Ipaddr    string `json:"ipaddr"`
	UserName  string `json:"userName"`
	BeginTime string `json:"beginTime"`
	EndTime   string `json:"endTime"`
}
