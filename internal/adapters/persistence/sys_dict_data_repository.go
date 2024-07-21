// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
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
	structEntity := make([]*model.SysDictDatum, 0)

	structEntity, err := this.db.Gen.SysDictDatum.WithContext(context.Background()).
		Where(this.db.Gen.SysDictDatum.DictType.Eq(typeStr),
			this.db.Gen.SysDictDatum.Status.Eq("0")).Find()
	if err != nil {
		return nil, err
	}
	return structEntity, nil
}
