package main

import (
	"RuoYi-Go/config"
	"RuoYi-Go/internal/adapters/dao"
	"RuoYi-Go/internal/adapters/handler"
	"RuoYi-Go/internal/application/usecase"
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/domain/model"
	"RuoYi-Go/internal/filter"
	"RuoYi-Go/internal/ports/output"
	"RuoYi-Go/pkg/cache"
	"testing"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"go.uber.org/zap"
)

// ===================== Mock Repositories (embedded interfaces) =====================

type mockUserRepo struct{ output.SysUserRepository }

func (m *mockUserRepo) QueryUserPage(_ common.PageRequest, _ *model.SysUserRequest) ([]*model.SysUser, int64, error) {
	return []*model.SysUser{{UserID: 1, UserName: "admin", NickName: "管理员", DeptID: 103, Status: "0", DelFlag: "0"}}, 1, nil
}
func (m *mockUserRepo) QueryUserByUserId(id int64) (*model.SysUser, error) {
	return &model.SysUser{UserID: id, UserName: "admin", NickName: "管理员", DeptID: 103, Status: "0", DelFlag: "0"}, nil
}
func (m *mockUserRepo) QueryUserList(_ *model.SysUserRequest) ([]*model.SysUser, error) {
	return []*model.SysUser{{UserID: 1}}, nil
}
func (m *mockUserRepo) AddUser(_ *model.SysUser) (*model.SysUser, error) {
	return &model.SysUser{UserID: 3}, nil
}
func (m *mockUserRepo) EditUser(u *model.SysUser) (*model.SysUser, int64, error) { return u, 1, nil }
func (m *mockUserRepo) CheckUserNameUnique(_ int64, _ string) (int64, error)     { return 0, nil }
func (m *mockUserRepo) CheckPhoneUnique(_ int64, _ string) (int64, error)        { return 0, nil }
func (m *mockUserRepo) CheckEmailUnique(_ int64, _ string) (int64, error)        { return 0, nil }

type mockRoleRepo struct{ output.SysRoleRepository }

func (m *mockRoleRepo) QueryRolesByUserId(_ int64) ([]*model.SysRole, error) {
	return []*model.SysRole{{RoleID: 1, RoleName: "管理员", RoleKey: "admin", DataScope: "1", Status: "0", DelFlag: "0"}}, nil
}
func (m *mockRoleRepo) QueryRolePage(_ common.PageRequest, _ *model.SysRoleRequest) ([]*model.SysRole, int64, error) {
	return []*model.SysRole{{RoleID: 1, RoleName: "管理员", RoleKey: "admin", DataScope: "1", Status: "0", DelFlag: "0"}}, 1, nil
}
func (m *mockRoleRepo) SelectRoleAll() ([]*model.SysRole, error) {
	return []*model.SysRole{{RoleID: 1, RoleName: "管理员", RoleKey: "admin", DataScope: "1"}}, nil
}
func (m *mockRoleRepo) AddRole(_ *model.SysRole) (*model.SysRole, error) {
	return &model.SysRole{RoleID: 1}, nil
}
func (m *mockRoleRepo) QueryRoleByID(_ int64) (*model.SysRole, error) {
	return &model.SysRole{RoleID: 1, RoleName: "管理员", DataScope: "1"}, nil
}
func (m *mockRoleRepo) QueryRoleList(_ *model.SysRoleRequest) ([]*model.SysRole, error) {
	return nil, nil
}

type mockMenuRepo struct{ output.SysMenuRepository }

func (m *mockMenuRepo) QueryMenuList(_ *model.SysMenuRequest) ([]*model.SysMenu, error) {
	return []*model.SysMenu{{MenuID: 1, MenuName: "系统管理", Perms: "system:user:list", MenuType: "M"}}, nil
}
func (m *mockMenuRepo) QueryMenuByID(_ int64) (*model.SysMenu, error) { return &model.SysMenu{}, nil }
func (m *mockMenuRepo) QueryMenuPage(_ common.PageRequest, _ *model.SysMenuRequest) ([]*model.SysMenu, int64, error) {
	return nil, 0, nil
}

