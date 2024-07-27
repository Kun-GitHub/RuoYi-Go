// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

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
		err := this.db.Gen.SysRole.WithContext(context.Background()).Select(this.db.Gen.SysRole.ALL).
			LeftJoin(this.db.Gen.SysUserRole, this.db.Gen.SysUserRole.RoleID.EqCol(this.db.Gen.SysRole.RoleID)).
			Where(this.db.Gen.SysUserRole.UserID.Eq(userId), this.db.Gen.SysRole.DelFlag.Eq("0"), this.db.Gen.SysRole.Status.Eq("0")).
			Scan(&roles)
		return err
	})
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (this *SysRoleRepository) QueryRolePage(pageReq common.PageRequest, request *model.SysRoleRequest) ([]*model.SysRole, int64, error) {
	structEntity := make([]*model.SysRole, 0)

	var status field.Expr
	var roleName field.Expr
	var roleKey field.Expr
	var timeField field.Expr
	if request != nil {
		if request.Status != "" {
			status = this.db.Gen.SysRole.Status.Eq(request.Status)
		}
		if request.RoleName != "" {
			roleName = this.db.Gen.SysRole.RoleName.Like("%" + request.RoleName + "%")
		}
		if request.RoleKey != "" {
			roleKey = this.db.Gen.SysRole.RoleKey.Like("%" + request.RoleKey + "%")
		}
		if request.BeginTime != "" && request.EndTime != "" {
			// 解析日期字符串
			t1, err1 := time.Parse("2006-01-02", request.BeginTime)
			t2, err2 := time.Parse("2006-01-02", request.EndTime)
			if err1 == nil && err2 == nil {
				// 设置一天的开始时间
				startOfDay := time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, t1.Location())
				// 设置一天的开始时间
				endOfDay := time.Date(t2.Year(), t2.Month(), t2.Day(), 23, 59, 59, 0, t2.Location())

				timeField = this.db.Gen.SysRole.CreateTime.Between(startOfDay, endOfDay)
			}
		}
	}

	structEntity, total, err := this.db.Gen.SysRole.WithContext(context.Background()).
		Where(status, roleName, roleKey, timeField, this.db.Gen.SysRole.DelFlag.Eq("0")).FindByPage((pageReq.PageNum-1)*pageReq.PageSize, pageReq.PageSize)
	if err != nil {
		return nil, 0, err
	}
	return structEntity, total, err
}

func (this *SysRoleRepository) QueryRoleList(request *model.SysRoleRequest) ([]*model.SysRole, error) {
	structEntity := make([]*model.SysRole, 0)

	var status field.Expr
	var roleName field.Expr
	var roleKey field.Expr
	var timeField field.Expr
	if request != nil {
		if request.Status != "" {
			status = this.db.Gen.SysRole.Status.Eq(request.Status)
		}
		if request.RoleName != "" {
			roleName = this.db.Gen.SysRole.RoleName.Like("%" + request.RoleName + "%")
		}
		if request.RoleKey != "" {
			roleKey = this.db.Gen.SysRole.RoleKey.Like("%" + request.RoleKey + "%")
		}
		if request.BeginTime != "" && request.EndTime != "" {
			// 解析日期字符串
			t1, err1 := time.Parse("2006-01-02", request.BeginTime)
			t2, err2 := time.Parse("2006-01-02", request.EndTime)
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
		Where(status, roleName, roleKey, timeField, this.db.Gen.SysRole.DelFlag.Eq("0")).Find()

	if err != nil {
		return nil, err
	}
	return structEntity, err
}

func (this *SysRoleRepository) QueryRoleByID(id int64) (*model.SysRole, error) {
	structEntity := &model.SysRole{}

	err := this.db.Gen.SysRole.WithContext(context.Background()).
		Where(this.db.Gen.SysRole.RoleID.Eq(id)).Scan(structEntity)

	if err != nil {
		return nil, err
	}
	return structEntity, nil
}

func (this *SysRoleRepository) AddRole(post *model.SysRole) (*model.SysRole, error) {
	post.Status = "0"
	post.DelFlag = "0"

	err := this.db.Gen.SysRole.WithContext(context.Background()).
		Save(post)
	return post, err
}

func (this *SysRoleRepository) EditRole(post *model.SysRole) (*model.SysRole, int64, error) {
	r, err := this.db.Gen.SysRole.WithContext(context.Background()).
		Where(this.db.Gen.SysRole.RoleID.Eq(post.RoleID), this.db.Gen.SysRole.DelFlag.Eq("0")).
		Updates(post)
	return post, r.RowsAffected, err
}

func (this *SysRoleRepository) DeleteRoleById(id int64) (int64, error) {
	r, err := this.db.Gen.SysRole.WithContext(context.Background()).
		Where(this.db.Gen.SysRole.RoleID.Eq(id)).Update(this.db.Gen.SysRole.DelFlag, "2")
	return r.RowsAffected, err
}
