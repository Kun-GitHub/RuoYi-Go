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
)

type SysUserRepository struct {
	db *dao.DatabaseStruct
}

func NewSysUserRepository(db *dao.DatabaseStruct) *SysUserRepository {
	return &SysUserRepository{db: db}
}

func (this *SysUserRepository) QueryUserInfoByUserName(username string) (*model.SysUser, error) {
	structEntity := &model.SysUser{}
	err := this.db.FindColumns(model.TableNameSysUser, structEntity,
		"user_name = ? and status = '0' and del_flag = '0'", username)
	if err != nil {
		return nil, err
	}
	return structEntity, nil
}

func (this *SysUserRepository) QueryUserInfoByUserId(userId int64) (*model.SysUser, error) {
	structEntity := &model.SysUser{}
	err := this.db.FindColumns(model.TableNameSysUser, structEntity,
		"user_id = ? and status = '0' and del_flag = '0'", userId)
	if err != nil {
		return nil, err
	}
	return structEntity, nil
}

func (this *SysUserRepository) QueryUserPage(pageReq common.PageRequest, user *model.SysUser) ([]*model.SysUser, int64, error) {
	structEntity := make([]*model.SysUser, 0)

	//total, err := this.db.PageQuery(func(db *gorm.DB) *gorm.DB {
	//	return db.Table(model.TableNameSysUser).Where("status = '0' and del_flag = '0'")
	//}, pageReq, &structEntity)

	var status field.Expr
	var deptID field.Expr
	var phonenumber field.Expr
	var userName field.Expr
	if user != nil {
		if user.Status != "" {
			status = this.db.Gen.SysUser.Status.Eq(user.Status)
		}
		if user.Phonenumber != "" {
			phonenumber = this.db.Gen.SysUser.Phonenumber.Like("%" + user.Phonenumber + "%")
		}
		if user.UserName != "" {
			userName = this.db.Gen.SysUser.UserName.Like("%" + user.UserName + "%")
		}
		if user.DeptID != 0 {
			deptID = this.db.Gen.SysUser.DeptID.Eq(user.DeptID)
		}
	}

	structEntity, err := this.db.Gen.SysUser.WithContext(context.Background()).
		Where(deptID, status, phonenumber, userName).Limit(pageReq.PageSize).Offset((pageReq.PageNum - 1) * pageReq.PageSize).Find()
	total, err := this.db.Gen.SysUser.WithContext(context.Background()).
		Where(deptID, status, phonenumber, userName).Limit(pageReq.PageSize).Offset((pageReq.PageNum - 1) * pageReq.PageSize).Count()

	if err != nil {
		return nil, 0, err
	}
	return structEntity, total, err
}

func (this *SysUserRepository) QueryUserList(user *model.SysUser) ([]*model.SysUser, error) {
	structEntity := make([]*model.SysUser, 0)

	var status field.Expr
	var deptID field.Expr
	var phonenumber field.Expr
	var userName field.Expr
	if user != nil {
		if user.Status != "" {
			status = this.db.Gen.SysUser.Status.Eq(user.Status)
		}
		if user.Phonenumber != "" {
			phonenumber = this.db.Gen.SysUser.Phonenumber.Like("%" + user.Phonenumber + "%")
		}
		if user.UserName != "" {
			userName = this.db.Gen.SysUser.UserName.Like("%" + user.UserName + "%")
		}
		if user.DeptID != 0 {
			deptID = this.db.Gen.SysUser.DeptID.Eq(user.DeptID)
		}
	}

	structEntity, err := this.db.Gen.SysUser.WithContext(context.Background()).
		Where(deptID, status, phonenumber, userName).Find()

	if err != nil {
		return nil, err
	}
	return structEntity, err
}
