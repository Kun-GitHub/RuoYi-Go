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
	"time"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
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
		Where(this.db.Gen.SysRole.RoleID.Eq(id)).
		UpdateSimple(this.db.Gen.SysRole.DelFlag.Value("2"),
			this.db.Gen.SysUser.UpdateBy.Value(this.db.User().UserName),
			this.db.Gen.SysUser.UpdateTime.Value(time.Now()))
	return r.RowsAffected, err
}

func (this *SysRoleRepository) ChangeRoleStatus(user *model.ChangeRoleStatusRequest) (int64, error) {
	r, err := this.db.Gen.SysRole.WithContext(context.Background()).
		Where(this.db.Gen.SysRole.RoleID.Eq(user.RoleId), this.db.Gen.SysRole.DelFlag.Eq("0")).
		UpdateSimple(this.db.Gen.SysRole.Status.Value(user.Status),
			this.db.Gen.SysUser.UpdateBy.Value(this.db.User().UserName),
			this.db.Gen.SysUser.UpdateTime.Value(time.Now()))
	return r.RowsAffected, err
}

func (this *SysRoleRepository) SelectRoleAll() ([]*model.SysRole, error) {
	roles := make([]*model.SysRole, 0)
	err := this.db.Gen.SysRole.WithContext(context.Background()).Where(this.db.Gen.SysRole.DelFlag.Eq("0")).Scan(&roles)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (this *SysRoleRepository) QueryAllocatedList(roleId int64, userName, phonenumber string, pageReq common.PageRequest) ([]*model.SysUser, int64, error) {
	users := make([]*model.SysUser, 0)

	// Prepare conditions
	conditions := []gen.Condition{
		this.db.Gen.SysUser.DelFlag.Eq("0"),
		this.db.Gen.SysUserRole.RoleID.Eq(roleId),
	}
	if userName != "" {
		conditions = append(conditions, this.db.Gen.SysUser.UserName.Like("%"+userName+"%"))
	}
	if phonenumber != "" {
		conditions = append(conditions, this.db.Gen.SysUser.Phonenumber.Like("%"+phonenumber+"%"))
	}

	// Count total
	count, err := this.db.Gen.SysUser.WithContext(context.Background()).
		Join(this.db.Gen.SysUserRole, this.db.Gen.SysUser.UserID.EqCol(this.db.Gen.SysUserRole.UserID)).
		Where(conditions...).
		Count()
	if err != nil {
		return nil, 0, err
	}

	// Query data
	err = this.db.Gen.SysUser.WithContext(context.Background()).
		Select(this.db.Gen.SysUser.ALL).
		Join(this.db.Gen.SysUserRole, this.db.Gen.SysUser.UserID.EqCol(this.db.Gen.SysUserRole.UserID)).
		Where(conditions...).
		Limit(pageReq.PageSize).
		Offset((pageReq.PageNum - 1) * pageReq.PageSize).
		Scan(&users)

	if err != nil {
		return nil, 0, err
	}

	return users, count, nil
}

func (this *SysRoleRepository) QueryUnallocatedList(roleId int64, userName, phonenumber string, pageReq common.PageRequest) ([]*model.SysUser, int64, error) {
	users := make([]*model.SysUser, 0)

	// Prepare conditions
	conditions := []gen.Condition{
		this.db.Gen.SysUser.DelFlag.Eq("0"),
		this.db.Gen.SysUser.UserID.Neq(common.ADMINID), // Exclude admin
		this.db.Gen.SysUserRole.UserID.IsNull(),        // UserID in join table is null (meaning no match)
	}
	if userName != "" {
		conditions = append(conditions, this.db.Gen.SysUser.UserName.Like("%"+userName+"%"))
	}
	if phonenumber != "" {
		conditions = append(conditions, this.db.Gen.SysUser.Phonenumber.Like("%"+phonenumber+"%"))
	}

	// Count total
	count, err := this.db.Gen.SysUser.WithContext(context.Background()).
		LeftJoin(this.db.Gen.SysUserRole,
			this.db.Gen.SysUser.UserID.EqCol(this.db.Gen.SysUserRole.UserID),
			this.db.Gen.SysUserRole.RoleID.Eq(roleId)).
		Where(conditions...).
		Count()

	if err != nil {
		return nil, 0, err
	}

	// Query data
	err = this.db.Gen.SysUser.WithContext(context.Background()).
		Select(this.db.Gen.SysUser.ALL).
		LeftJoin(this.db.Gen.SysUserRole,
			this.db.Gen.SysUser.UserID.EqCol(this.db.Gen.SysUserRole.UserID),
			this.db.Gen.SysUserRole.RoleID.Eq(roleId)).
		Where(conditions...).
		Limit(pageReq.PageSize).
		Offset((pageReq.PageNum - 1) * pageReq.PageSize).
		Scan(&users)

	if err != nil {
		return nil, 0, err
	}

	return users, count, nil
}

func (this *SysRoleRepository) InsertAuthUsers(roleId int64, userIds []int64) error {
	userRoles := make([]*model.SysUserRole, 0)
	for _, userId := range userIds {
		userRoles = append(userRoles, &model.SysUserRole{
			UserID: userId,
			RoleID: roleId,
		})
	}
	return this.db.Gen.SysUserRole.WithContext(context.Background()).Create(userRoles...)
}

func (this *SysRoleRepository) DeleteAuthUser(userRole *model.SysUserRole) int64 {
	info, _ := this.db.Gen.SysUserRole.WithContext(context.Background()).
		Where(this.db.Gen.SysUserRole.RoleID.Eq(userRole.RoleID), this.db.Gen.SysUserRole.UserID.Eq(userRole.UserID)).
		Delete()
	return info.RowsAffected
}

func (this *SysRoleRepository) DeleteAuthUsers(roleId int64, userIds []int64) int64 {
	info, _ := this.db.Gen.SysUserRole.WithContext(context.Background()).
		Where(this.db.Gen.SysUserRole.RoleID.Eq(roleId), this.db.Gen.SysUserRole.UserID.In(userIds...)).
		Delete()
	return info.RowsAffected
}
