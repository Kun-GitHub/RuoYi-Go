// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package persistence

import (
	"RuoYi-Go/internal/domain/model"
	"RuoYi-Go/pkg/db"
	"gorm.io/gorm"
)

type SysRoleRepository struct {
	db *rydb.DatabaseStruct
}

func NewSysRoleRepository(db *rydb.DatabaseStruct) *SysRoleRepository {
	return &SysRoleRepository{db: db}
}

func (this *SysRoleRepository) QueryRolesByUserId(userId int64) ([]*model.SysRole, error) {
	roles := make([]*model.SysRole, 0)
	err := this.db.Transactional(func(db *gorm.DB) error {
		err := db.Table("sys_role sr").Select("sr.*").
			Joins("LEFT JOIN sys_user_role sur ON sur.role_id = sr.role_id").
			Where("sr.status = '0' and sur.user_id = ?", userId).
			Find(&roles).Error
		return err
	})
	if err != nil {
		return nil, err
	}
	return roles, nil
}
