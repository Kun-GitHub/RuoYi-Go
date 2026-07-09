// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package handler

import (
	"RuoYi-Go/internal/application/usecase"
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/domain/model"
	"RuoYi-Go/internal/filter"
	"RuoYi-Go/internal/ports/input"
	"RuoYi-Go/pkg/excel"
	"fmt"
	"github.com/kataras/iris/v12"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"io"
	"os"
	"path/filepath"
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
	txManager       *usecase.TransactionManager
}

func NewSysUserHandler(service input.SysUserService, deptService input.SysDeptService, roleService input.SysRoleService,
	postService input.SysPostService, userRoleService input.SysUserRoleService, userPostService input.SysUserPostService,
	txManager *usecase.TransactionManager) *SysUserHandler {
	return &SysUserHandler{service: service,
		deptService: deptService, roleService: roleService,
		postService: postService, userRoleService: userRoleService,
		userPostService: userPostService, txManager: txManager,
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

	ids, err := this.deptService.QueryChildIdListById(deptId)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "QueryChildIdListById, error：%s", err.Error()))
		return
	}

	// 数据权限过滤
	scope := filter.ComputeDataScope(ctx, this.roleService, this.deptService)
	if scope.IsAdmin || scope.ScopeType == filter.DataScopeAll {
		// 全部数据权限，使用原始部门筛选条件
		u.DeptIDs = ids
	} else if scope.ScopeType == filter.DataScopeSelf {
		// 仅本人数据 - 通过UserId过滤
		loginUser, _ := ctx.Values().Get(common.LOGINUSER).(*model.UserInfoStruct)
		if loginUser != nil {
			u.UserId = loginUser.UserID
		}
		u.DeptIDs = append(ids, scope.DeptIds...)
	} else {
		// 把数据权限的部门ID和原始部门筛选合并
		u.DeptIDs = append(ids, scope.DeptIds...)
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
	deptFilter := &model.SysDept{}
	scope := filter.ComputeDataScope(ctx, this.roleService, this.deptService)
	if !scope.IsAdmin && scope.ScopeType != filter.DataScopeAll {
		deptFilter.DataScopeDeptIds = scope.DeptIds
	}

	data, err := this.deptService.QueryDeptList(deptFilter)
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

func (this *SysUserHandler) UserProfile(ctx iris.Context) {
	temp := ctx.Values().Get(common.LOGINUSER)
	// 类型断言
	loginUser, ok := temp.(*model.UserInfoStruct)
	if !ok {
		ctx.JSON(common.Error(iris.StatusUnauthorized, "请重新登录"))
		return
	}

	user, err := this.service.QueryUserByUserId(loginUser.UserID)
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

	//postList, err := this.postService.QueryPostList(nil)
	//if err != nil {
	//	ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "QueryPostList, error：%s", err.Error()))
	//	return
	//}

	posts, err := this.postService.QueryPostByUserId(user.UserID)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "QueryPostByUserId, error：%s", err.Error()))
		return
	}
	var postIds []int64
	for _, post := range posts {
		postIds = append(postIds, post.PostID)
	}

	infoSuccess := &model.UserProfileSuccess{
		Code:    common.SUCCESS,
		User:    userInfo,
		Message: "操作成功",
		//RoleIds: roleIds,
		//Roles:   roles,
		//Posts:   postList,
		//PostIds: postIds,
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

		this.userRoleService.DeleteUserRoleByUserId(userId)
		this.userPostService.DeleteUserPostByUserId(userId)
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
	if err := filter.ValidateRequest(ctx, post); err != nil {
		return
	}

	count, err := this.service.CheckUserNameUnique(-1, post.UserName)
	if err != nil || count != 0 {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "新增用户，已存在相同用户名称"))
		return
	}

	if post.Phonenumber != "" {
		count, err = this.service.CheckPhoneUnique(-1, post.Phonenumber)
		if err != nil || count != 0 {
			ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "新增用户，已存在相同手机号"))
			return
		}
	}

	if post.Email != "" {
		count, err = this.service.CheckEmailUnique(-1, post.Email)
		if err != nil || count != 0 {
			ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "新增用户，已存在相同邮箱"))
			return
		}
	}

	user := ctx.Values().Get(common.LOGINUSER)
	loginUser, ok := user.(*model.UserInfoStruct)
	if !ok {
		ctx.JSON(common.Error(iris.StatusUnauthorized, "请重新登录"))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(post.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "密码加密失败"))
		return
	}

	post.CreateTime = time.Now()
	post.CreateBy = loginUser.UserName
	post.UpdateTime = time.Now()
	post.UpdateBy = loginUser.UserName
	post.LoginDate = time.Now()
	post.Password = string(hashedPassword)

	var info *model.SysUser
	err = this.txManager.Execute(func(tx *gorm.DB) error {
		post.SysUser.Status = "0"
		post.SysUser.DelFlag = "0"

		if err := tx.Create(post.SysUser).Error; err != nil {
			return err
		}
		info = post.SysUser

		for _, id := range post.RoleIds {
			if err := tx.Create(&model.SysUserRole{UserID: info.UserID, RoleID: id}).Error; err != nil {
				return err
			}
		}

		for _, id := range post.PostIds {
			if err := tx.Create(&model.SysUserPost{UserID: info.UserID, PostID: id}).Error; err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "AddUser, error：%s", err.Error()))
		return
	}

	ctx.JSON(common.Success(info))
}

