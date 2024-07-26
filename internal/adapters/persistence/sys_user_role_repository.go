// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package persistence

import (
	"RuoYi-Go/internal/adapters/dao"
	"RuoYi-Go/internal/domain/model"
	"context"
)

type SysUserRoleRepository struct {
	db *dao.DatabaseStruct
}

func NewSysUserRoleRepository(db *dao.DatabaseStruct) *SysUserRoleRepository {
	return &SysUserRoleRepository{db: db}
}

func (this *SysUserRoleRepository) AddUserRole(post *model.SysUserRole) (*model.SysUserRole, error) {
	err := this.db.Gen.SysUserRole.WithContext(context.Background()).
		Save(post)
	return post, err
}

func (this *SysUserRoleRepository) DeleteUserRoleByUserId(id int64) (int64, error) {
	r, err := this.db.Gen.SysUserRole.WithContext(context.Background()).
		Where(this.db.Gen.SysUserRole.UserID.Eq(id)).Delete()
	return r.RowsAffected, err
}
