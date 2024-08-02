// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package input

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/domain/model"
)

// SysUserService 输入端口接口
type SysUserService interface {
	QueryUserByUserName(username string) (*model.SysUser, error)
	QueryUserLikeUserName(username string) ([]*model.SysUser, error)
	QueryUserByUserId(userId int64) (*model.SysUser, error)
	QueryUserPage(pageReq common.PageRequest, user *model.SysUserRequest) ([]*model.UserInfoStruct, int64, error)
	QueryUserList(user *model.SysUserRequest) ([]*model.SysUser, error)
	DeleteUserByUserId(userId int64) (int64, error)
	ChangeUserStatus(user *model.ChangeUserStatusRequest) (int64, error)
	ResetUserPwd(user *model.ResetUserPwdRequest) (int64, error)
	AddUser(post *model.SysUser) (*model.SysUser, error)
	EditUser(post *model.SysUser) (*model.SysUser, int64, error)
	CheckUserNameUnique(id int64, typeStr string) (int64, error)
}
