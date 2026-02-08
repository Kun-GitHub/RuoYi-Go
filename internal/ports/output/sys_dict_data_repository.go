// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package output

import "RuoYi-Go/internal/domain/model"

type SysDictDataRepository interface {
	QueryDictDatasByType(typeStr string) ([]*model.SysDictDatum, error)
	Get(id uint) (*model.SysDictDatum, error)
	List(page, size int, dictLabel, dictType, status string) ([]*model.SysDictDatum, int64, error)
	Create(data *model.SysDictDatum) error
	Update(data *model.SysDictDatum) error
	Delete(ids []int64) error
}
