// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package service

import (
	"RuoYi-Go/internal/domain/model"
	rydb "RuoYi-Go/pkg/db"
	"gorm.io/gorm"
)

type SysMenuStruct struct {
	model.SysMenu
	ParentName string           `gorm:"-" json:"parentName,omitempty"` // gorm:"-" 忽略
	Children   []*SysMenuStruct `gorm:"-" json:"children,omitempty"`
}

func QueryAllMenus() ([]*SysMenuStruct, error) {
	var menus []*SysMenuStruct
	err := rydb.DB.Transactional(func(db *gorm.DB) error {
		err := db.Table("sys_menu AS m").Select("m.menu_id, m.parent_id, m.menu_name, m.path, m.component, m.query, m.visible, m.status, COALESCE(m.perms, '') AS perms, m.is_frame, m.is_cache, m.menu_type, m.icon, m.order_num, m.create_time").
			Where("m.menu_type IN (?, ?) AND m.status = '0'", "M", "C").
			Order("m.parent_id, m.order_num").
			Find(&menus).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func QueryMenusByUserId(userId int64) ([]*model.SysMenu, error) {
	roles := make([]*model.SysMenu, 0)
	err := rydb.DB.Transactional(func(db *gorm.DB) error {
		err := db.Table("sys_menu sm").Distinct("sm.*").
			Joins("left join sys_role_menu srm on srm.menu_id = sm.menu_id").
			Joins("left join sys_user_role sur on sur.role_id = srm.role_id").
			Where("sm.status = '0' and sur.user_id = ?", userId).Order("sm.menu_id").
			Find(&roles).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return roles, nil
}
