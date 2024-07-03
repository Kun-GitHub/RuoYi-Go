// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package input

import (
	"RuoYi-Go/internal/domain/model"
)

// SysDeptService 输入端口接口
type SysDeptService interface {
	QueryRolesByDeptId(deptId int64) (*model.SysDept, error)
	QueryDeptList(dept *model.SysDept) ([]*model.SysDept, error)
}
