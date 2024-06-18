// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package persistence

import (
	"RuoYi-Go/internal/domain/model"
	"RuoYi-Go/pkg/db"
)

type SysDeptRepository struct {
	db *rydb.DatabaseStruct
}

func NewSysDeptRepository(db *rydb.DatabaseStruct) *SysDeptRepository {
	return &SysDeptRepository{db: db}
}

func (this *SysDeptRepository) QueryRolesByDeptId(deptId int64) (*model.SysDept, error) {
	structEntity := &model.SysDept{}
	err := this.db.FindColumns(model.TableNameSysDept, structEntity,
		"dept_id = ? and status = '0' and del_flag = '0'", deptId)
	if err != nil {
		return nil, err
	}
	return structEntity, nil
}
