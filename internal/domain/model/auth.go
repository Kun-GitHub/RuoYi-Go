// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package model

type LoginUserStruct struct {
	*SysUser
	Admin bool       `json:"admin"`
	Roles []*SysRole `json:"roles"`
}

type LoginSuccess struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Token   string `json:"token"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Code     string `json:"code"`
	Uuid     string `json:"uuid"`
}

type GetInfoSuccess struct {
	Code        int              `json:"code"`
	Message     string           `json:"msg"`
	Permissions []string         `json:"permissions"`
	User        *LoginUserStruct `json:"user"`
	Roles       []string         `json:"roles"`
}
