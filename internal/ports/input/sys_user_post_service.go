// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package input

import "RuoYi-Go/internal/domain/model"

// SysUserPostService 输入端口接口
type SysUserPostService interface {
	AddUserPost(post *model.SysUserPost) (*model.SysUserPost, error)
	DeleteUserPostByUserId(userId int64) (int64, error)
}
