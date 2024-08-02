// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package handler

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/domain/model"
	"RuoYi-Go/internal/filter"
	"RuoYi-Go/internal/ports/input"
	"github.com/kataras/iris/v12"
	"strconv"
	"strings"
	"time"
)

type SysUserHandler struct {
	service         input.SysUserService
	deptService     input.SysDeptService
	roleService     input.SysRoleService
	postService     input.SysPostService
	userRoleService input.SysUserRoleService
	userPostService input.SysUserPostService
}

func NewSysUserHandler(service input.SysUserService, deptService input.SysDeptService, roleService input.SysRoleService,
	postService input.SysPostService, userRoleService input.SysUserRoleService, userPostService input.SysUserPostService) *SysUserHandler {
	return &SysUserHandler{service: service,
		deptService: deptService, roleService: roleService,
		postService: postService, userRoleService: userRoleService,
		userPostService: userPostService,
	}
}

// GenerateCaptchaImage
func (this *SysUserHandler) UserPage(ctx iris.Context) {
	// 获取查询参数
	pageNumStr := ctx.URLParamDefault("pageNum", "1")
	pageSizeStr := ctx.URLParamDefault("pageSize", "10")

	// 使用 Query() 方法获取所有的查询参数
	allParams := ctx.Request().URL.Query()
	// 从 url.Values 结构体中获取参数
	beginTimeList, _ := allParams["params[beginTime]"]
	endTimeList, _ := allParams["params[endTime]"]
	// 假设我们只关心第一个值，我们可以这样获取：
	beginTime := ""
	if len(beginTimeList) > 0 {
		beginTime = beginTimeList[0]
	}
	endTime := ""
	if len(endTimeList) > 0 {
		endTime = endTimeList[0]
	}

	pageNum, _ := strconv.Atoi(pageNumStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)
	l := common.PageRequest{
		pageNum,
		pageSize,
	}

	status := ctx.URLParam("status")
	deptIdStr := ctx.URLParam("deptId")
	var deptId int64
	userName := ctx.URLParam("userName")
	phonenumber := ctx.URLParam("phonenumber")

	if deptIdStr != "" {
		deptId, _ = strconv.ParseInt(deptIdStr, 10, 64)
	}

	u := &model.SysUserRequest{
		Status:      status,
		DeptID:      deptId,
		UserName:    userName,
		Phonenumber: phonenumber,
		BeginTime:   beginTime,
		EndTime:     endTime,
	}

	d, t, err := this.service.QueryUserPage(l, u)
	if err != nil {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "UserPage, error：%s", err.Error()))
		return
	}

	data := &common.PageResponse{
		Rows:    d,
		Total:   t,
		Message: "操作成功",
		Code:    iris.StatusOK,
	}

	ctx.JSON(data)
}

func (this *SysUserHandler) DeptTree(ctx iris.Context) {
	data, err := this.deptService.QueryDeptList(nil)
	if err != nil {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "DeptTree, error：%s", err.Error()))
		return
	}

	data = buildDeptTree(data)
	ctx.JSON(common.Success(data))
}

// buildDeptTree constructs a tree of departments from a flat list.
func buildDeptTree(depts []*model.SysDept) []*model.SysDept {
	// Create a map to store the department by its ID.
	idToDept := make(map[int64]*model.SysDept)
	for _, dept := range depts {
		idToDept[dept.DeptID] = dept
	}

	// Create a slice to hold the root departments.
	var rootDepts []*model.SysDept

	// Iterate over the departments and set up the parent-child relationships.
	for _, dept := range depts {
		if parentId, exists := idToDept[dept.ParentID]; exists {
			// If the parent department exists, add the current department as its child.
			dept.ID = dept.DeptID
			dept.Label = dept.DeptName

			parentId.Children = append(parentId.Children, dept)
		} else {
			dept.ID = dept.DeptID
			dept.Label = dept.DeptName
			// If the parent department does not exist, it's a root department.
			rootDepts = append(rootDepts, dept)
		}
	}

	// Return the root departments.
	return rootDepts
}

