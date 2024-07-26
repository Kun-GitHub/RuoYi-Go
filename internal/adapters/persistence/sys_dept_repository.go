// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
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
		Where(deptName, status, this.db.Gen.SysDept.DelFlag.Eq("0")).Find()
	if err != nil {
		return nil, err
	}
	return structEntity, nil
}

func (this *SysDeptRepository) QueryDeptListExcludeById(id int64) ([]*model.SysDept, error) {
	structEntity := make([]*model.SysDept, 0)

	var idField field.Expr
	if id != 0 {
		idField = this.db.Gen.SysDept.DeptID.Neq(id)
	}

	structEntity, err := this.db.Gen.SysDept.WithContext(context.Background()).
		Where(idField, this.db.Gen.SysDept.DelFlag.Eq("0")).Find()
	if err != nil {
		return nil, err
	}
	return structEntity, nil
}

func (this *SysDeptRepository) QueryDeptById(id int64) (*model.SysDept, error) {
	structEntity := &model.SysDept{}

	err := this.db.Gen.SysDept.WithContext(context.Background()).
		Where(this.db.Gen.SysDept.DeptID.Eq(id)).Scan(structEntity)

	if err != nil {
		return nil, err
	}
	return structEntity, nil
}

func (this *SysDeptRepository) AddDept(post *model.SysDept) (*model.SysDept, error) {
	post.Status = "0"
	post.DelFlag = "0"

	err := this.db.Gen.SysDept.WithContext(context.Background()).
		Save(post)
	return post, err
}

func (this *SysDeptRepository) EditDept(post *model.SysDept) (*model.SysDept, int64, error) {
	r, err := this.db.Gen.SysDept.WithContext(context.Background()).
		Where(this.db.Gen.SysDept.DeptID.Eq(post.DeptID), this.db.Gen.SysDept.DelFlag.Eq("0")).
		Updates(post)
	return post, r.RowsAffected, err
}

func (this *SysDeptRepository) DeleteDeptById(id int64) (int64, error) {
	r, err := this.db.Gen.SysDept.WithContext(context.Background()).
		Where(this.db.Gen.SysDept.DeptID.Eq(id)).Update(this.db.Gen.SysDept.DelFlag, "2")
	return r.RowsAffected, err
}
