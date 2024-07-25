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

type SysJobRepository struct {
	db *dao.DatabaseStruct
}

func NewSysJobRepository(db *dao.DatabaseStruct) *SysJobRepository {
	return &SysJobRepository{db: db}
}

func (this *SysJobRepository) QueryJobByID(id int64) (*model.SysJob, error) {
	structEntity := &model.SysJob{}

	err := this.db.Gen.SysJob.WithContext(context.Background()).
		Where(this.db.Gen.SysJob.JobID.Eq(id)).Scan(structEntity)

	if err != nil {
		return nil, err
	}
	return structEntity, nil
}

func (this *SysJobRepository) QueryJobList(request *model.SysJobRequest) ([]*model.SysJob, error) {
	structEntity := make([]*model.SysJob, 0)

	var jobName field.Expr
	var jobGroup field.Expr
	var status field.Expr
	if request != nil {
		if request.JobName != "" {
			jobName = this.db.Gen.SysJob.JobName.Like("%" + request.JobName + "%")
		}
		if request.JobGroup != "" {
			jobGroup = this.db.Gen.SysJob.JobGroup.Like("%" + request.JobGroup + "%")
		}
		if request.Status != "" {
			status = this.db.Gen.SysJob.Status.Eq(request.Status)
		}
	}

	structEntity, err := this.db.Gen.SysJob.WithContext(context.Background()).
		Where(jobName, jobGroup, status).Find()
	if err != nil {
		return nil, err
	}
	return structEntity, nil
}

func (this *SysJobRepository) QueryJobPage(pageReq common.PageRequest, request *model.SysJobRequest) ([]*model.SysJob, int64, error) {
	structEntity := make([]*model.SysJob, 0)

	var jobName field.Expr
	var jobGroup field.Expr
	var status field.Expr
	if request != nil {
		if request.JobName != "" {
			jobName = this.db.Gen.SysJob.JobName.Like("%" + request.JobName + "%")
		}
		if request.JobGroup != "" {
			jobGroup = this.db.Gen.SysJob.JobGroup.Like("%" + request.JobGroup + "%")
		}
		if request.Status != "" {
			status = this.db.Gen.SysJob.Status.Eq(request.Status)
		}
	}

	structEntity, total, err := this.db.Gen.SysJob.WithContext(context.Background()).
		Where(jobName, jobGroup, status).FindByPage((pageReq.PageNum-1)*pageReq.PageSize, pageReq.PageSize)
	if err != nil {
		return nil, 0, err
	}
	return structEntity, total, err
}

func (this *SysJobRepository) AddJob(post *model.SysJob) (*model.SysJob, error) {
	err := this.db.Gen.SysJob.WithContext(context.Background()).
		Save(post)
	return post, err
}

func (this *SysJobRepository) EditJob(post *model.SysJob) (*model.SysJob, int64, error) {
	r, err := this.db.Gen.SysJob.WithContext(context.Background()).
		Where(this.db.Gen.SysJob.JobID.Eq(post.JobID)).
		Updates(post)
	return post, r.RowsAffected, err
}

func (this *SysJobRepository) DeleteJobById(id int64) (int64, error) {
	r, err := this.db.Gen.SysJob.WithContext(context.Background()).
		Where(this.db.Gen.SysJob.JobID.Eq(id)).Delete()
	return r.RowsAffected, err
}
