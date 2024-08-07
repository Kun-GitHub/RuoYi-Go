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

const TableNameSysNotice = "sys_notice"

// SysNotice mapped from table <sys_notice>
type SysNotice struct {
	NoticeID      int64     `gorm:"column:notice_id;primaryKey;comment:公告ID" json:"noticeId"`                                // 公告ID
	NoticeTitle   string    `gorm:"column:notice_title;not null;comment:公告标题" json:"noticeTitle" validate:"required"`        // 公告标题
	NoticeType    string    `gorm:"column:notice_type;not null;comment:公告类型（1通知 2公告）" json:"noticeType" validate:"required"` // 公告类型（1通知 2公告）
	NoticeContent string    `gorm:"column:notice_content;comment:公告内容" json:"noticeContent"`                                 // 公告内容
	Status        string    `gorm:"column:status;comment:公告状态（0正常 1关闭）" json:"status"`                                       // 公告状态（0正常 1关闭）
	CreateBy      string    `gorm:"column:create_by;comment:创建者" json:"createBy"`                                            // 创建者
	CreateTime    time.Time `gorm:"column:create_time;comment:创建时间" json:"createTime"`                                       // 创建时间
	UpdateBy      string    `gorm:"column:update_by;comment:更新者" json:"updateBy"`                                            // 更新者
	UpdateTime    time.Time `gorm:"column:update_time;comment:更新时间" json:"updateTime"`                                       // 更新时间
	Remark        string    `gorm:"column:remark;comment:备注" json:"remark"`                                                  // 备注
}

// TableName SysNotice's table name
func (*SysNotice) TableName() string {
	return TableNameSysNotice
}

// SysNoticeRequest mapped from table <SysNoticeRequest>
type SysNoticeRequest struct {
	NoticeTitle string `json:"noticeTitle"`
	NoticeType  string `json:"noticeType"`
	CreateBy    string `json:"createBy"`
}
