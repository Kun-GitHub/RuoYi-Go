// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package persistence

import (
	"RuoYi-Go/internal/adapters/dao"
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/domain/model"
	"context"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"time"
)

type SysRoleRepository struct {
	db *dao.DatabaseStruct
}

func NewSysRoleRepository(db *dao.DatabaseStruct) *SysRoleRepository {
	return &SysRoleRepository{db: db}
}

func (this *SysRoleRepository) QueryRolesByUserId(userId int64) ([]*model.SysRole, error) {
	roles := make([]*model.SysRole, 0)
	err := this.db.Transactional(func(db *gorm.DB) error {
		err := db.Table("sys_role sr").Select("sr.*").
			Joins("LEFT JOIN sys_user_role sur ON sur.role_id = sr.role_id").
			Where("sr.status = '0' and sr.del_flag = '0' and sur.user_id = ?", userId).
			Find(&roles).Error
		return err
	})
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (this *SysRoleRepository) QueryRolePage(pageReq common.PageRequest, user *model.SysRoleRequest) ([]*model.SysRole, int64, error) {
	structEntity := make([]*model.SysRole, 0)

	var status field.Expr
	var roleName field.Expr
	var roleKey field.Expr
	var timeField field.Expr
	if user != nil {
		if user.Status != "" {
			status = this.db.Gen.SysRole.Status.Eq(user.Status)
		}
		if user.RoleName != "" {
			roleName = this.db.Gen.SysRole.RoleName.Like("%" + user.RoleName + "%")
		}
		if user.RoleKey != "" {
			roleKey = this.db.Gen.SysRole.RoleKey.Like("%" + user.RoleKey + "%")
		}
		if user.BeginTime != "" && user.EndTime != "" {
			// 解析日期字符串
			t1, err1 := time.Parse("2006-01-02", user.BeginTime)
			t2, err2 := time.Parse("2006-01-02", user.EndTime)
			if err1 == nil && err2 == nil {
				// 设置一天的开始时间
				startOfDay := time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, t1.Location())
				// 设置一天的开始时间
				endOfDay := time.Date(t2.Year(), t2.Month(), t2.Day(), 23, 59, 59, 0, t2.Location())

				timeField = this.db.Gen.SysRole.CreateTime.Between(startOfDay, endOfDay)
			}
		}
	}

	structEntity, err := this.db.Gen.SysRole.WithContext(context.Background()).
		Where(status, roleName, roleKey, timeField).Limit(pageReq.PageSize).Offset((pageReq.PageNum - 1) * pageReq.PageSize).Find()
	total, err := this.db.Gen.SysRole.WithContext(context.Background()).
		Where(status, roleName, roleKey, timeField).Limit(pageReq.PageSize).Offset((pageReq.PageNum - 1) * pageReq.PageSize).Count()

	if err != nil {
		return nil, 0, err
	}
	return structEntity, total, err
}