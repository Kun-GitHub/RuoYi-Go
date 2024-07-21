// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package persistence

import (
	"RuoYi-Go/internal/adapters/dao"
	"RuoYi-Go/internal/domain/model"
	"context"
	"gorm.io/gen/field"
)

type SysDeptRepository struct {
	db *dao.DatabaseStruct
}

func NewSysDeptRepository(db *dao.DatabaseStruct) *SysDeptRepository {
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

func (this *SysDeptRepository) QueryDeptList(dept *model.SysDept) ([]*model.SysDept, error) {
	structEntity := make([]*model.SysDept, 0)

	var status field.Expr
	var deptName field.Expr
	if dept != nil {
		if dept.Status != "" {
			status = this.db.Gen.SysDept.Status.Eq(dept.Status)
		}

		if dept.DeptName != "" {
			deptName = this.db.Gen.SysDept.DeptName.Like("%" + dept.DeptName + "%")
		}
	}

	structEntity, err := this.db.Gen.SysDept.WithContext(context.Background()).
		Where(deptName, status).Find()
	if err != nil {
		return nil, err
	}
	return structEntity, nil
}
