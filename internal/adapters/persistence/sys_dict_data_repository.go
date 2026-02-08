// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package persistence

import (
	"RuoYi-Go/internal/adapters/dao"
	"RuoYi-Go/internal/domain/model"
	"context"
)

type SysDictDataRepository struct {
	db *dao.DatabaseStruct
}

func NewSysDictDataRepository(db *dao.DatabaseStruct) *SysDictDataRepository {
	return &SysDictDataRepository{db: db}
}

func (this *SysDictDataRepository) QueryDictDatasByType(typeStr string) ([]*model.SysDictDatum, error) {
	structEntity, err := this.db.Gen.SysDictDatum.WithContext(context.Background()).
		Where(this.db.Gen.SysDictDatum.DictType.Eq(typeStr),
			this.db.Gen.SysDictDatum.Status.Eq("0")).
		Order(this.db.Gen.SysDictDatum.DictSort).
		Find()
	if err != nil {
		return nil, err
	}
	return structEntity, nil
}

func (this *SysDictDataRepository) Get(id uint) (*model.SysDictDatum, error) {
	return this.db.Gen.SysDictDatum.WithContext(context.Background()).Where(this.db.Gen.SysDictDatum.DictCode.Eq(int64(id))).First()
}

func (this *SysDictDataRepository) List(page, size int, dictLabel, dictType, status string) ([]*model.SysDictDatum, int64, error) {
	q := this.db.Gen.SysDictDatum.WithContext(context.Background())
	if dictLabel != "" {
		q = q.Where(this.db.Gen.SysDictDatum.DictLabel.Like("%" + dictLabel + "%"))
	}
	if dictType != "" {
		q = q.Where(this.db.Gen.SysDictDatum.DictType.Eq(dictType))
	}
	if status != "" {
		q = q.Where(this.db.Gen.SysDictDatum.Status.Eq(status))
	}
	return q.Order(this.db.Gen.SysDictDatum.DictSort).FindByPage((page-1)*size, size)
}

func (this *SysDictDataRepository) Create(data *model.SysDictDatum) error {
	return this.db.Gen.SysDictDatum.WithContext(context.Background()).Create(data)
}

func (this *SysDictDataRepository) Update(data *model.SysDictDatum) error {
	_, err := this.db.Gen.SysDictDatum.WithContext(context.Background()).Where(this.db.Gen.SysDictDatum.DictCode.Eq(data.DictCode)).Updates(data)
	return err
}

func (this *SysDictDataRepository) Delete(ids []int64) error {
	_, err := this.db.Gen.SysDictDatum.WithContext(context.Background()).Where(this.db.Gen.SysDictDatum.DictCode.In(ids...)).Delete()
	return err
}
