// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package output

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/domain/model"
)

type SysDictTypeRepository interface {
	QueryDictTypeByDictID(postId int64) (*model.SysDictType, error)
	QueryDictTypeList(post *model.SysDictTypeRequest) ([]*model.SysDictType, error)
	QueryDictTypePage(pageReq common.PageRequest, r *model.SysDictTypeRequest) ([]*model.SysDictType, int64, error)
}
