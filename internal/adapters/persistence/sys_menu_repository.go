// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package persistence

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/domain/model"
	"RuoYi-Go/pkg/db"
	"gorm.io/gorm"
)

type SysMenuRepository struct {
	db *rydb.DatabaseStruct
}

func NewSysMenuRepository(db *rydb.DatabaseStruct) *SysMenuRepository {
	return &SysMenuRepository{db: db}
}

func (this *SysMenuRepository) QueryMenusByUserId(userId int64) ([]*model.SysMenu, error) {
	var menus []*model.SysMenu
	err := this.db.Transactional(func(db *gorm.DB) error {
		if userId == common.ADMINID {
			err := db.Table("sys_menu AS sm").Select("sm.menu_id, sm.parent_id, sm.menu_name, sm.path, sm.component, sm.query, sm.visible, sm.status, sm.perms, sm.is_frame, sm.is_cache, sm.menu_type, sm.icon, sm.order_num, sm.create_time").
				Where("sm.menu_type IN (?, ?) AND sm.status = '0'", "M", "C").
				Order("sm.parent_id, sm.order_num").
				Find(&menus).Error
			return err
		} else {
			err := db.Table("sys_menu AS sm").Select("sm.menu_id, sm.parent_id, sm.menu_name, sm.path, sm.component, sm.query, sm.visible, sm.status, sm.perms, sm.is_frame, sm.is_cache, sm.menu_type, sm.icon, sm.order_num, sm.create_time").
				Joins("LEFT JOIN sys_role_menu srm on srm.menu_id = sm.menu_id").
				Joins("LEFT JOIN sys_user_role sur on sur.role_id = srm.role_id").
				Where("sm.menu_type IN (?, ?) AND sm.status = '0' and srm.user_id = ? ", "M", "C", userId).
				Order("sm.parent_id, sm.order_num").
				Find(&menus).Error
			return err
		}
	})
	if err != nil {
		return nil, err
	}
	return menus, nil
}
