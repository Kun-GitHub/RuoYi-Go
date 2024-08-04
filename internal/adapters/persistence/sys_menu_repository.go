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

type SysMenuRepository struct {
	db *dao.DatabaseStruct
}

func NewSysMenuRepository(db *dao.DatabaseStruct) *SysMenuRepository {
	return &SysMenuRepository{db: db}
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
	var err error

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

	if request.UserId == common.ADMINID {
		structEntity, err = this.db.Gen.SysMenu.WithContext(context.Background()).Distinct(this.db.Gen.SysMenu.ALL).
			Where(configName, status, this.db.Gen.SysMenu.MenuType.In("M", "C")).
			Order(this.db.Gen.SysMenu.ParentID, this.db.Gen.SysMenu.OrderNum).Order(this.db.Gen.SysMenu.ParentID, this.db.Gen.SysMenu.OrderNum).Find()
	} else {
		structEntity, err = this.db.Gen.SysMenu.WithContext(context.Background()).Distinct(this.db.Gen.SysMenu.ALL).
			LeftJoin(this.db.Gen.SysRoleMenu, this.db.Gen.SysRoleMenu.MenuID.EqCol(this.db.Gen.SysMenu.MenuID)).
			LeftJoin(this.db.Gen.SysUserRole, this.db.Gen.SysUserRole.RoleID.EqCol(this.db.Gen.SysRoleMenu.RoleID)).
			Where(configName, status, this.db.Gen.SysMenu.MenuType.In("M", "C"),
				this.db.Gen.SysUserRole.UserID.Eq(request.UserId)).Order(this.db.Gen.SysMenu.ParentID, this.db.Gen.SysMenu.OrderNum).Find()
	}

	if err != nil {
		return nil, err
	}
	return structEntity, nil
}

func (this *SysMenuRepository) QueryMenuPage(pageReq common.PageRequest, request *model.SysMenuRequest) ([]*model.SysMenu, int64, error) {
	structEntity := make([]*model.SysMenu, 0)
	var err error
	var total int64

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

	if request.UserId == common.ADMINID {
		structEntity, total, err = this.db.Gen.SysMenu.WithContext(context.Background()).Distinct(this.db.Gen.SysMenu.ALL).
			Where(configName, status, this.db.Gen.SysMenu.MenuType.In("M", "C")).
			Order(this.db.Gen.SysMenu.ParentID, this.db.Gen.SysMenu.OrderNum).Order(this.db.Gen.SysMenu.ParentID, this.db.Gen.SysMenu.OrderNum).
			FindByPage((pageReq.PageNum-1)*pageReq.PageSize, pageReq.PageSize)
	} else {
		structEntity, total, err = this.db.Gen.SysMenu.WithContext(context.Background()).Distinct(this.db.Gen.SysMenu.ALL).
			LeftJoin(this.db.Gen.SysRoleMenu, this.db.Gen.SysRoleMenu.MenuID.EqCol(this.db.Gen.SysMenu.MenuID)).
			LeftJoin(this.db.Gen.SysUserRole, this.db.Gen.SysUserRole.RoleID.EqCol(this.db.Gen.SysRoleMenu.RoleID)).
			Where(configName, status, this.db.Gen.SysMenu.MenuType.In("M", "C"),
				this.db.Gen.SysUserRole.UserID.Eq(request.UserId)).Order(this.db.Gen.SysMenu.ParentID, this.db.Gen.SysMenu.OrderNum).
			FindByPage((pageReq.PageNum-1)*pageReq.PageSize, pageReq.PageSize)
	}

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
