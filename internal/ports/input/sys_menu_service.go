// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package input

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/domain/model"
)

// SysMenuService 输入端口接口
type SysMenuService interface {
	QueryMenuByID(id int64) (*model.SysMenu, error)
	QueryMenuList(request *model.SysMenuRequest) ([]*model.SysMenu, error)
	QueryMenuPage(pageReq common.PageRequest, r *model.SysMenuRequest) ([]*model.SysMenu, int64, error)
	AddMenu(post *model.SysMenu) (*model.SysMenu, error)
	EditMenu(post *model.SysMenu) (*model.SysMenu, int64, error)
	DeleteMenuById(id int64) (int64, error)
}