type mockDeptRepo struct{ output.SysDeptRepository }

func (m *mockDeptRepo) QueryDeptList(_ *model.SysDept) ([]*model.SysDept, error) {
	return []*model.SysDept{
		{DeptID: 100, ParentID: 0, DeptName: "总公司", Status: "0", DelFlag: "0"},
		{DeptID: 101, ParentID: 100, DeptName: "研发部", Status: "0", DelFlag: "0"},
	}, nil
}
func (m *mockDeptRepo) QueryDeptById(_ int64) (*model.SysDept, error) {
	return &model.SysDept{DeptID: 100, DeptName: "测试部门", Status: "0", DelFlag: "0"}, nil
}
func (m *mockDeptRepo) QueryChildIdListById(_ int64) ([]int64, error) { return nil, nil }

type mockPostRepo struct{ output.SysPostRepository }

func (m *mockPostRepo) QueryPostPage(_ common.PageRequest, _ *model.SysPostRequest) ([]*model.SysPost, int64, error) {
	return []*model.SysPost{{PostID: 1, PostName: "项目经理", PostCode: "pm", Status: "0"}}, 1, nil
}
func (m *mockPostRepo) QueryPostList(_ *model.SysPostRequest) ([]*model.SysPost, error) {
	return []*model.SysPost{{PostID: 1, PostName: "项目经理", PostCode: "pm", Status: "0"}}, nil
}
func (m *mockPostRepo) QueryPostByPostId(_ int64) (*model.SysPost, error) {
	return &model.SysPost{}, nil
}
func (m *mockPostRepo) QueryPostByUserId(_ int64) ([]*model.SysPost, error) { return nil, nil }
func (m *mockPostRepo) SelectPostAll() ([]*model.SysPost, error)            { return nil, nil }

type mockDictTypeRepo struct{ output.SysDictTypeRepository }

func (m *mockDictTypeRepo) QueryDictTypeList(_ *model.SysDictTypeRequest) ([]*model.SysDictType, error) {
	return []*model.SysDictType{{DictID: 1, DictName: "用户性别", DictType: "sys_user_sex", Status: "0"}}, nil
}
func (m *mockDictTypeRepo) QueryDictTypePage(_ common.PageRequest, _ *model.SysDictTypeRequest) ([]*model.SysDictType, int64, error) {
	return []*model.SysDictType{{DictID: 1, DictName: "用户性别", DictType: "sys_user_sex", Status: "0"}}, 1, nil
}
func (m *mockDictTypeRepo) QueryDictTypeByDictID(_ int64) (*model.SysDictType, error) {
	return &model.SysDictType{}, nil
}
func (m *mockDictTypeRepo) AddDictType(_ *model.SysDictType) (*model.SysDictType, error) {
	return &model.SysDictType{DictID: 1}, nil
}
func (m *mockDictTypeRepo) EditDictType(p *model.SysDictType) (*model.SysDictType, int64, error) {
	return p, 1, nil
}
func (m *mockDictTypeRepo) DeleteDictTypeById(_ int64) (int64, error)            { return 1, nil }
func (m *mockDictTypeRepo) CheckDictTypeUnique(_ int64, _ string) (int64, error) { return 0, nil }

type mockDictDataRepo struct{ output.SysDictDataRepository }

func (m *mockDictDataRepo) QueryDictDatasByType(_ string) ([]*model.SysDictDatum, error) {
	return []*model.SysDictDatum{
		{DictCode: 1, DictLabel: "男", DictValue: "0", DictType: "sys_user_sex"},
		{DictCode: 2, DictLabel: "女", DictValue: "1", DictType: "sys_user_sex"},
	}, nil
}
func (m *mockDictDataRepo) Get(_ uint) (*model.SysDictDatum, error) {
	return &model.SysDictDatum{}, nil
}
func (m *mockDictDataRepo) List(_ int, _ int, _ string, _ string, _ string) ([]*model.SysDictDatum, int64, error) {
	return []*model.SysDictDatum{{DictCode: 1, DictLabel: "男", DictValue: "0", DictType: "sys_user_sex"}}, 1, nil
}

