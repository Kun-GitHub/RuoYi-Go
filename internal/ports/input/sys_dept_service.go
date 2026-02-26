// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package input

import (
	"RuoYi-Go/internal/domain/model"
)

// SysDeptService 输入端口接口
type SysDeptService interface {
	QueryDeptList(dept *model.SysDept) ([]*model.SysDept, error)
	QueryDeptListExcludeById(id int64) ([]*model.SysDept, error)
	QueryDeptById(id int64) (*model.SysDept, error)
	AddDept(post *model.SysDept) (*model.SysDept, error)
	EditDept(post *model.SysDept) (*model.SysDept, int64, error)
	DeleteDeptById(id int64) (int64, error)
	QueryChildIdListById(id int64) ([]int64, error)
}