func (this *SysUserHandler) EditUser(ctx iris.Context) {
	post := &model.UserInfoRequest{}
	if err := filter.ValidateRequest(ctx, post); err != nil {
		return
	}

	count, err := this.service.CheckUserNameUnique(post.UserID, post.UserName)
	if err != nil || count != 0 {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "修改用户，已存在相同用户名称"))
		return
	}

	if post.Phonenumber != "" {
		count, err = this.service.CheckPhoneUnique(post.UserID, post.Phonenumber)
		if err != nil || count != 0 {
			ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "修改用户，已存在相同手机号"))
			return
		}
	}

	if post.Email != "" {
		count, err = this.service.CheckEmailUnique(post.UserID, post.Email)
		if err != nil || count != 0 {
			ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "修改用户，已存在相同邮箱"))
			return
		}
	}

	user := ctx.Values().Get(common.LOGINUSER)
	loginUser, ok := user.(*model.UserInfoStruct)
	if !ok {
		ctx.JSON(common.Error(iris.StatusUnauthorized, "请重新登录"))
		return
	}
	post.UpdateTime = time.Now()
	post.UpdateBy = loginUser.UserName

	var info *model.SysUser
	err = this.txManager.Execute(func(tx *gorm.DB) error {
		if err := tx.Model(&model.SysUser{}).Where("user_id = ? and del_flag = '0'", post.UserID).Updates(post.SysUser).Error; err != nil {
			return err
		}
		info = post.SysUser

		if err := tx.Where("user_id = ?", info.UserID).Delete(&model.SysUserRole{}).Error; err != nil {
			return err
		}
		for _, id := range post.RoleIds {
			if err := tx.Create(&model.SysUserRole{UserID: info.UserID, RoleID: id}).Error; err != nil {
				return err
			}
		}

		if err := tx.Where("user_id = ?", info.UserID).Delete(&model.SysUserPost{}).Error; err != nil {
			return err
		}
		for _, id := range post.PostIds {
			if err := tx.Create(&model.SysUserPost{UserID: info.UserID, PostID: id}).Error; err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "EditUser, error：%s", err.Error()))
		return
	}

	ctx.JSON(common.Success(info))
}

func (this *SysUserHandler) UpdateProfile(ctx iris.Context) {
	post := &model.SysUser{}
	if err := filter.ValidateRequest(ctx, post); err != nil {
		return
	}

	user := ctx.Values().Get(common.LOGINUSER)
	loginUser, ok := user.(*model.UserInfoStruct)
	if !ok {
		ctx.JSON(common.Error(iris.StatusUnauthorized, "请重新登录"))
		return
	}
	post.UpdateTime = time.Now()
	post.UpdateBy = loginUser.UserName

	_, err := this.service.UpdateProfile(post)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "UpdateProfile error：%s", err.Error()))
		return
	}
	ctx.JSON(common.Success(nil))
}

