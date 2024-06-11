// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package services

import (
	"RuoYi-Go/internal/models"
	rydb "RuoYi-Go/pkg/db"
	"gorm.io/gorm"
)

func GetUserRoles(userId int64) ([]*models.SysRole, error) {
	roles := make([]*models.SysRole, 0)
	err := rydb.DB.Transactional(func(db *gorm.DB) error {
		err := db.Table("sys_role sr").Select("sr.*").
			Joins("LEFT JOIN sys_user_role sur ON sur.role_id = sr.role_id").
			Where("sur.user_id = ?", userId).
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
