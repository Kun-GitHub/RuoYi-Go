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
	"gorm.io/gorm"
)

type SysMenuRepository struct {
	db *dao.DatabaseStruct
}

func NewSysMenuRepository(db *dao.DatabaseStruct) *SysMenuRepository {
	return &SysMenuRepository{db: db}
}

func (this *SysMenuRepository) QueryMenusByUserId(userId int64) ([]*model.SysMenu, error) {
	var menus []*model.SysMenu
	err := this.db.Transactional(func(db *gorm.DB) error {
		if userId == common.ADMINID {
			err := db.Table("sys_menu AS sm").Select("sm.menu_id, sm.parent_id, sm.menu_name, sm.path, sm.component, sm.query, sm.visible, sm.status, sm.perms, sm.is_frame, sm.is_cache, sm.menu_type, sm.icon, sm.order_num, sm.create_time").
				Where("sm.menu_type IN (?, ?) AND sm.status = '0'", "M", "C").
				Order("sm.parent_id, sm.order_num").
				Find(&menus).Error
			return err
		} else {
			err := db.Table("sys_menu AS sm").Select("sm.menu_id, sm.parent_id, sm.menu_name, sm.path, sm.component, sm.query, sm.visible, sm.status, sm.perms, sm.is_frame, sm.is_cache, sm.menu_type, sm.icon, sm.order_num, sm.create_time").
				Joins("LEFT JOIN sys_role_menu srm on srm.menu_id = sm.menu_id").
				Joins("LEFT JOIN sys_user_role sur on sur.role_id = srm.role_id").
				Where("sm.menu_type IN (?, ?) AND sm.status = '0' and sur.user_id = ? ", "M", "C", userId).
				Order("sm.parent_id, sm.order_num").
				Find(&menus).Error
			return err
		}
	})
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (this *SysMenuRepository) QueryMenuByID(id int64) (*model.SysMenu, error) {
	structEntity := &model.SysMenu{}

	err := this.db.Gen.SysMenu.WithContext(context.Background()).
		Where(this.db.Gen.SysMenu.MenuID.Eq(id)).Scan(structEntity)

	if err != nil {
		return nil, err
	}
	return structEntity, nil
}

func (this *SysMenuRepository) QueryMenuList(request *model.SysMenuRequest) ([]*model.SysMenu, error) {
	structEntity := make([]*model.SysMenu, 0)

	var configName field.Expr
	var status field.Expr
	if request != nil {
		if request.MenuName != "" {
			configName = this.db.Gen.SysMenu.MenuName.Like("%" + request.MenuName + "%")
		}
		if request.Status != "" {
			status = this.db.Gen.SysMenu.Status.Eq(request.Status)
		}
	}

	structEntity, err := this.db.Gen.SysMenu.WithContext(context.Background()).
		Where(configName, status).Find()
	if err != nil {
		return nil, err
	}
	return structEntity, nil
}

func (this *SysMenuRepository) QueryMenuPage(pageReq common.PageRequest, request *model.SysMenuRequest) ([]*model.SysMenu, int64, error) {
	structEntity := make([]*model.SysMenu, 0)

	var configName field.Expr
	var status field.Expr
	if request != nil {
		if request.MenuName != "" {
			configName = this.db.Gen.SysMenu.MenuName.Like("%" + request.MenuName + "%")
		}
		if request.Status != "" {
			status = this.db.Gen.SysMenu.Status.Eq(request.Status)
		}
	}

	structEntity, total, err := this.db.Gen.SysMenu.WithContext(context.Background()).
		Where(configName, status).FindByPage((pageReq.PageNum-1)*pageReq.PageSize, pageReq.PageSize)
	if err != nil {
		return nil, 0, err
	}
	return structEntity, total, err
}

func (this *SysMenuRepository) AddMenu(post *model.SysMenu) (*model.SysMenu, error) {
	err := this.db.Gen.SysMenu.WithContext(context.Background()).
		Save(post)
	return post, err
}

func (this *SysMenuRepository) EditMenu(post *model.SysMenu) (*model.SysMenu, int64, error) {
	r, err := this.db.Gen.SysMenu.WithContext(context.Background()).
		Where(this.db.Gen.SysMenu.MenuID.Eq(post.MenuID)).
		Updates(post)
	return post, r.RowsAffected, err
}

func (this *SysMenuRepository) DeleteMenuById(id int64) (int64, error) {
	r, err := this.db.Gen.SysMenu.WithContext(context.Background()).
		Where(this.db.Gen.SysMenu.MenuID.Eq(id)).Delete()
	return r.RowsAffected, err
}
