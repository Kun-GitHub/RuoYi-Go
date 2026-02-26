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

// SysUserRepository 用户仓储实现类
// 负责用户数据的持久化操作，基于GORM Gen框架实现
type SysUserRepository struct {
	db *dao.DatabaseStruct
}

// NewSysUserRepository 创建用户仓储实例
// 参数:
//   - db: 数据库结构体
//
// 返回值: 用户仓储
func NewSysUserRepository(db *dao.DatabaseStruct) *SysUserRepository {
	return &SysUserRepository{db: db}
}

// QueryUserByUserName 根据用户名查询用户
// 从数据库中查询未被删除的指定用户名的用户信息
// 参数:
//   - username: 用户名
//
// 返回值:
//   - *model.SysUser: 用户信息
//   - error: 错误信息
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

// QueryUserLikeUserName 模糊查询用户名
// 根据用户名关键字进行模糊匹配，查询正常的未删除用户
// 参数:
//   - username: 用户名关键字
//
// 返回值:
//   - []*model.SysUser: 符合条件的用户列表
//   - error: 错误信息
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

// QueryUserByUserId 根据用户ID查询用户
// 从数据库中查询未被删除的指定ID的用户信息
// 参数:
//   - userId: 用户ID
//
// 返回值:
//   - *model.SysUser: 用户信息
//   - error: 错误信息
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

