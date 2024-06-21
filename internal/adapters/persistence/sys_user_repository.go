// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package persistence

import (
	"RuoYi-Go/internal/domain/model"
	"RuoYi-Go/pkg/db"
)

type SysUserRepository struct {
	db *rydb.DatabaseStruct
}

func NewSysUserRepository(db *rydb.DatabaseStruct) *SysUserRepository {
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

func (this *SysUserRepository) QueryUserInfoByUserId(userId string) (*model.SysUser, error) {
	structEntity := &model.SysUser{}
	err := this.db.FindColumns(model.TableNameSysUser, structEntity,
		"user_id = ? and status = '0' and del_flag = '0'", userId)
	if err != nil {
		return nil, err
	}
	return structEntity, nil
}

func (this *SysUserRepository) QueryUserPage(userId int64, username string, phone string, status string, deptId int64) ([]*model.UserList, error) {
	return make([]*model.UserList, 0), nil
}