type mockConfigRepo struct{ output.SysConfigRepository }

func (m *mockConfigRepo) QueryConfigPage(_ common.PageRequest, _ *model.SysConfigRequest) ([]*model.SysConfig, int64, error) {
	return []*model.SysConfig{{ConfigID: 1, ConfigName: "初始密码", ConfigKey: "sys.user.initPassword", ConfigValue: "123456", ConfigType: "Y"}}, 1, nil
}
func (m *mockConfigRepo) QueryConfigByKey(_ string) (*model.SysConfig, error) {
	return &model.SysConfig{ConfigID: 1, ConfigName: "初始密码", ConfigKey: "sys.user.initPassword", ConfigValue: "123456", ConfigType: "Y"}, nil
}
func (m *mockConfigRepo) QueryConfigByID(_ int64) (*model.SysConfig, error) {
	return &model.SysConfig{}, nil
}
func (m *mockConfigRepo) QueryConfigList(_ *model.SysConfigRequest) ([]*model.SysConfig, error) {
	return nil, nil
}
func (m *mockConfigRepo) AddConfig(_ *model.SysConfig) (*model.SysConfig, error) {
	return &model.SysConfig{ConfigID: 1}, nil
}
func (m *mockConfigRepo) EditConfig(p *model.SysConfig) (*model.SysConfig, int64, error) {
	return p, 1, nil
}
func (m *mockConfigRepo) DeleteConfigById(_ int64) (int64, error)                { return 1, nil }
func (m *mockConfigRepo) CheckConfigNameUnique(_ int64, _ string) (int64, error) { return 0, nil }

type mockNoticeRepo struct{ output.SysNoticeRepository }

func (m *mockNoticeRepo) QueryNoticePage(_ common.PageRequest, _ *model.SysNoticeRequest) ([]*model.SysNotice, int64, error) {
	return []*model.SysNotice{{NoticeID: 1, NoticeTitle: "测试公告", NoticeType: "1", Status: "0"}}, 1, nil
}
func (m *mockNoticeRepo) QueryNoticeByID(_ int64) (*model.SysNotice, error) {
	return &model.SysNotice{}, nil
}
func (m *mockNoticeRepo) QueryNoticeList(_ *model.SysNoticeRequest) ([]*model.SysNotice, error) {
	return nil, nil
}
func (m *mockNoticeRepo) AddNotice(_ *model.SysNotice) (*model.SysNotice, error) {
	return &model.SysNotice{NoticeID: 1}, nil
}
func (m *mockNoticeRepo) EditNotice(p *model.SysNotice) (*model.SysNotice, int64, error) {
	return p, 1, nil
}
func (m *mockNoticeRepo) DeleteNoticeById(_ int64) (int64, error) { return 1, nil }

type mockOperLogRepo struct{ output.SysOperLogRepository }

func (m *mockOperLogRepo) Create(_ *model.SysOperLog) error { return nil }
func (m *mockOperLogRepo) List(_ int, _ int, _ string, _ string, _ string, _ string, _ string, _ []string) ([]*model.SysOperLog, int64, error) {
	return nil, 0, nil
}
func (m *mockOperLogRepo) Get(_ int64) (*model.SysOperLog, error) { return nil, nil }
func (m *mockOperLogRepo) Delete(_ []int64) error                 { return nil }
func (m *mockOperLogRepo) Clean() error                           { return nil }

type mockUserRoleRepo struct{ output.SysUserRoleRepository }

func (m *mockUserRoleRepo) AddUserRole(p *model.SysUserRole) (*model.SysUserRole, error) {
	return p, nil
}
func (m *mockUserRoleRepo) DeleteUserRoleByUserId(_ int64) (int64, error) { return 0, nil }

type mockUserPostRepo struct{ output.SysUserPostRepository }

func (m *mockUserPostRepo) AddUserPost(p *model.SysUserPost) (*model.SysUserPost, error) {
	return p, nil
}
func (m *mockUserPostRepo) DeleteUserPostByUserId(_ int64) (int64, error) { return 0, nil }

type mockRoleMenuRepo struct{ output.SysRoleMenuRepository }