// QueryUserPage 分页查询用户列表
// 支持多种筛选条件的分页查询，返回用户列表和总记录数
// 参数:
//   - pageReq: 分页请求参数
//   - user: 用户查询条件
//
// 返回值:
//   - []*model.SysUser: 用户列表
//   - int64: 总记录数
//   - error: 错误信息
func (this *SysUserRepository) QueryUserPage(pageReq common.PageRequest, user *model.SysUserRequest) ([]*model.SysUser, int64, error) {
	structEntity := make([]*model.SysUser, 0)

	//total, err := this.db.PageQuery(func(db *gorm.DB) *gorm.DB {
	//	return db.Table(model.TableNameSysUser).Where("status = '0' and del_flag = '0'")
	//}, pageReq, &structEntity)

	var status field.Expr
	var deptID field.Expr
	var deptIDs field.Expr
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
		if len(user.DeptIDs) > 0 {
			deptID = this.db.Gen.SysUser.DeptID.In(user.DeptIDs...)
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
		Where(deptID, status, phonenumber, userName, timeField, deptIDs, this.db.Gen.SysUser.DelFlag.Eq("0")).
		Order(this.db.Gen.SysUser.UserID).FindByPage((pageReq.PageNum-1)*pageReq.PageSize, pageReq.PageSize)
	if err != nil {
		return nil, 0, err
	}
	return structEntity, total, err
}

// QueryUserList 查询用户列表
// 不分页的用户列表查询，支持多种筛选条件
// 参数:
//   - user: 用户查询条件
//
// 返回值:
//   - []*model.SysUser: 用户列表
//   - error: 错误信息
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

// DeleteUserByUserId 逻辑删除用户
// 将用户标记为删除状态（del_flag=2），并更新操作信息
// 参数:
//   - userId: 用户ID
//
// 返回值:
//   - int64: 影响的行数
//   - error: 错误信息
func (this *SysUserRepository) DeleteUserByUserId(userId int64) (int64, error) {
	r, err := this.db.Gen.SysUser.WithContext(context.Background()).
		Where(this.db.Gen.SysUser.UserID.Eq(userId)).
		UpdateSimple(this.db.Gen.SysUser.DelFlag.Value("2"),
			this.db.Gen.SysUser.UpdateBy.Value(this.db.User().UserName),
			this.db.Gen.SysUser.UpdateTime.Value(time.Now()))
	return r.RowsAffected, err
}

// ChangeUserStatus 修改用户状态
// 启用或禁用用户账户，更新用户状态和操作信息
// 参数:
//   - user: 用户状态修改请求
//
// 返回值:
//   - int64: 影响的行数
//   - error: 错误信息
func (this *SysUserRepository) ChangeUserStatus(user *model.ChangeUserStatusRequest) (int64, error) {
	r, err := this.db.Gen.SysUser.WithContext(context.Background()).
		Where(this.db.Gen.SysUser.UserID.Eq(user.UserID), this.db.Gen.SysUser.DelFlag.Eq("0")).
		UpdateSimple(this.db.Gen.SysUser.Status.Value(user.Status),
			this.db.Gen.SysUser.UpdateBy.Value(this.db.User().UserName),
			this.db.Gen.SysUser.UpdateTime.Value(time.Now()))
	return r.RowsAffected, err
}

// ResetUserPwd 重置用户密码
// 为用户设置新的加密密码，并更新操作信息
// 参数:
//   - user: 密码重置请求
//
// 返回值:
//   - int64: 影响的行数
//   - error: 错误信息
func (this *SysUserRepository) ResetUserPwd(user *model.ResetUserPwdRequest) (int64, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	r, err := this.db.Gen.SysUser.WithContext(context.Background()).
		Where(this.db.Gen.SysUser.UserID.Eq(user.UserID), this.db.Gen.SysUser.DelFlag.Eq("0")).
		UpdateSimple(this.db.Gen.SysUser.Password.Value(string(hashedPassword)),
			this.db.Gen.SysUser.UpdateBy.Value(this.db.User().UserName),
			this.db.Gen.SysUser.UpdateTime.Value(time.Now()))
	return r.RowsAffected, err
}

// AddUser 添加新用户
// 创建新的用户记录，设置默认状态和删除标记
// 参数:
//   - post: 用户信息
//
// 返回值:
//   - *model.SysUser: 创建的用户信息
//   - error: 错误信息
func (this *SysUserRepository) AddUser(post *model.SysUser) (*model.SysUser, error) {
	post.Status = "0"
	post.DelFlag = "0"

	err := this.db.Gen.SysUser.WithContext(context.Background()).
		Save(post)
	return post, err
}

// EditUser 编辑用户信息
// 更新现有用户的完整信息
// 参数:
//   - post: 更新后的用户信息
//
// 返回值:
//   - *model.SysUser: 更新后的用户信息
//   - int64: 影响的行数
//   - error: 错误信息
func (this *SysUserRepository) EditUser(post *model.SysUser) (*model.SysUser, int64, error) {
	r, err := this.db.Gen.SysUser.WithContext(context.Background()).
		Where(this.db.Gen.SysUser.UserID.Eq(post.UserID), this.db.Gen.SysUser.DelFlag.Eq("0")).
		Updates(post)
	return post, r.RowsAffected, err
}

// CheckUserNameUnique 检查用户名唯一性
// 验证指定用户名是否已被其他未删除用户使用
// 参数:
//   - id: 排除的用户ID（编辑时使用）
//   - name: 待检查的用户名
//
// 返回值:
//   - int64: 重复数量
//   - error: 错误信息
func (this *SysUserRepository) CheckUserNameUnique(id int64, name string) (int64, error) {
	r, err := this.db.Gen.SysUser.WithContext(context.Background()).
		Where(this.db.Gen.SysUser.UserName.Eq(name), this.db.Gen.SysUser.UserID.Neq(id), this.db.Gen.SysUser.DelFlag.Eq("0")).Count()
	return r, err
}

// UserLogin 记录用户登录
// 更新用户的最后登录时间
// 参数:
//   - user: 用户信息
//
// 返回值:
//   - int64: 影响的行数
//   - error: 错误信息
func (this *SysUserRepository) UserLogin(user *model.SysUser) (int64, error) {
	r, err := this.db.Gen.SysUser.WithContext(context.Background()).
		Where(this.db.Gen.SysUser.UserID.Eq(user.UserID), this.db.Gen.SysUser.DelFlag.Eq("0")).
		Update(this.db.Gen.SysUser.LoginDate, time.Now())
	return r.RowsAffected, err
}
