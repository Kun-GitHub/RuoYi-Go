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

type SysPostRepository struct {
	db *dao.DatabaseStruct
}

func NewSysPostRepository(db *dao.DatabaseStruct) *SysPostRepository {
	return &SysPostRepository{db: db}
}

func (this *SysPostRepository) QueryPostByPostId(postId int64) (*model.SysPost, error) {
	structEntity := &model.SysPost{}

	err := this.db.Gen.SysPost.WithContext(context.Background()).
		Where(this.db.Gen.SysPost.PostID.Eq(postId)).Scan(structEntity)

	if err != nil {
		return nil, err
	}
	return structEntity, nil
}

func (this *SysPostRepository) QueryPostByUserId(userId int64) ([]*model.SysPost, error) {
	structEntity := make([]*model.SysPost, 0)

	structEntity, err := this.db.Gen.SysPost.WithContext(context.Background()).Select(this.db.Gen.SysPost.ALL).
		LeftJoin(this.db.Gen.SysUserPost, this.db.Gen.SysUserPost.PostID.EqCol(this.db.Gen.SysPost.PostID)).
		Where(this.db.Gen.SysUserPost.UserID.Eq(userId)).Find()

	if err != nil {
		return nil, err
	}
	return structEntity, nil
}

func (this *SysPostRepository) QueryPostList(post *model.SysPost) ([]*model.SysPost, error) {
	structEntity := make([]*model.SysPost, 0)

	var status field.Expr
	var name field.Expr
	if post != nil {
		if post.Status != "" {
			status = this.db.Gen.SysPost.Status.Eq(post.Status)
		}

		if post.PostName != "" {
			name = this.db.Gen.SysPost.PostName.Like("%" + post.PostName + "%")
		}
	}

	structEntity, err := this.db.Gen.SysPost.WithContext(context.Background()).
		Where(name, status).Find()
	if err != nil {
		return nil, err
	}
	return structEntity, nil
}

func (this *SysPostRepository) QueryPostPage(pageReq common.PageRequest, request *model.SysPostRequest) ([]*model.SysPost, int64, error) {
	structEntity := make([]*model.SysPost, 0)

	var status field.Expr
	var postName field.Expr
	var postCode field.Expr
	if request != nil {
		if request.Status != "" {
			status = this.db.Gen.SysPost.Status.Eq(request.Status)
		}
		if request.PostName != "" {
			postName = this.db.Gen.SysPost.PostName.Like("%" + request.PostName + "%")
		}
		if request.PostCode != "" {
			postCode = this.db.Gen.SysPost.PostCode.Like("%" + request.PostCode + "%")
		}
	}

	structEntity, err := this.db.Gen.SysPost.WithContext(context.Background()).
		Where(status, postName, postCode).Limit(pageReq.PageSize).Offset((pageReq.PageNum - 1) * pageReq.PageSize).Find()
	total, err := this.db.Gen.SysPost.WithContext(context.Background()).
		Where(status, postName, postCode).Limit(pageReq.PageSize).Offset((pageReq.PageNum - 1) * pageReq.PageSize).Count()

	if err != nil {
		return nil, 0, err
	}
	return structEntity, total, err
}
