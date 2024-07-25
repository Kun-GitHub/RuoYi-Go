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

type SysLogininforRepository struct {
	db *dao.DatabaseStruct
}

func NewSysLogininforRepository(db *dao.DatabaseStruct) *SysLogininforRepository {
	return &SysLogininforRepository{db: db}
}

func (this *SysLogininforRepository) QueryLogininforByID(id int64) (*model.SysLogininfor, error) {
	structEntity := &model.SysLogininfor{}

	err := this.db.Gen.SysLogininfor.WithContext(context.Background()).
		Where(this.db.Gen.SysLogininfor.InfoID.Eq(id)).Scan(structEntity)

	if err != nil {
		return nil, err
	}
	return structEntity, nil
}

func (this *SysLogininforRepository) QueryLogininforList(request *model.SysLogininforRequest) ([]*model.SysLogininfor, error) {
	structEntity := make([]*model.SysLogininfor, 0)

	var status field.Expr
	var ipaddr field.Expr
	var userName field.Expr
	var timeField field.Expr
	if request != nil {
		if request.Status != "" {
			status = this.db.Gen.SysLogininfor.Status.Eq(request.Status)
		}
		if request.Ipaddr != "" {
			ipaddr = this.db.Gen.SysLogininfor.Ipaddr.Like("%" + request.Ipaddr + "%")
		}
		if len(request.UserName) > 0 {
			userName = this.db.Gen.SysLogininfor.UserName.Like("%" + request.UserName + "%")
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

				timeField = this.db.Gen.SysLogininfor.LoginTime.Between(startOfDay, endOfDay)
			}
		}
	}

	structEntity, err := this.db.Gen.SysLogininfor.WithContext(context.Background()).
		Where(status, ipaddr, userName, timeField).Find()
	if err != nil {
		return nil, err
	}
	return structEntity, nil
}

func (this *SysLogininforRepository) QueryLogininforPage(pageReq common.PageRequest, request *model.SysLogininforRequest) ([]*model.SysLogininfor, int64, error) {
	structEntity := make([]*model.SysLogininfor, 0)

	var status field.Expr
	var ipaddr field.Expr
	var userName field.Expr
	var timeField field.Expr
	if request != nil {
		if request.Status != "" {
			status = this.db.Gen.SysLogininfor.Status.Eq(request.Status)
		}
		if request.Ipaddr != "" {
			ipaddr = this.db.Gen.SysLogininfor.Ipaddr.Like("%" + request.Ipaddr + "%")
		}
		if len(request.UserName) > 0 {
			userName = this.db.Gen.SysLogininfor.UserName.Like("%" + request.UserName + "%")
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

				timeField = this.db.Gen.SysLogininfor.LoginTime.Between(startOfDay, endOfDay)
			}
		}
	}

	structEntity, total, err := this.db.Gen.SysLogininfor.WithContext(context.Background()).
		Where(status, ipaddr, userName, timeField).FindByPage((pageReq.PageNum-1)*pageReq.PageSize, pageReq.PageSize)
	if err != nil {
		return nil, 0, err
	}
	return structEntity, total, err
}

func (this *SysLogininforRepository) AddLogininfor(post *model.SysLogininfor) (*model.SysLogininfor, error) {
	err := this.db.Gen.SysLogininfor.WithContext(context.Background()).
		Save(post)
	return post, err
}

func (this *SysLogininforRepository) EditLogininfor(post *model.SysLogininfor) (*model.SysLogininfor, int64, error) {
	r, err := this.db.Gen.SysLogininfor.WithContext(context.Background()).
		Where(this.db.Gen.SysLogininfor.InfoID.Eq(post.InfoID)).
		Updates(post)
	return post, r.RowsAffected, err
}

func (this *SysLogininforRepository) DeleteLogininforById(id int64) (int64, error) {
	r, err := this.db.Gen.SysLogininfor.WithContext(context.Background()).
		Where(this.db.Gen.SysLogininfor.InfoID.Eq(id)).Delete()
	return r.RowsAffected, err
}