func (this *SysUserHandler) UserInfo(ctx iris.Context) {
	userIdStr := ctx.Params().GetString("userId")
	if userIdStr == "" {
		ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid userIdStr"))
		return
	}

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "ParseInt error：%s", err.Error()))
		return
	}

	user, err := this.service.QueryUserByUserId(userId)
	if err != nil {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "QueryUserByUserId, error：%s", err.Error()))
		return
	}
	user.Password = ""

	userInfo := &model.UserInfoStruct{}
	userInfo.SysUser = user

	if user.UserID == common.ADMINID {
		userInfo.Admin = true
	}

	roles, err := this.roleService.QueryRolesByUserId(user.UserID)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "QueryRolesByUserId, error：%s", err.Error()))
		return
	}
	userInfo.Roles = roles

	dept, err := this.deptService.QueryDeptById(user.DeptID)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "QueryDeptById, error：%s", err.Error()))
		return
	}
	userInfo.Dept = dept

	var roleIds []int64
	for _, role := range roles {
		roleIds = append(roleIds, role.RoleID)
	}

	postList, err := this.postService.QueryPostList(nil)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "QueryPostList, error：%s", err.Error()))
		return
	}

	posts, err := this.postService.QueryPostByUserId(user.UserID)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "QueryPostByUserId, error：%s", err.Error()))
		return
	}
	var postIds []int64
	for _, post := range posts {
		postIds = append(postIds, post.PostID)
	}

	infoSuccess := &model.UserInfoSuccess{
		Code:    common.SUCCESS,
		User:    userInfo,
		Message: "操作成功",
		RoleIds: roleIds,
		Roles:   roles,
		Posts:   postList,
		PostIds: postIds,
	}

	ctx.JSON(infoSuccess)
}

//func (this *SysUserHandler) UserInfoByNoneUserId(ctx iris.Context) {
//	roles, err := this.roleService.QueryRoleList(nil)
//	if err != nil {
//		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "QueryRolesByUserId, error：%s", err.Error()))
//		return
//	}
//
//	postList, err := this.postService.QueryPostList(nil)
//	if err != nil {
//		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "QueryPostList, error：%s", err.Error()))
//		return
//	}
//
//	infoSuccess := &model.UserInfoNoneUserIdSuccess{
//		Code:    common.SUCCESS,
//		Message: "操作成功",
//		Roles:   roles,
//		Posts:   postList,
//	}
//
//	ctx.JSON(infoSuccess)
//}

func (this *SysUserHandler) ChangeUserStatus(ctx iris.Context) {
	u := &model.ChangeUserStatusRequest{}
	// Attempt to read and bind the JSON request body to the 'user' variable
	if err := filter.ValidateRequest(ctx, u); err != nil {
		//ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid JSON, error:%s", err.Error()))
		return
	}

	_, err := this.service.ChangeUserStatus(u)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "ChangeUserStatus error：%s", err.Error()))
		return
	}
	ctx.JSON(common.Success(nil))
}

func (this *SysUserHandler) ResetUserPwd(ctx iris.Context) {
	u := &model.ResetUserPwdRequest{}
	// Attempt to read and bind the JSON request body to the 'user' variable
	if err := filter.ValidateRequest(ctx, u); err != nil {
		//ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid JSON, error:%s", err.Error()))
		return
	}

	_, err := this.service.ResetUserPwd(u)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "ResetUserPwd error：%s", err.Error()))
		return
	}
	ctx.JSON(common.Success(nil))
}

func (this *SysUserHandler) DeleteUser(ctx iris.Context) {
	userIdStr := ctx.Params().GetString("userIds")
	if userIdStr == "" {
		ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid userIdStr"))
		return
	}

	parts := strings.Split(userIdStr, ",")
	for _, part := range parts {
		userId, err := strconv.ParseInt(part, 10, 64)
		if err != nil {
			ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "ParseInt error：%s", err.Error()))
			return
		}

		_, err = this.service.DeleteUserByUserId(userId)
		if err != nil {
			ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "DeleteUserByUserId error：%s", err.Error()))
			return
		}
	}

	ctx.JSON(common.Success(nil))
}

func (this *SysUserHandler) AuthRoleByUserId(ctx iris.Context) {
	userIdStr := ctx.Params().GetString("userId")
	if userIdStr == "" {
		ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid userIdStr"))
		return
	}

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "ParseInt error：%s", err.Error()))
		return
	}

	user, err := this.service.QueryUserByUserId(userId)
	if err != nil {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "QueryUserByUserId, error：%s", err.Error()))
		return
	}
	user.Password = ""

	userInfo := &model.UserInfoStruct{}
	userInfo.SysUser = user

	if user.UserID == common.ADMINID {
		userInfo.Admin = true
	}

	roles, err := this.roleService.QueryRolesByUserId(user.UserID)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "QueryRolesByUserId, error：%s", err.Error()))
		return
	}
	userInfo.Roles = roles

	dept, err := this.deptService.QueryDeptById(user.DeptID)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "QueryDeptById, error：%s", err.Error()))
		return
	}
	userInfo.Dept = dept

	roleList, err := this.roleService.QueryRoleList(nil)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "QueryRoleList, error：%s", err.Error()))
		return
	}

	infoSuccess := &model.AuthRoleSuccess{
		Code:    common.SUCCESS,
		User:    userInfo,
		Message: "操作成功",
		Roles:   roleList,
	}

	ctx.JSON(infoSuccess)
}

