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

const TableNameSysDictDatum = "sys_dict_data"

// SysDictDatum mapped from table <sys_dict_data>
type SysDictDatum struct {
	DictCode   int64     `gorm:"column:dict_code;primaryKey;comment:字典编码" json:"dictCode"` // 字典编码
	DictSort   int32     `gorm:"column:dict_sort;comment:字典排序" json:"dictSort"`            // 字典排序
	DictLabel  string    `gorm:"column:dict_label;comment:字典标签" json:"dictLabel"`          // 字典标签
	DictValue  string    `gorm:"column:dict_value;comment:字典键值" json:"dictValue"`          // 字典键值
	DictType   string    `gorm:"column:dict_type;comment:字典类型" json:"dictType"`            // 字典类型
	CSSClass   string    `gorm:"column:css_class;comment:样式属性（其他样式扩展）" json:"cssClass"`    // 样式属性（其他样式扩展）
	ListClass  string    `gorm:"column:list_class;comment:表格回显样式" json:"listClass"`        // 表格回显样式
	IsDefault  string    `gorm:"column:is_default;comment:是否默认（Y是 N否）" json:"isDefault"`   // 是否默认（Y是 N否）
	Status     string    `gorm:"column:status;comment:状态（0正常 1停用）" json:"status"`          // 状态（0正常 1停用）
	CreateBy   string    `gorm:"column:create_by;comment:创建者" json:"createBy"`             // 创建者
	CreateTime time.Time `gorm:"column:create_time;comment:创建时间" json:"createTime"`        // 创建时间
	UpdateBy   string    `gorm:"column:update_by;comment:更新者" json:"updateBy"`             // 更新者
	UpdateTime time.Time `gorm:"column:update_time;comment:更新时间" json:"updateTime"`        // 更新时间
	Remark     string    `gorm:"column:remark;comment:备注" json:"remark"`                   // 备注
}

// TableName SysDictDatum's table name
func (*SysDictDatum) TableName() string {
	return TableNameSysDictDatum
}