func (m *mockRoleMenuRepo) AddRoleMenu(p *model.SysRoleMenu) (*model.SysRoleMenu, error) {
	return p, nil
}
func (m *mockRoleMenuRepo) DeleteRoleMenuByRoleId(_ int64) (int64, error) { return 0, nil }

type mockLoginRepo struct{ output.SysLogininforRepository }

func (m *mockLoginRepo) AddLogininfor(_ *model.SysLogininfor) (*model.SysLogininfor, error) {
	return nil, nil
}
func (m *mockLoginRepo) QueryLogininforByID(_ int64) (*model.SysLogininfor, error) { return nil, nil }
func (m *mockLoginRepo) QueryLogininforList(_ *model.SysLogininforRequest) ([]*model.SysLogininfor, error) {
	return nil, nil
}
func (m *mockLoginRepo) QueryLogininforPage(_ common.PageRequest, _ *model.SysLogininforRequest) ([]*model.SysLogininfor, int64, error) {
	return nil, 0, nil
}
func (m *mockLoginRepo) EditLogininfor(_ *model.SysLogininfor) (*model.SysLogininfor, int64, error) {
	return nil, 0, nil
}
func (m *mockLoginRepo) DeleteLogininforById(_ int64) (int64, error) { return 0, nil }

// ===================== Test Setup =====================

