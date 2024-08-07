// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:            db,
		SysConfig:     newSysConfig(db, opts...),
		SysDept:       newSysDept(db, opts...),
		SysDictDatum:  newSysDictDatum(db, opts...),
		SysDictType:   newSysDictType(db, opts...),
		SysJob:        newSysJob(db, opts...),
		SysJobLog:     newSysJobLog(db, opts...),
		SysLogininfor: newSysLogininfor(db, opts...),
		SysMenu:       newSysMenu(db, opts...),
		SysNotice:     newSysNotice(db, opts...),
		SysOperLog:    newSysOperLog(db, opts...),
		SysPost:       newSysPost(db, opts...),
		SysRole:       newSysRole(db, opts...),
		SysRoleDept:   newSysRoleDept(db, opts...),
		SysRoleMenu:   newSysRoleMenu(db, opts...),
		SysUser:       newSysUser(db, opts...),
		SysUserPost:   newSysUserPost(db, opts...),
		SysUserRole:   newSysUserRole(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	SysConfig     sysConfig
	SysDept       sysDept
	SysDictDatum  sysDictDatum
	SysDictType   sysDictType
	SysJob        sysJob
	SysJobLog     sysJobLog
	SysLogininfor sysLogininfor
	SysMenu       sysMenu
	SysNotice     sysNotice
	SysOperLog    sysOperLog
	SysPost       sysPost
	SysRole       sysRole
	SysRoleDept   sysRoleDept
	SysRoleMenu   sysRoleMenu
	SysUser       sysUser
	SysUserPost   sysUserPost
	SysUserRole   sysUserRole
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:            db,
		SysConfig:     q.SysConfig.clone(db),
		SysDept:       q.SysDept.clone(db),
		SysDictDatum:  q.SysDictDatum.clone(db),
		SysDictType:   q.SysDictType.clone(db),
		SysJob:        q.SysJob.clone(db),
		SysJobLog:     q.SysJobLog.clone(db),
		SysLogininfor: q.SysLogininfor.clone(db),
		SysMenu:       q.SysMenu.clone(db),
		SysNotice:     q.SysNotice.clone(db),
		SysOperLog:    q.SysOperLog.clone(db),
		SysPost:       q.SysPost.clone(db),
		SysRole:       q.SysRole.clone(db),
		SysRoleDept:   q.SysRoleDept.clone(db),
		SysRoleMenu:   q.SysRoleMenu.clone(db),
		SysUser:       q.SysUser.clone(db),
		SysUserPost:   q.SysUserPost.clone(db),
		SysUserRole:   q.SysUserRole.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:            db,
		SysConfig:     q.SysConfig.replaceDB(db),
		SysDept:       q.SysDept.replaceDB(db),
		SysDictDatum:  q.SysDictDatum.replaceDB(db),
		SysDictType:   q.SysDictType.replaceDB(db),
		SysJob:        q.SysJob.replaceDB(db),
		SysJobLog:     q.SysJobLog.replaceDB(db),
		SysLogininfor: q.SysLogininfor.replaceDB(db),
		SysMenu:       q.SysMenu.replaceDB(db),
		SysNotice:     q.SysNotice.replaceDB(db),
		SysOperLog:    q.SysOperLog.replaceDB(db),
		SysPost:       q.SysPost.replaceDB(db),
		SysRole:       q.SysRole.replaceDB(db),
		SysRoleDept:   q.SysRoleDept.replaceDB(db),
		SysRoleMenu:   q.SysRoleMenu.replaceDB(db),
		SysUser:       q.SysUser.replaceDB(db),
		SysUserPost:   q.SysUserPost.replaceDB(db),
		SysUserRole:   q.SysUserRole.replaceDB(db),
	}
}

type queryCtx struct {
	SysConfig     *sysConfigDo
	SysDept       *sysDeptDo
	SysDictDatum  *sysDictDatumDo
	SysDictType   *sysDictTypeDo
	SysJob        *sysJobDo
	SysJobLog     *sysJobLogDo
	SysLogininfor *sysLogininforDo
	SysMenu       *sysMenuDo
	SysNotice     *sysNoticeDo
	SysOperLog    *sysOperLogDo
	SysPost       *sysPostDo
	SysRole       *sysRoleDo
	SysRoleDept   *sysRoleDeptDo
	SysRoleMenu   *sysRoleMenuDo
	SysUser       *sysUserDo
	SysUserPost   *sysUserPostDo
	SysUserRole   *sysUserRoleDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		SysConfig:     q.SysConfig.WithContext(ctx),
		SysDept:       q.SysDept.WithContext(ctx),
		SysDictDatum:  q.SysDictDatum.WithContext(ctx),
		SysDictType:   q.SysDictType.WithContext(ctx),
		SysJob:        q.SysJob.WithContext(ctx),
		SysJobLog:     q.SysJobLog.WithContext(ctx),
		SysLogininfor: q.SysLogininfor.WithContext(ctx),
		SysMenu:       q.SysMenu.WithContext(ctx),
		SysNotice:     q.SysNotice.WithContext(ctx),
		SysOperLog:    q.SysOperLog.WithContext(ctx),
		SysPost:       q.SysPost.WithContext(ctx),
		SysRole:       q.SysRole.WithContext(ctx),
		SysRoleDept:   q.SysRoleDept.WithContext(ctx),
		SysRoleMenu:   q.SysRoleMenu.WithContext(ctx),
		SysUser:       q.SysUser.WithContext(ctx),
		SysUserPost:   q.SysUserPost.WithContext(ctx),
		SysUserRole:   q.SysUserRole.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