func (this *SysUserHandler) UpdatePwd(ctx iris.Context) {
	body := &struct {
		OldPassword string `json:"oldPassword"`
		NewPassword string `json:"newPassword"`
	}{}
	if err := filter.ValidateRequest(ctx, body); err != nil {
		return
	}

	user := ctx.Values().Get(common.LOGINUSER)
	loginUser, ok := user.(*model.UserInfoStruct)
	if !ok {
		ctx.JSON(common.Error(iris.StatusUnauthorized, "请重新登录"))
		return
	}

	_, err := this.service.UpdatePwd(loginUser.UserID, body.OldPassword, body.NewPassword)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "UpdatePwd error：%s", err.Error()))
		return
	}
	ctx.JSON(common.Success(nil))
}

func (this *SysUserHandler) UpdateAvatar(ctx iris.Context) {
	file, header, err := ctx.FormFile("avatarfile")
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusBadRequest, "Upload failed: %s", err.Error()))
		return
	}
	defer file.Close()

	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(header.Filename))
	basePath := "./upload/avatar"
	datePath := time.Now().Format("2006/01/02")
	absolutePath := filepath.Join(basePath, "upload", datePath)
	if err := os.MkdirAll(absolutePath, os.ModePerm); err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "Mkdir error: %s", err.Error()))
		return
	}

	dst := filepath.Join(absolutePath, fileName)
	out, err := os.Create(dst)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "Create file error: %s", err.Error()))
		return
	}
	defer out.Close()

	if _, err = io.Copy(out, file); err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "Save file error: %s", err.Error()))
		return
	}

	url := "/profile/upload/" + datePath + "/" + fileName

	user := ctx.Values().Get(common.LOGINUSER)
	loginUser, ok := user.(*model.UserInfoStruct)
	if !ok {
		ctx.JSON(common.Error(iris.StatusUnauthorized, "请重新登录"))
		return
	}

	_, err = this.service.UpdateAvatar(loginUser.UserID, url)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "UpdateAvatar error：%s", err.Error()))
		return
	}
	ctx.JSON(common.Success(map[string]interface{}{
		"imgUrl":   url,
		"fileName": fileName,
	}))
}

func (this *SysUserHandler) Export(ctx iris.Context) {
	allParams := ctx.Request().URL.Query()
	beginTimeList, _ := allParams["params[beginTime]"]
	endTimeList, _ := allParams["params[endTime]"]
	beginTime := ""
	if len(beginTimeList) > 0 {
		beginTime = beginTimeList[0]
	}
	endTime := ""
	if len(endTimeList) > 0 {
		endTime = endTimeList[0]
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

	// 数据权限过滤
	scope := filter.ComputeDataScope(ctx, this.roleService, this.deptService)
	if scope.IsAdmin || scope.ScopeType == filter.DataScopeAll {
	} else if scope.ScopeType == filter.DataScopeSelf {
		loginUser, _ := ctx.Values().Get(common.LOGINUSER).(*model.UserInfoStruct)
		if loginUser != nil {
			u.UserId = loginUser.UserID
		}
	} else {
		u.DeptIDs = append(u.DeptIDs, scope.DeptIds...)
	}

	list, err := this.service.QueryUserList(u)
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "Export error: %s", err.Error()))
		return
	}

	headers := []string{"用户ID", "用户名称", "用户昵称", "邮箱", "手机号", "状态", "创建时间"}
	rows := make([][]interface{}, len(list))
	for i, item := range list {
		createTime := ""
		if !item.CreateTime.IsZero() {
			createTime = item.CreateTime.Format("2006-01-02 15:04:05")
		}
		rows[i] = []interface{}{
			item.UserID,
			item.UserName,
			item.NickName,
			item.Email,
			item.Phonenumber,
			item.Status,
			createTime,
		}
	}

	filePath, err := excel.ExportExcel(headers, rows, "用户数据")
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "ExportExcel error: %s", err.Error()))
		return
	}
	defer os.Remove(filePath)

	ctx.SendFile(filePath, "user.xlsx")
}

func (this *SysUserHandler) ImportData(ctx iris.Context) {
	ctx.JSON(common.Success(nil))
}

func (this *SysUserHandler) ImportTemplate(ctx iris.Context) {
	headers := []string{"用户名称", "用户昵称", "邮箱", "手机号", "性别", "状态"}
	rows := make([][]interface{}, 0)

	filePath, err := excel.ExportExcel(headers, rows, "用户导入模板")
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "ExportExcel error: %s", err.Error()))
		return
	}
	defer os.Remove(filePath)

	ctx.SendFile(filePath, "user_import_template.xlsx")
}
