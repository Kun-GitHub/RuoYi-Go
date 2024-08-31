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

type SysConfigRepository struct {
	db *dao.DatabaseStruct
}

func NewSysConfigRepository(db *dao.DatabaseStruct) *SysConfigRepository {
	return &SysConfigRepository{db: db}
}

func (this *SysConfigRepository) QueryConfigByID(dictID int64) (*model.SysConfig, error) {
	structEntity := &model.SysConfig{}

	err := this.db.Gen.SysConfig.WithContext(context.Background()).
		Where(this.db.Gen.SysConfig.ConfigID.Eq(dictID)).Scan(structEntity)

	if err != nil {
		return nil, err
	}
	return structEntity, nil
}

func (this *SysConfigRepository) QueryConfigList(request *model.SysConfigRequest) ([]*model.SysConfig, error) {
	structEntity := make([]*model.SysConfig, 0)

	var configType field.Expr
	var configName field.Expr
	var configKey field.Expr
	var timeField field.Expr
	if request != nil {
		if request.ConfigType != "" {
			configType = this.db.Gen.SysConfig.ConfigType.Eq(request.ConfigType)
		}
		if request.ConfigName != "" {
			configName = this.db.Gen.SysConfig.ConfigName.Like("%" + request.ConfigName + "%")
		}
		if request.ConfigKey != "" {
			configKey = this.db.Gen.SysConfig.ConfigKey.Like("%" + request.ConfigKey + "%")
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

				timeField = this.db.Gen.SysConfig.CreateTime.Between(startOfDay, endOfDay)
			}
		}
	}

	structEntity, err := this.db.Gen.SysConfig.WithContext(context.Background()).
		Where(configType, configName, configKey, timeField).Find()
	if err != nil {
		return nil, err
	}
	return structEntity, nil
}

func (this *SysConfigRepository) QueryConfigPage(pageReq common.PageRequest, request *model.SysConfigRequest) ([]*model.SysConfig, int64, error) {
	structEntity := make([]*model.SysConfig, 0)

	var configType field.Expr
	var configName field.Expr
	var configKey field.Expr
	var timeField field.Expr
	if request != nil {
		if request.ConfigType != "" {
			configType = this.db.Gen.SysConfig.ConfigType.Eq(request.ConfigType)
		}
		if request.ConfigName != "" {
			configName = this.db.Gen.SysConfig.ConfigName.Like("%" + request.ConfigName + "%")
		}
		if request.ConfigKey != "" {
			configKey = this.db.Gen.SysConfig.ConfigKey.Like("%" + request.ConfigKey + "%")
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

				timeField = this.db.Gen.SysConfig.CreateTime.Between(startOfDay, endOfDay)
			}
		}
	}

	structEntity, total, err := this.db.Gen.SysConfig.WithContext(context.Background()).
		Where(configType, configName, configKey, timeField).FindByPage((pageReq.PageNum-1)*pageReq.PageSize, pageReq.PageSize)
	if err != nil {
		return nil, 0, err
	}
	return structEntity, total, err
}

func (this *SysConfigRepository) AddConfig(post *model.SysConfig) (*model.SysConfig, error) {
	err := this.db.Gen.SysConfig.WithContext(context.Background()).
		Save(post)
	return post, err
}

func (this *SysConfigRepository) EditConfig(post *model.SysConfig) (*model.SysConfig, int64, error) {
	r, err := this.db.Gen.SysConfig.WithContext(context.Background()).
		Where(this.db.Gen.SysConfig.ConfigID.Eq(post.ConfigID)).
		Updates(post)
	return post, r.RowsAffected, err
}

func (this *SysConfigRepository) DeleteConfigById(id int64) (int64, error) {
	r, err := this.db.Gen.SysConfig.WithContext(context.Background()).
		Where(this.db.Gen.SysConfig.ConfigID.Eq(id)).Delete()
	return r.RowsAffected, err
}

func (this *SysConfigRepository) CheckConfigNameUnique(id int64, name string) (int64, error) {
	r, err := this.db.Gen.SysConfig.WithContext(context.Background()).
		Where(this.db.Gen.SysConfig.ConfigName.Eq(name), this.db.Gen.SysConfig.ConfigID.Neq(id)).Count()
	return r, err
}

func (this *SysConfigRepository) QueryConfigByKey(configKey string) (*model.SysConfig, error) {
	structEntity := &model.SysConfig{}

	err := this.db.Gen.SysConfig.WithContext(context.Background()).
		Where(this.db.Gen.SysConfig.ConfigKey.Eq(configKey)).
		Limit(1).Scan(structEntity)

	if err != nil {
		return nil, err
	}
	return structEntity, nil
}
