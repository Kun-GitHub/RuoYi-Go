// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package output

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/domain/model"
)

type SysConfigRepository interface {
	QueryConfigByID(postId int64) (*model.SysConfig, error)
	QueryConfigList(post *model.SysConfigRequest) ([]*model.SysConfig, error)
	QueryConfigPage(pageReq common.PageRequest, r *model.SysConfigRequest) ([]*model.SysConfig, int64, error)
	AddConfig(post *model.SysConfig) (*model.SysConfig, error)
	EditConfig(post *model.SysConfig) (*model.SysConfig, int64, error)
	DeleteConfigById(id int64) (int64, error)
	CheckConfigNameUnique(id int64, name string) (int64, error)
}
