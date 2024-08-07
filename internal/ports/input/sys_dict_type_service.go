// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package input

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/domain/model"
)

// SysDictTypeService 输入端口接口
type SysDictTypeService interface {
	QueryDictTypeByDictID(postId int64) (*model.SysDictType, error)
	QueryDictTypeList(post *model.SysDictTypeRequest) ([]*model.SysDictType, error)
	QueryDictTypePage(pageReq common.PageRequest, r *model.SysDictTypeRequest) ([]*model.SysDictType, int64, error)
	AddDictType(post *model.SysDictType) (*model.SysDictType, error)
	EditDictType(post *model.SysDictType) (*model.SysDictType, int64, error)
	DeleteDictTypeById(id int64) (int64, error)
	CheckDictTypeUnique(id int64, typeStr string) (int64, error)
}
