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
	"golang.org/x/crypto/bcrypt"
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
		"user_name = ? and del_flag = '0'", username)
	if err != nil {
		return nil, err
	}
	return structEntity, nil
}

func (this *SysUserRepository) QueryUserInfoByUserId(userId int64) (*model.SysUser, error) {
	structEntity := &model.SysUser{}
	err := this.db.FindColumns(model.TableNameSysUser, structEntity,
		"user_id = ? and del_flag = '0'", userId)
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
		Where(deptID, status, phonenumber, userName, this.db.Gen.SysUser.DelFlag.Eq("0")).
		Order(this.db.Gen.SysUser.UserID).Limit(pageReq.PageSize).Offset((pageReq.PageNum - 1) * pageReq.PageSize).Find()
	total, err := this.db.Gen.SysUser.WithContext(context.Background()).
		Where(deptID, status, phonenumber, userName, this.db.Gen.SysUser.DelFlag.Eq("0")).
		Order(this.db.Gen.SysUser.UserID).Limit(pageReq.PageSize).Offset((pageReq.PageNum - 1) * pageReq.PageSize).Count()

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
		Where(deptID, status, phonenumber, userName, this.db.Gen.SysUser.DelFlag.Eq("0")).
		Order(this.db.Gen.SysUser.UserID).Find()

	if err != nil {
		return nil, err
	}
	return structEntity, err
}

func (this *SysUserRepository) DeleteUserByUserId(userId int64) (int64, error) {
	r, err := this.db.Gen.SysUser.WithContext(context.Background()).
		Where(this.db.Gen.SysUser.UserID.Eq(userId)).Update(this.db.Gen.SysUser.DelFlag, "2")
	return r.RowsAffected, err
}

func (this *SysUserRepository) ChangeUserStatus(user model.ChangeUserStatusRequest) (int64, error) {
	r, err := this.db.Gen.SysUser.WithContext(context.Background()).
		Where(this.db.Gen.SysUser.UserID.Eq(user.UserID)).Update(this.db.Gen.SysUser.Status, user.Status)
	return r.RowsAffected, err
}

func (this *SysUserRepository) ResetUserPwd(user model.ResetUserPwdRequest) (int64, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	r, err := this.db.Gen.SysUser.WithContext(context.Background()).
		Where(this.db.Gen.SysUser.UserID.Eq(user.UserID)).Update(this.db.Gen.SysUser.Password, string(hashedPassword))
	return r.RowsAffected, err
}
