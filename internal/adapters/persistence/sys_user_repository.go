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
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gen/field"
	"time"
)

type SysUserRepository struct {
	db *dao.DatabaseStruct
}

func NewSysUserRepository(db *dao.DatabaseStruct) *SysUserRepository {
	return &SysUserRepository{db: db}
}

func (this *SysUserRepository) QueryUserByUserName(username string) (*model.SysUser, error) {
	structEntity := &model.SysUser{}
	err := this.db.Gen.SysUser.WithContext(context.Background()).
		Where(this.db.Gen.SysUser.UserName.Eq(username), this.db.Gen.SysUser.DelFlag.Eq("0")).
		Scan(structEntity)
	if err != nil {
		return nil, err
	}
	return structEntity, nil
}

func (this *SysUserRepository) QueryUserLikeUserName(username string) ([]*model.SysUser, error) {
	structEntity := make([]*model.SysUser, 0)
	structEntity, err := this.db.Gen.SysUser.WithContext(context.Background()).
		Where(this.db.Gen.SysUser.UserName.Like("%"+username+"%"), this.db.Gen.SysUser.DelFlag.Eq("0"),
			this.db.Gen.SysUser.Status.Eq("0")).Find()
	if err != nil {
		return nil, err
	}
	return structEntity, nil
}

func (this *SysUserRepository) QueryUserByUserId(userId int64) (*model.SysUser, error) {
	structEntity := &model.SysUser{}
	err := this.db.Gen.SysUser.WithContext(context.Background()).
		Where(this.db.Gen.SysUser.UserID.Eq(userId), this.db.Gen.SysUser.DelFlag.Eq("0")).
		Scan(structEntity)
	if err != nil {
		return nil, err
	}
	return structEntity, nil
}

func (this *SysUserRepository) QueryUserPage(pageReq common.PageRequest, user *model.SysUserRequest) ([]*model.SysUser, int64, error) {
	structEntity := make([]*model.SysUser, 0)

	//total, err := this.db.PageQuery(func(db *gorm.DB) *gorm.DB {
	//	return db.Table(model.TableNameSysUser).Where("status = '0' and del_flag = '0'")
	//}, pageReq, &structEntity)

	var status field.Expr
	var deptID field.Expr
	var phonenumber field.Expr
	var userName field.Expr
	var timeField field.Expr
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
		if user.BeginTime != "" && user.EndTime != "" {
			// 解析日期字符串
			t1, err1 := time.Parse("2006-01-02", user.BeginTime)
			t2, err2 := time.Parse("2006-01-02", user.EndTime)
			if err1 == nil && err2 == nil {
				// 设置一天的开始时间
				startOfDay := time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, t1.Location())
				// 设置一天的开始时间
				endOfDay := time.Date(t2.Year(), t2.Month(), t2.Day(), 23, 59, 59, 0, t2.Location())

				timeField = this.db.Gen.SysUser.CreateTime.Between(startOfDay, endOfDay)
			}
		}
	}

	structEntity, total, err := this.db.Gen.SysUser.WithContext(context.Background()).
		Where(deptID, status, phonenumber, userName, timeField, this.db.Gen.SysUser.DelFlag.Eq("0")).
		Order(this.db.Gen.SysUser.UserID).FindByPage((pageReq.PageNum-1)*pageReq.PageSize, pageReq.PageSize)
	if err != nil {
		return nil, 0, err
	}
	return structEntity, total, err
}

func (this *SysUserRepository) QueryUserList(user *model.SysUserRequest) ([]*model.SysUser, error) {
	structEntity := make([]*model.SysUser, 0)

	var status field.Expr
	var deptID field.Expr
	var phonenumber field.Expr
	var userName field.Expr
	var timeField field.Expr
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
		if user.BeginTime != "" && user.EndTime != "" {
			// 解析日期字符串
			t1, err1 := time.Parse("2006-01-02", user.BeginTime)
			t2, err2 := time.Parse("2006-01-02", user.EndTime)
			if err1 == nil && err2 == nil {
				// 设置一天的开始时间
				startOfDay := time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, t1.Location())
				// 设置一天的开始时间
				endOfDay := time.Date(t2.Year(), t2.Month(), t2.Day(), 23, 59, 59, 0, t2.Location())

				timeField = this.db.Gen.SysUser.CreateTime.Between(startOfDay, endOfDay)
			}
		}
	}

	structEntity, err := this.db.Gen.SysUser.WithContext(context.Background()).
		Where(deptID, status, phonenumber, userName, timeField, this.db.Gen.SysUser.DelFlag.Eq("0")).
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

func (this *SysUserRepository) ChangeUserStatus(user *model.ChangeUserStatusRequest) (int64, error) {
	r, err := this.db.Gen.SysUser.WithContext(context.Background()).
		Where(this.db.Gen.SysUser.UserID.Eq(user.UserID), this.db.Gen.SysUser.DelFlag.Eq("0")).
		Update(this.db.Gen.SysUser.Status, user.Status)
	return r.RowsAffected, err
}

func (this *SysUserRepository) ResetUserPwd(user *model.ResetUserPwdRequest) (int64, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	r, err := this.db.Gen.SysUser.WithContext(context.Background()).
		Where(this.db.Gen.SysUser.UserID.Eq(user.UserID), this.db.Gen.SysUser.DelFlag.Eq("0")).
		Update(this.db.Gen.SysUser.Password, string(hashedPassword))
	return r.RowsAffected, err
}

func (this *SysUserRepository) AddUser(post *model.SysUser) (*model.SysUser, error) {
	post.Status = "0"
	post.DelFlag = "0"

	err := this.db.Gen.SysUser.WithContext(context.Background()).
		Save(post)
	return post, err
}

func (this *SysUserRepository) EditUser(post *model.SysUser) (*model.SysUser, int64, error) {
	r, err := this.db.Gen.SysUser.WithContext(context.Background()).
		Where(this.db.Gen.SysUser.UserID.Eq(post.UserID), this.db.Gen.SysUser.DelFlag.Eq("0")).
		Updates(post)
	return post, r.RowsAffected, err
}

func (this *SysUserRepository) CheckUserNameUnique(id int64, name string) (int64, error) {
	r, err := this.db.Gen.SysUser.WithContext(context.Background()).
		Where(this.db.Gen.SysUser.UserName.Eq(name), this.db.Gen.SysUser.UserID.Neq(id), this.db.Gen.SysUser.DelFlag.Eq("0")).Count()
	return r, err
}
