// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package persistence

import (
	"RuoYi-Go/internal/domain/model"
	"gorm.io/gorm"
)

type DemoRepository struct {
	db *gorm.DB
}

func NewDemoRepository() *DemoRepository {
	return &DemoRepository{}
}

func (r *DemoRepository) GetDemoByID(id uint) (*model.Demo, error) {
	var demo model.Demo
	if err := r.db.First(&demo, id).Error; err != nil {
		return nil, err
	}
	return &demo, nil
}
