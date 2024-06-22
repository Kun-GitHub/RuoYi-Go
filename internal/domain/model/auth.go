// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package model

type LoginUserStruct struct {
	*SysUser
	Admin bool       `gorm:"-" json:"admin"`
	Roles []*SysRole `gorm:"-" json:"roles"`
	Dept  *SysDept   `gorm:"-" json:"dept"`
}

type LoginSuccess struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Token   string `json:"token"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Code     string `json:"code" validate:"required"`
	Uuid     string `json:"uuid" validate:"required"`
}

type GetInfoSuccess struct {
	Code        int              `json:"code"`
	Message     string           `json:"msg"`
	Permissions []string         `json:"permissions"`
	User        *LoginUserStruct `json:"user"`
	Roles       []string         `json:"roles"`
}
