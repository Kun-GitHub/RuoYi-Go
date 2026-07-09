package service

import (
	"RuoYi-Go/internal/domain/model"
)

const (
	DataScopeAll          = "1"
	DataScopeCustom       = "2"
	DataScopeDept         = "3"
	DataScopeDeptAndChild = "4"
	DataScopeSelf         = "5"
)

type DataScopeResult struct {
	ScopeType string  // 1-5
	DeptIds   []int64 // applicable dept IDs
	UserId    int64   // for self scope
	IsAdmin   bool    // admin bypass
}

type DeptRepository interface {
	QueryChildIdListById(deptId int64) ([]int64, error)
}

type DataScopeService struct {
	deptRepo DeptRepository
}

func NewDataScopeService(deptRepo DeptRepository) *DataScopeService {
	return &DataScopeService{deptRepo: deptRepo}
}

func (s *DataScopeService) GetDataScope(userId int64, deptId int64, roles []*model.SysRole) *DataScopeResult {
	result := &DataScopeResult{ScopeType: DataScopeAll}

	if len(roles) == 0 {
		return result
	}

	for _, role := range roles {
		if role.DataScope == "" {
			continue
		}
		switch role.DataScope {
		case DataScopeAll:
			return &DataScopeResult{ScopeType: DataScopeAll, IsAdmin: false}
		case DataScopeCustom:
			if len(role.DeptIds) > 0 {
				result.ScopeType = DataScopeCustom
				result.DeptIds = append(result.DeptIds, role.DeptIds...)
			}
		case DataScopeDept:
			result.ScopeType = DataScopeDept
			result.DeptIds = append(result.DeptIds, deptId)
		case DataScopeDeptAndChild:
			children, _ := s.deptRepo.QueryChildIdListById(deptId)
			result.ScopeType = DataScopeDeptAndChild
			result.DeptIds = append(result.DeptIds, deptId)
			result.DeptIds = append(result.DeptIds, children...)
		case DataScopeSelf:
			result.ScopeType = DataScopeSelf
			result.UserId = userId
		}
	}
	return result
}
