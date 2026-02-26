// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package usecase

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/domain/model"
	"RuoYi-Go/internal/ports/input"
	"RuoYi-Go/internal/ports/output"
	"RuoYi-Go/pkg/cache"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
)

// SysUserService 用户服务实现类
// 负责处理用户相关的业务逻辑，包括用户查询、增删改查、状态管理等
type SysUserService struct {
	repo     output.SysUserRepository
	roleRepo output.SysRoleRepository
	deptRepo output.SysDeptRepository
	cache    *cache.FreeCacheClient
	logger   *zap.Logger
}

// NewPageSysUserService 创建分页用户服务实例
// 用于需要分页查询用户信息的场景
//
// 参数:
//
//	repo: 用户仓储接口
//	roleRepo: 角色仓储接口
//	deptRepo: 部门仓储接口
//	cache: 缓存客户端
//	logger: 日志记录器
//
// 返回值:
//
//	用户服务接口
func NewPageSysUserService(repo output.SysUserRepository, roleRepo output.SysRoleRepository, deptRepo output.SysDeptRepository, cache *cache.FreeCacheClient, logger *zap.Logger) input.SysUserService {
	return &SysUserService{repo: repo, roleRepo: roleRepo, deptRepo: deptRepo, cache: cache, logger: logger}
}

// NewSysUserService 创建基础用户服务实例
// 用于基本的用户操作场景
//
// 参数:
//
//	repo: 用户仓储接口
//	cache: 缓存客户端
//	logger: 日志记录器
//
// 返回值:
//
//	用户服务接口
func NewSysUserService(repo output.SysUserRepository, cache *cache.FreeCacheClient, logger *zap.Logger) input.SysUserService {
	return &SysUserService{repo: repo, cache: cache, logger: logger}
}

