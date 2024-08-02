// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package output

import "RuoYi-Go/internal/domain/model"

type SysUserPostRepository interface {
	AddUserPost(post *model.SysUserPost) (*model.SysUserPost, error)
	DeleteUserPostByUserId(userId int64) (int64, error)
}