func newTestApp() *iris.Application {
	logger := zap.NewNop()
	fc := cache.NewFreeCacheClient(1024 * 1024)
	dbS := &dao.DatabaseStruct{}
	redisCli := &cache.RedisClient{}

	userRepo := &mockUserRepo{}
	roleRepo := &mockRoleRepo{}
	menuRepo := &mockMenuRepo{}
	deptRepo := &mockDeptRepo{}
	postRepo := &mockPostRepo{}

	userSvc := usecase.NewPageSysUserService(userRepo, roleRepo, deptRepo, fc, logger)
	roleSvc := usecase.NewSysRoleService(roleRepo, fc, logger)
	menuSvc := usecase.NewSysMenuService(menuRepo, fc, logger)
	deptSvc := usecase.NewSysDeptService(deptRepo, fc, logger)
	postSvc := usecase.NewSysPostService(postRepo, fc, logger)
	dictTypeSvc := usecase.NewSysDictTypeService(&mockDictTypeRepo{}, fc, logger)
	dictDataSvc := usecase.NewSysDictDataService(&mockDictDataRepo{}, fc, logger)
	configSvc := usecase.NewSysConfigService(&mockConfigRepo{}, fc, logger)
	noticeSvc := usecase.NewSysNoticeService(&mockNoticeRepo{}, fc, logger)
	operLogSvc := usecase.NewSysOperLogService(&mockOperLogRepo{}, fc, logger)
	userRoleSvc := usecase.NewSysUserRoleService(&mockUserRoleRepo{}, fc, logger)
	userPostSvc := usecase.NewSysUserPostService(&mockUserPostRepo{}, fc, logger)
	roleMenuSvc := usecase.NewSysRoleMenuService(&mockRoleMenuRepo{}, fc, logger)

	ms := filter.NewServerMiddleware(dbS, redisCli, logger, config.AppConfig{}, userSvc, menuSvc, operLogSvc)

	userH := handler.NewSysUserHandler(userSvc, deptSvc, roleSvc, postSvc, userRoleSvc, userPostSvc, usecase.NewTransactionManager(dbS))
	roleH := handler.NewSysRoleHandler(roleSvc, roleMenuSvc, deptSvc)
	menuH := handler.NewSysMenuHandler(menuSvc)
	deptH := handler.NewSysDeptHandler(deptSvc, roleSvc)
	postH := handler.NewSysPostHandler(postSvc)
	dictTypeH := handler.NewSysDictTypeHandler(dictTypeSvc)
	dictDataH := handler.NewSysDictDataHandler(dictDataSvc)
	configH := handler.NewSysConfigHandler(configSvc)
	noticeH := handler.NewSysNoticeHandler(noticeSvc)
	authH := handler.NewAuthHandler(
		usecase.NewAuthService(userSvc, roleSvc, deptSvc, configSvc, &mockLoginRepo{}, menuSvc, nil, logger),
		logger,
	)

	app := iris.New()

	// test middleware: inject loginUser (bypass JWT/Redis)
	app.Use(func(ctx iris.Context) {
		sysUser := &model.SysUser{UserID: 1, DeptID: 103, UserName: "admin", NickName: "管理员", Status: "0", DelFlag: "0"}
		loginUser := &model.UserInfoStruct{SysUser: sysUser, Admin: true}
		ctx.Values().Set(common.LOGINUSER, loginUser)
		ctx.Values().Set(common.USER_ID, "1")
		ctx.Next()
	})

	app.Get("/getInfo", authH.GetInfo)
	app.Get("/getRouters", menuH.GetRouters)
	app.Get("/system/user/list", ms.PermissionMiddleware("system:user:list"), userH.UserPage)
	app.Get("/system/user/profile", userH.UserProfile)
	app.Get("/system/user/{userId:uint}", ms.PermissionMiddleware("system:user:query"), userH.UserInfo)
	app.Post("/system/user", ms.PermissionMiddleware("system:user:add"), userH.AddUser)
	app.Get("/system/role/list", ms.PermissionMiddleware("system:role:list"), roleH.RolePage)
	app.Get("/system/role/{roleId:uint}", ms.PermissionMiddleware("system:role:query"), roleH.RoleInfo)
	app.Get("/system/role/optionselect", roleH.OptionSelect)
	app.Post("/system/role", ms.PermissionMiddleware("system:role:add"), roleH.AddRoleInfo)
	app.Get("/system/menu/list", ms.PermissionMiddleware("system:menu:list"), menuH.MenuList)
	app.Get("/system/menu/treeselect", menuH.TreeSelect)
	app.Get("/system/dept/list", ms.PermissionMiddleware("system:dept:list"), deptH.DeptList)
	app.Get("/system/dept/{deptId:uint}", ms.PermissionMiddleware("system:dept:query"), deptH.DeptInfo)
	app.Get("/system/post/list", ms.PermissionMiddleware("system:post:list"), postH.PostPage)
	app.Get("/system/post/optionselect", postH.OptionSelect)
	app.Get("/system/dict/type/list", ms.PermissionMiddleware("system:dict:type:list"), dictTypeH.DictTypePage)
	app.Get("/system/dict/type/optionselect", dictTypeH.DictTypeList)
	app.Get("/system/dict/data/type/{dictType:string}", dictDataH.DictType)
	app.Get("/system/dict/data/list", ms.PermissionMiddleware("system:dict:list"), dictDataH.List)
	app.Get("/system/config/list", ms.PermissionMiddleware("system:config:list"), configH.ConfigPage)
	app.Get("/system/config/configKey/{configKey:string}", configH.ConfigInfoByKey)
	app.Get("/system/notice/list", ms.PermissionMiddleware("system:notice:list"), noticeH.NoticePage)

	return app
}

// ===================== Test Cases =====================

func TestGetInfo(t *testing.T) {
	httptest.New(t, newTestApp()).GET("/getInfo").Expect().
		Status(iris.StatusOK).JSON().Object().ValueEqual("code", 200)
}

func TestGetRouters(t *testing.T) {
	httptest.New(t, newTestApp()).GET("/getRouters").Expect().
		Status(iris.StatusOK).JSON().Object().ValueEqual("code", 200)
}

func TestSystemUserList(t *testing.T) {
	httptest.New(t, newTestApp()).GET("/system/user/list").WithQuery("pageNum", "1").WithQuery("pageSize", "10").
		Expect().Status(iris.StatusOK).JSON().Object().ValueEqual("code", 200)
}

func TestSystemUserProfile(t *testing.T) {
	httptest.New(t, newTestApp()).GET("/system/user/profile").
		Expect().Status(iris.StatusOK).JSON().Object().ValueEqual("code", 200)
}

func TestSystemUserQuery(t *testing.T) {
	httptest.New(t, newTestApp()).GET("/system/user/1").
		Expect().Status(iris.StatusOK).JSON().Object().ValueEqual("code", 200)
}