// QueryUserByUserName 根据用户名查询用户信息
// 支持缓存机制，优先从缓存获取，缓存未命中则查询数据库并更新缓存
//
// 参数:
//
//	username: 用户名
//
// 返回值:
//
//	*model.SysUser: 用户信息
//	error: 错误信息
func (this *SysUserService) QueryUserByUserName(username string) (*model.SysUser, error) {
	structEntity := &model.SysUser{}
	// 尝试从缓存中获取
	userBytes, err := this.cache.Get([]byte(fmt.Sprintf("UserName:%s", username)))
	if err == nil {
		// 缓存命中
		err = json.Unmarshal(userBytes, &structEntity)
		if err == nil && structEntity.UserID != 0 {
			// 缓存命中
			return structEntity, nil
		}
	}

	structEntity, err = this.repo.QueryUserByUserName(username)
	if err != nil {
		this.logger.Error("查询用户信息失败", zap.Error(err))
		return nil, err
	} else if structEntity.UserID != 0 {
		// 序列化用户对象并存入缓存
		userBytes, err = json.Marshal(structEntity)
		if err == nil {
			this.cache.Set([]byte(fmt.Sprintf("UserName:%s", username)), userBytes, common.EXPIRESECONDS)          // 第三个参数是过期时间，0表示永不过期
			this.cache.Set([]byte(fmt.Sprintf("UserID:%d", structEntity.UserID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
		}
		return structEntity, nil
	}

	this.logger.Debug("查询用户信息失败", zap.Error(err))
	return nil, fmt.Errorf("查询用户信息失败", zap.Error(err))
}

// QueryUserLikeUserName 模糊查询用户名
// 根据用户名关键字进行模糊匹配查询
//
// 参数:
//
//	username: 用户名关键字
//
// 返回值:
//
//	[]*model.SysUser: 符合条件的用户列表
//	error: 错误信息
func (this *SysUserService) QueryUserLikeUserName(username string) ([]*model.SysUser, error) {
	structEntity, err := this.repo.QueryUserLikeUserName(username)
	if err != nil {
		this.logger.Error("查询用户信息失败", zap.Error(err))
		return nil, err
	}
	return structEntity, nil
}

// QueryUserByUserId 根据用户ID查询用户信息
// 支持缓存机制，优先从缓存获取，缓存未命中则查询数据库并更新缓存
//
// 参数:
//
//	userId: 用户ID
//
// 返回值:
//
//	*model.SysUser: 用户信息
//	error: 错误信息
func (this *SysUserService) QueryUserByUserId(userId int64) (*model.SysUser, error) {
	structEntity := &model.SysUser{}
	if userId == 0 {
		return structEntity, nil
	}

	// 尝试从缓存中获取
	userBytes, err := this.cache.Get([]byte(fmt.Sprintf("UserID:%d", userId)))
	if err == nil {
		// 缓存命中
		err = json.Unmarshal(userBytes, &structEntity)
		if err == nil && structEntity.UserID != 0 {
			// 缓存命中
			return structEntity, nil
		}
	}

	structEntity, err = this.repo.QueryUserByUserId(userId)
	if err != nil {
		this.logger.Error("查询用户信息失败", zap.Error(err))
		return nil, err
	} else if structEntity.UserID != 0 {
		// 序列化用户对象并存入缓存
		userBytes, err = json.Marshal(structEntity)
		if err == nil {
			this.cache.Set([]byte(fmt.Sprintf("UserID:%d", structEntity.UserID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
		}
		return structEntity, nil
	}

	this.logger.Debug("查询用户信息失败", zap.Error(err))
	return nil, fmt.Errorf("查询用户信息失败", zap.Error(err))
}

// QueryUserPage 分页查询用户列表
// 支持多种筛选条件的分页查询，并关联查询用户的角色和部门信息
// 参数:
//   - pageReq: 分页请求参数
//   - u: 用户查询条件
//
// 返回值:
//   - []*model.UserInfoStruct: 用户信息列表（包含角色和部门信息）
//   - int64: 总记录数
//   - error: 错误信息
func (this *SysUserService) QueryUserPage(pageReq common.PageRequest, u *model.SysUserRequest) ([]*model.UserInfoStruct, int64, error) {
	data, total, err := this.repo.QueryUserPage(pageReq, u)
	if err != nil {
		this.logger.Error("查询用户分页信息失败", zap.Error(err))
		return nil, 0, err
	}

	userList := make([]*model.UserInfoStruct, 0)
	for _, user := range data {
		userInfo := &model.UserInfoStruct{}
		userInfo.SysUser = user
		if user.UserID == common.ADMINID {
			userInfo.Admin = true
		}

		roles, err := this.roleRepo.QueryRolesByUserId(user.UserID)
		if err != nil {
			this.logger.Error("QueryRolesByUserId error,", zap.Error(err))
			return nil, 0, fmt.Errorf("getInfo error", zap.Error(err))
		}
		userInfo.Roles = roles

		dept, err := this.deptRepo.QueryDeptById(user.DeptID)
		if err != nil {
			this.logger.Error("QueryDeptById error,", zap.Error(err))
			return nil, 0, fmt.Errorf("getInfo error", zap.Error(err))
		}
		userInfo.Dept = dept
		userList = append(userList, userInfo)
	}

	return userList, total, nil
}

// QueryUserList 查询用户列表
// 不分页的用户列表查询，支持多种筛选条件
// 参数:
//   - u: 用户查询条件
//
// 返回值:
//   - []*model.SysUser: 用户列表
//   - error: 错误信息
func (this *SysUserService) QueryUserList(u *model.SysUserRequest) ([]*model.SysUser, error) {
	data, err := this.repo.QueryUserList(u)
	if err != nil {
		this.logger.Error("查询用户列表信息失败", zap.Error(err))
		return nil, err
	}
	return data, nil
}

// DeleteUserByUserId 根据用户ID删除用户
// 逻辑删除用户，并清除相关缓存
// 参数:
//   - userId: 用户ID
//
// 返回值:
//   - int64: 影响的行数
//   - error: 错误信息
func (this *SysUserService) DeleteUserByUserId(userId int64) (int64, error) {
	result, err := this.repo.DeleteUserByUserId(userId)
	if err != nil {
		this.logger.Error("删除用户信息失败", zap.Error(err))
		return 0, err
	}
	if result > 0 {
		this.cache.Del(fmt.Sprintf("UserID:%d", userId))
	}
	return result, nil
}

// ChangeUserStatus 修改用户状态
// 启用或禁用用户账户，并更新缓存
// 参数:
//   - user: 用户状态修改请求
//
// 返回值:
//   - int64: 影响的行数
//   - error: 错误信息
func (this *SysUserService) ChangeUserStatus(user *model.ChangeUserStatusRequest) (int64, error) {
	result, err := this.repo.ChangeUserStatus(user)
	if err != nil {
		this.logger.Error("修改用户状态失败", zap.Error(err))
		return 0, err
	}
	if result > 0 {
		structEntity, err := this.repo.QueryUserByUserId(user.UserID)
		if err != nil {
			this.logger.Error("查询用户信息失败", zap.Error(err))
			return 0, err
		} else if structEntity.UserID != 0 {
			// 序列化用户对象并存入缓存
			userBytes, err := json.Marshal(structEntity)
			if err == nil {
				this.cache.Set([]byte(fmt.Sprintf("UserID:%d", structEntity.UserID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
			}
			return result, nil
		}
	}
	return result, nil
}

// ResetUserPwd 重置用户密码
// 为用户设置新的密码
// 参数:
//   - user: 密码重置请求
//
// 返回值:
//   - int64: 影响的行数
//   - error: 错误信息
func (this *SysUserService) ResetUserPwd(user *model.ResetUserPwdRequest) (int64, error) {
	result, err := this.repo.ResetUserPwd(user)
	if err != nil {
		this.logger.Error("修改用户状态失败", zap.Error(err))
		return 0, err
	}
	return result, nil
}

// AddUser 添加新用户
// 创建新用户记录，并将用户信息缓存
// 参数:
//   - post: 用户信息
//
// 返回值:
//   - *model.SysUser: 创建的用户信息
//   - error: 错误信息
func (this *SysUserService) AddUser(post *model.SysUser) (*model.SysUser, error) {
	data, err := this.repo.AddUser(post)
	if err != nil {
		this.logger.Error("AddUser", zap.Error(err))
		return nil, err
	}
	if data != nil && data.UserID != 0 {
		// 序列化用户对象并存入缓存
		userBytes, err := json.Marshal(data)
		if err == nil {
			this.cache.Set([]byte(fmt.Sprintf("UserID:%d", data.UserID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
		}
	}
	return data, nil
}

// EditUser 编辑用户信息
// 更新现有用户的信息，并同步更新缓存
// 参数:
//   - post: 更新后的用户信息
//
// 返回值:
//   - *model.SysUser: 更新后的用户信息
//   - int64: 影响的行数
//   - error: 错误信息
func (this *SysUserService) EditUser(post *model.SysUser) (*model.SysUser, int64, error) {
	data, result, err := this.repo.EditUser(post)
	if err != nil {
		this.logger.Error("EditUser", zap.Error(err))
		return nil, 0, err
	}
	if data != nil && data.UserID != 0 && result == 1 {
		// 序列化用户对象并存入缓存
		userBytes, err := json.Marshal(data)
		if err == nil {
			this.cache.Set([]byte(fmt.Sprintf("UserID:%d", data.UserID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
		}
	}
	return data, result, nil
}

// CheckUserNameUnique 检查用户名唯一性
// 验证指定用户名是否已被其他用户使用
// 参数:
//   - id: 排除的用户ID（编辑时使用）
//   - name: 待检查的用户名
//
// 返回值:
//   - int64: 重复数量（0表示唯一，>0表示重复）
//   - error: 错误信息
func (this *SysUserService) CheckUserNameUnique(id int64, name string) (int64, error) {
	result, err := this.repo.CheckUserNameUnique(id, name)
	if err != nil {
		this.logger.Error("CheckUserNameUnique", zap.Error(err))
		return -1, err
	}
	return result, nil
}

// UserLogin 记录用户登录信息
// 更新用户的最后登录时间
// 参数:
//   - user: 用户信息
//
// 返回值:
//   - int64: 影响的行数
//   - error: 错误信息
func (this *SysUserService) UserLogin(user *model.SysUser) (int64, error) {
	result, err := this.repo.UserLogin(user)
	if err != nil {
		this.logger.Error("UserLogin", zap.Error(err))
		return -1, err
	}
	return result, nil
}
