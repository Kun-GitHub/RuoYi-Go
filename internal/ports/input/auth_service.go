// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package input

import (
	"RuoYi-Go/internal/domain/model"
)

// AuthService 输入端口接口
type AuthService interface {
	Login(l model.LoginRequest) (*model.LoginSuccess, error)
	Logout(token string) error
	GetInfo(loginUser *model.UserInfoStruct) (*model.UserInfoStruct, []string, []string, error)
}