func (this *SysUserHandler) AuthRole(ctx iris.Context) {
	userIdStr := ctx.URLParamDefault("userId", "0")
	if userIdStr == "" {
		ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid userIdStr"))
		return
	}

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "ParseInt error：%s", err.Error()))
		return
	}
	this.userRoleService.DeleteUserRoleByUserId(userId)

	roleIdStr := ctx.URLParamDefault("roleIds", "0")
	if roleIdStr == "" {
		ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid roleIdStr"))
		return
	}

	roleIds := strings.Split(roleIdStr, ",")
	for _, id := range roleIds {
		roleId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "ParseInt error：%s", err.Error()))
			return
		}
		this.userRoleService.AddUserRole(&model.SysUserRole{
			userId,
			roleId,
		})
	}

	ctx.JSON(common.Success(nil))
}

func (this *SysUserHandler) AddUser(ctx iris.Context) {
	post := &model.UserInfoRequest{}
	// Attempt to read and bind the JSON request body to the 'user' variable
	if err := filter.ValidateRequest(ctx, post); err != nil {
		//ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid JSON, error:%s", err.Error()))
		return
	}

	count, err := this.service.CheckUserNameUnique(-1, post.UserName)
	if err != nil || count != 0 {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "新增用户，已存在相同用户名称"))
		return
	}

	user := ctx.Values().Get(common.LOGINUSER)
	// 类型断言
	loginUser, ok := user.(*model.UserInfoStruct)
	if !ok {
		ctx.JSON(common.Error(iris.StatusUnauthorized, "请重新登录"))
		return
	}
	post.CreateTime = time.Now()
	post.CreateBy = loginUser.UserName
	post.UpdateTime = time.Now()
	post.UpdateBy = loginUser.UserName
	post.LoginDate = time.Now()

	info, err := this.service.AddUser(post.SysUser)
	if err != nil {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "AddUser, error：%s", err.Error()))
		return
	}

	for _, id := range post.RoleIds {
		this.userRoleService.AddUserRole(&model.SysUserRole{
			info.UserID,
			id,
		})
	}

	for _, id := range post.PostIds {
		this.userPostService.AddUserPost(&model.SysUserPost{
			info.UserID,
			id,
		})
	}

	ctx.JSON(common.Success(info))
}

func (this *SysUserHandler) EditUser(ctx iris.Context) {
	post := &model.UserInfoRequest{}
	// Attempt to read and bind the JSON request body to the 'user' variable
	if err := filter.ValidateRequest(ctx, post); err != nil {
		//ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Invalid JSON, error:%s", err.Error()))
		return
	}

	count, err := this.service.CheckUserNameUnique(post.UserID, post.UserName)
	if err != nil || count != 0 {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "修改用户，已存在相同用户名称"))
		return
	}

	user := ctx.Values().Get(common.LOGINUSER)
	// 类型断言
	loginUser, ok := user.(*model.UserInfoStruct)
	if !ok {
		ctx.JSON(common.Error(iris.StatusUnauthorized, "请重新登录"))
		return
	}
	post.UpdateTime = time.Now()
	post.UpdateBy = loginUser.UserName

	info, _, err := this.service.EditUser(post.SysUser)
	if err != nil {
		//this.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "EditRole, error：%s", err.Error()))
		return
	}

	this.userRoleService.DeleteUserRoleByUserId(info.UserID)
	for _, id := range post.RoleIds {
		this.userRoleService.AddUserRole(&model.SysUserRole{
			info.UserID,
			id,
		})
	}

	this.userPostService.DeleteUserPostByUserId(info.UserID)
	for _, id := range post.PostIds {
		this.userPostService.AddUserPost(&model.SysUserPost{
			info.UserID,
			id,
		})
	}

	ctx.JSON(common.Success(info))
}
