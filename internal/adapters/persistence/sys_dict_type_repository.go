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
	"gorm.io/gen/field"
	"time"
)

type SysDictTypeRepository struct {
	db *dao.DatabaseStruct
}

func NewSysDictTypeRepository(db *dao.DatabaseStruct) *SysDictTypeRepository {
	return &SysDictTypeRepository{db: db}
}

func (this *SysDictTypeRepository) QueryDictTypeByDictID(dictID int64) (*model.SysDictType, error) {
	structEntity := &model.SysDictType{}

	err := this.db.Gen.SysDictType.WithContext(context.Background()).
		Where(this.db.Gen.SysDictType.DictID.Eq(dictID)).Scan(structEntity)

	if err != nil {
		return nil, err
	}
	return structEntity, nil
}

func (this *SysDictTypeRepository) QueryDictTypeList(request *model.SysDictTypeRequest) ([]*model.SysDictType, error) {
	structEntity := make([]*model.SysDictType, 0)

	var status field.Expr
	var dictName field.Expr
	var dictType field.Expr
	var timeField field.Expr
	if request != nil {
		if request.Status != "" {
			status = this.db.Gen.SysDictType.Status.Eq(request.Status)
		}
		if request.DictName != "" {
			dictName = this.db.Gen.SysDictType.DictName.Like("%" + request.DictName + "%")
		}
		if request.DictType != "" {
			dictType = this.db.Gen.SysDictType.DictType.Like("%" + request.DictType + "%")
		}
		if request.BeginTime != "" && request.EndTime != "" {
			// 解析日期字符串
			t1, err1 := time.Parse("2006-01-02", request.BeginTime)
			t2, err2 := time.Parse("2006-01-02", request.EndTime)
			if err1 == nil && err2 == nil {
				// 设置一天的开始时间
				startOfDay := time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, t1.Location())
				// 设置一天的开始时间
				endOfDay := time.Date(t2.Year(), t2.Month(), t2.Day(), 23, 59, 59, 0, t2.Location())

				timeField = this.db.Gen.SysDictType.CreateTime.Between(startOfDay, endOfDay)
			}
		}
	}

	structEntity, err := this.db.Gen.SysDictType.WithContext(context.Background()).
		Where(status, dictName, dictType, timeField).Find()
	if err != nil {
		return nil, err
	}
	return structEntity, nil
}

func (this *SysDictTypeRepository) QueryDictTypePage(pageReq common.PageRequest, request *model.SysDictTypeRequest) ([]*model.SysDictType, int64, error) {
	structEntity := make([]*model.SysDictType, 0)

	var status field.Expr
	var dictName field.Expr
	var dictType field.Expr
	var timeField field.Expr
	if request != nil {
		if request.Status != "" {
			status = this.db.Gen.SysDictType.Status.Eq(request.Status)
		}
		if request.DictName != "" {
			dictName = this.db.Gen.SysDictType.DictName.Like("%" + request.DictName + "%")
		}
		if request.DictType != "" {
			dictType = this.db.Gen.SysDictType.DictType.Like("%" + request.DictType + "%")
		}
		if request.BeginTime != "" && request.EndTime != "" {
			// 解析日期字符串
			t1, err1 := time.Parse("2006-01-02", request.BeginTime)
			t2, err2 := time.Parse("2006-01-02", request.EndTime)
			if err1 == nil && err2 == nil {
				// 设置一天的开始时间
				startOfDay := time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, t1.Location())
				// 设置一天的开始时间
				endOfDay := time.Date(t2.Year(), t2.Month(), t2.Day(), 23, 59, 59, 0, t2.Location())

				timeField = this.db.Gen.SysDictType.CreateTime.Between(startOfDay, endOfDay)
			}
		}
	}

	structEntity, total, err := this.db.Gen.SysDictType.WithContext(context.Background()).
		Where(status, dictName, dictType, timeField).FindByPage((pageReq.PageNum-1)*pageReq.PageSize, pageReq.PageSize)
	if err != nil {
		return nil, 0, err
	}
	return structEntity, total, err
}

func (this *SysDictTypeRepository) AddDictType(post *model.SysDictType) (*model.SysDictType, error) {
	err := this.db.Gen.SysDictType.WithContext(context.Background()).
		Save(post)
	return post, err
}

func (this *SysDictTypeRepository) EditDictType(post *model.SysDictType) (*model.SysDictType, int64, error) {
	r, err := this.db.Gen.SysDictType.WithContext(context.Background()).
		Where(this.db.Gen.SysDictType.DictID.Eq(post.DictID)).
		Updates(post)
	return post, r.RowsAffected, err
}

func (this *SysDictTypeRepository) DeleteDictTypeById(id int64) (int64, error) {
	r, err := this.db.Gen.SysDictType.WithContext(context.Background()).
		Where(this.db.Gen.SysDictType.DictID.Eq(id)).Delete()
	return r.RowsAffected, err
}

func (this *SysDictTypeRepository) CheckDictTypeUnique(id int64, name string) (int64, error) {
	r, err := this.db.Gen.SysDictType.WithContext(context.Background()).
		Where(this.db.Gen.SysDictType.DictType.Eq(name), this.db.Gen.SysDictType.DictID.Neq(id)).Count()
	return r, err
}
