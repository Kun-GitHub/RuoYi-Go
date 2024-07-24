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
)

type SysNoticeRepository struct {
	db *dao.DatabaseStruct
}

func NewSysNoticeRepository(db *dao.DatabaseStruct) *SysNoticeRepository {
	return &SysNoticeRepository{db: db}
}

func (this *SysNoticeRepository) QueryNoticeByID(id int64) (*model.SysNotice, error) {
	structEntity := &model.SysNotice{}

	err := this.db.Gen.SysNotice.WithContext(context.Background()).
		Where(this.db.Gen.SysNotice.NoticeID.Eq(id)).Scan(structEntity)

	if err != nil {
		return nil, err
	}
	return structEntity, nil
}

func (this *SysNoticeRepository) QueryNoticeList(request *model.SysNoticeRequest) ([]*model.SysNotice, error) {
	structEntity := make([]*model.SysNotice, 0)

	var noticeType field.Expr
	var configName field.Expr
	var configKey field.Expr
	if request != nil {
		if request.NoticeType != "" {
			noticeType = this.db.Gen.SysNotice.NoticeType.Eq(request.NoticeType)
		}
		if request.NoticeTitle != "" {
			configName = this.db.Gen.SysNotice.NoticeTitle.Like("%" + request.NoticeTitle + "%")
		}
		if len(request.CreateBy) > 0 {
			configKey = this.db.Gen.SysNotice.CreateBy.Like("%" + request.CreateBy + "%")
		}
	}

	structEntity, err := this.db.Gen.SysNotice.WithContext(context.Background()).
		Where(noticeType, configName, configKey).Find()
	if err != nil {
		return nil, err
	}
	return structEntity, nil
}

func (this *SysNoticeRepository) QueryNoticePage(pageReq common.PageRequest, request *model.SysNoticeRequest) ([]*model.SysNotice, int64, error) {
	structEntity := make([]*model.SysNotice, 0)

	var noticeType field.Expr
	var configName field.Expr
	var configKey field.Expr
	if request != nil {
		if request.NoticeType != "" {
			noticeType = this.db.Gen.SysNotice.NoticeType.Eq(request.NoticeType)
		}
		if request.NoticeTitle != "" {
			configName = this.db.Gen.SysNotice.NoticeTitle.Like("%" + request.NoticeTitle + "%")
		}
		if len(request.CreateBy) > 0 {
			configKey = this.db.Gen.SysNotice.CreateBy.Like("%" + request.CreateBy + "%")
		}
	}

	structEntity, total, err := this.db.Gen.SysNotice.WithContext(context.Background()).
		Where(noticeType, configName, configKey).FindByPage((pageReq.PageNum-1)*pageReq.PageSize, pageReq.PageSize)
	if err != nil {
		return nil, 0, err
	}
	return structEntity, total, err
}
