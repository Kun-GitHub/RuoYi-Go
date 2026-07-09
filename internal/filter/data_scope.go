// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package filter

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/domain/model"

	"github.com/kataras/iris/v12"
)

const (
	DataScopeAll          = "1"
	DataScopeCustom       = "2"
	DataScopeDept         = "3"
	DataScopeDeptAndChild = "4"
	DataScopeSelf         = "5"
)

// DeptRepo 数据权限需要的部门仓储接口
type DeptRepo interface {
	QueryChildIdListById(deptId int64) ([]int64, error)
}

// RoleService 数据权限需要的角色服务接口
type RoleService interface {
	QueryRolesByUserId(userId int64) ([]*model.SysRole, error)
}

// DataScopeResult 数据权限计算结果
type DataScopeResult struct {
	ScopeType string  // 1-5
	DeptIds   []int64 // 过滤的部门ID
	UserId    int64   // 仅本人的用户ID
	IsAdmin   bool
}

// ComputeDataScope 计算当前用户的数据权限
func ComputeDataScope(ctx iris.Context, roleSvc RoleService, deptRepo DeptRepo) *DataScopeResult {
	loginUser, _ := ctx.Values().Get(common.LOGINUSER).(*model.UserInfoStruct)
	if loginUser == nil {
		return &DataScopeResult{ScopeType: DataScopeAll}
	}
	if loginUser.Admin {
		return &DataScopeResult{ScopeType: DataScopeAll, IsAdmin: true}
	}

	roles, err := roleSvc.QueryRolesByUserId(loginUser.UserID)
	if err != nil || len(roles) == 0 {
		return &DataScopeResult{ScopeType: DataScopeAll}
	}

	result := &DataScopeResult{}
	for _, role := range roles {
		if role.DataScope == "" {
			continue
		}
		switch role.DataScope {
		case DataScopeAll:
			return &DataScopeResult{ScopeType: DataScopeAll}
		case DataScopeCustom:
			result.ScopeType = DataScopeCustom
			result.DeptIds = append(result.DeptIds, role.DeptIds...)
		case DataScopeDept:
			result.ScopeType = DataScopeDept
			result.DeptIds = append(result.DeptIds, loginUser.DeptID)
		case DataScopeDeptAndChild:
			children, _ := deptRepo.QueryChildIdListById(loginUser.DeptID)
			result.ScopeType = DataScopeDeptAndChild
			result.DeptIds = append(result.DeptIds, loginUser.DeptID)
			result.DeptIds = append(result.DeptIds, children...)
		case DataScopeSelf:
			result.ScopeType = DataScopeSelf
			result.UserId = loginUser.UserID
		}
	}
	return result
}

// GetDataScopeDeptIds 获取数据权限过滤的部门ID列表
// 返回空数组表示不过滤，返回nil表示全部数据权限
// 返回[-1]表示仅本人数据
func GetDataScopeDeptIds(ctx iris.Context, roleSvc RoleService, deptRepo DeptRepo) []int64 {
	result := ComputeDataScope(ctx, roleSvc, deptRepo)
	if result.IsAdmin || result.ScopeType == DataScopeAll {
		return nil
	}
	if result.ScopeType == DataScopeSelf {
		return []int64{-1}
	}
	return result.DeptIds
}
