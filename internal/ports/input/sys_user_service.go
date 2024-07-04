// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package input

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/domain/model"
)

// SysUserService 输入端口接口
type SysUserService interface {
	QueryUserInfoByUserName(username string) (*model.SysUser, error)
	QueryUserInfoByUserId(userId int64) (*model.SysUser, error)
	QueryUserPage(pageReq common.PageRequest, user *model.SysUser) ([]*model.UserInfoStruct, int64, error)
	QueryUserList(user *model.SysUser) ([]*model.SysUser, error)
}
