// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package service

import (
	"RuoYi-Go/internal/domain/model"
	rydb "RuoYi-Go/pkg/db"
	"gorm.io/gorm"
)

func QueryRolesByUserId(userId int64) ([]*model.SysRole, error) {
	roles := make([]*model.SysRole, 0)
	err := rydb.DB.Transactional(func(db *gorm.DB) error {
		err := db.Table("sys_role sr").Select("sr.*").
			Joins("LEFT JOIN sys_user_role sur ON sur.role_id = sr.role_id").
			Where("sr.status = '0' and sur.user_id = ?", userId).
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