func TestSystemRoleList(t *testing.T) {
	httptest.New(t, newTestApp()).GET("/system/role/list").WithQuery("pageNum", "1").WithQuery("pageSize", "10").
		Expect().Status(iris.StatusOK).JSON().Object().ValueEqual("code", 200)
}

func TestSystemRoleQuery(t *testing.T) {
	httptest.New(t, newTestApp()).GET("/system/role/1").
		Expect().Status(iris.StatusOK).JSON().Object().ValueEqual("code", 200)
}

func TestSystemRoleOptionSelect(t *testing.T) {
	httptest.New(t, newTestApp()).GET("/system/role/optionselect").
		Expect().Status(iris.StatusOK).JSON().Object().ValueEqual("code", 200)
}

func TestSystemMenuList(t *testing.T) {
	httptest.New(t, newTestApp()).GET("/system/menu/list").
		Expect().Status(iris.StatusOK).JSON().Object().ValueEqual("code", 200)
}

func TestSystemMenuTreeSelect(t *testing.T) {
	httptest.New(t, newTestApp()).GET("/system/menu/treeselect").
		Expect().Status(iris.StatusOK).JSON().Object().ValueEqual("code", 200)
}

func TestSystemDeptList(t *testing.T) {
	httptest.New(t, newTestApp()).GET("/system/dept/list").
		Expect().Status(iris.StatusOK).JSON().Object().ValueEqual("code", 200)
}

func TestSystemDeptQuery(t *testing.T) {
	httptest.New(t, newTestApp()).GET("/system/dept/100").
		Expect().Status(iris.StatusOK).JSON().Object().ValueEqual("code", 200)
}

func TestSystemPostList(t *testing.T) {
	httptest.New(t, newTestApp()).GET("/system/post/list").WithQuery("pageNum", "1").WithQuery("pageSize", "10").
		Expect().Status(iris.StatusOK).JSON().Object().ValueEqual("code", 200)
}

func TestSystemPostOptionSelect(t *testing.T) {
	httptest.New(t, newTestApp()).GET("/system/post/optionselect").
		Expect().Status(iris.StatusOK).JSON().Object().ValueEqual("code", 200)
}

func TestSystemDictTypeList(t *testing.T) {
	httptest.New(t, newTestApp()).GET("/system/dict/type/list").WithQuery("pageNum", "1").WithQuery("pageSize", "10").
		Expect().Status(iris.StatusOK).JSON().Object().ValueEqual("code", 200)
}

func TestSystemDictTypeOptionSelect(t *testing.T) {
	httptest.New(t, newTestApp()).GET("/system/dict/type/optionselect").
		Expect().Status(iris.StatusOK).JSON().Object().ValueEqual("code", 200)
}

func TestSystemDictDataByType(t *testing.T) {
	httptest.New(t, newTestApp()).GET("/system/dict/data/type/sys_user_sex").
		Expect().Status(iris.StatusOK).JSON().Object().ValueEqual("code", 200)
}

func TestSystemDictDataList(t *testing.T) {
	httptest.New(t, newTestApp()).GET("/system/dict/data/list").WithQuery("pageNum", "1").WithQuery("pageSize", "10").
		Expect().Status(iris.StatusOK).JSON().Object().ValueEqual("code", 200)
}

func TestSystemConfigList(t *testing.T) {
	httptest.New(t, newTestApp()).GET("/system/config/list").WithQuery("pageNum", "1").WithQuery("pageSize", "10").
		Expect().Status(iris.StatusOK).JSON().Object().ValueEqual("code", 200)
}

func TestSystemConfigByKey(t *testing.T) {
	httptest.New(t, newTestApp()).GET("/system/config/configKey/sys.user.initPassword").
		Expect().Status(iris.StatusOK).JSON().Object().ValueEqual("code", 200)
}

func TestSystemNoticeList(t *testing.T) {
	httptest.New(t, newTestApp()).GET("/system/notice/list").WithQuery("pageNum", "1").WithQuery("pageSize", "10").
		Expect().Status(iris.StatusOK).JSON().Object().ValueEqual("code", 200)
}
