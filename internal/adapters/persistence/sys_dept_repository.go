// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package persistence

import (
	"RuoYi-Go/internal/adapters/dao"
	"RuoYi-Go/internal/domain/model"
	"context"
	"fmt"
	"gorm.io/gen/field"
	"strings"
	"time"
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
		Where(this.db.Gen.SysDept.DeptID.Eq(id)).
		UpdateSimple(this.db.Gen.SysDept.DelFlag.Value("2"),
			this.db.Gen.SysUser.UpdateBy.Value(this.db.User().UserName),
			this.db.Gen.SysUser.UpdateTime.Value(time.Now()))
	return r.RowsAffected, err
}

func (this *SysDeptRepository) QueryChildIdListById(id int64) ([]int64, error) {
	if id == 0 {
		return []int64{}, nil
	}

	// 查询所有子部门ID（包括直接子部门和间接子部门）
	var childIds []int64

	// 使用原生SQL进行精确匹配，避免ID误匹配问题
	// Ancestors格式为: "0" 或 "0,100" 或 "0,100,101"
	// 需要匹配以下四种情况：
	// 1. 完全匹配: id (当id就是根节点时)
	// 2. 开头匹配: id + ",%"  (如: 100,101,102)
	// 3. 结尾匹配: "%," + id  (如: 0,100)
	// 4. 中间匹配: "%," + id + ",%" (如: 0,100,101)

	var structEntities []*model.SysDept
	var err error

	// 根据不同情况构建不同的查询条件
	// 对于非根节点，使用复杂的OR条件
	structEntities, err = this.db.Gen.SysDept.WithContext(context.Background()).
		Select(this.db.Gen.SysDept.DeptID, this.db.Gen.SysDept.Ancestors).
		Where(this.db.Gen.SysDept.DelFlag.Eq("0")).
		Order(this.db.Gen.SysDept.DeptID).
		Find()

	// 在内存中过滤，确保精确匹配
	filteredEntities := make([]*model.SysDept, 0)
	idStr := fmt.Sprintf("%d", id)
	for _, entity := range structEntities {
		ancestors := entity.Ancestors
		// 精确匹配四种情况
		if ancestors == idStr || // 完全匹配
			strings.HasPrefix(ancestors, idStr+",") || // 开头匹配
			strings.HasSuffix(ancestors, ","+idStr) || // 结尾匹配
			strings.Contains(ancestors, ","+idStr+",") { // 中间匹配
			filteredEntities = append(filteredEntities, entity)
		}
	}
	structEntities = filteredEntities

	if err != nil {
		return nil, err
	}

	// 提取ID列表
	childIds = make([]int64, len(structEntities))
	for i, entity := range structEntities {
		childIds[i] = entity.DeptID
	}

	// 过滤掉自身ID（如果在结果中）
	filteredIds := make([]int64, 0, len(childIds))
	for _, childId := range childIds {
		if childId != id {
			filteredIds = append(filteredIds, childId)
		}
	}

	return filteredIds, nil
}
