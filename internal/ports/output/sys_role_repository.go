// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package output

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/domain/model"
)

type SysRoleRepository interface {
	QueryRolesByUserId(userId int64) ([]*model.SysRole, error)
	QueryRolePage(pageReq common.PageRequest, user *model.SysRoleRequest) ([]*model.SysRole, int64, error)
}
