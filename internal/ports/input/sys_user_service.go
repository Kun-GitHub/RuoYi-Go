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
	QueryUserInfoByUserId(userId string) (*model.SysUser, error)
	QueryUserPage(pageReq common.PageRequest, userId int64, username string, phone string, status string, deptId int64) (*common.PageResponse, error)
}
