// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package usecase

import (
	"RuoYi-Go/internal/adapters/dao"
	"gorm.io/gorm"
)

type TransactionManager struct {
	db *dao.DatabaseStruct
}

func NewTransactionManager(db *dao.DatabaseStruct) *TransactionManager {
	return &TransactionManager{db: db}
}

func (tm *TransactionManager) Execute(txFunc func(tx *gorm.DB) error) error {
	return tm.db.Transactional(txFunc)
}
