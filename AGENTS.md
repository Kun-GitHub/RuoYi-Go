# RuoYi-Go Project Guide

## Overview
RuoYi-Go 是若依（RuoYi-Vue v3.9.1）的 Go 语言移植，使用 DDD 六边形架构 + Iris v12 框架，API 与 Java 版 1:1 对齐。

## Directories
```
RuoYi-Go/                      # Go 后端项目
├── di/container.go            # DI 容器 + 123 条路由注册
├── config/config.yaml         # 数据库/Redis 配置
├── internal/
│   ├── adapters/
│   │   ├── handler/           # HTTP 处理器（类似 Controller）
│   │   ├── persistence/       # Repository 实现（GORM Gen）
│   │   └── dao/               # GORM Gen 生成的 DAO 层
│   ├── application/usecase/   # 应用服务（协调业务逻辑）
│   ├── domain/
│   │   ├── model/             # 领域模型（贫血模型，gen 生成）
│   │   └── service/           # 领域服务（密码、数据权限等）
│   ├── filter/                # 中间件（认证/权限/操作日志/数据权限）
│   ├── ports/
│   │   ├── input/             # 输入端口接口（Service）
│   │   └── output/            # 输出端口接口（Repository）
│   ├── server/http_server.go  # DI 解析工厂
│   └── common/                # 公共常量/工具
└── pkg/
    ├── cache/                 # FreeCache + Redis
    ├── excel/                 # Excel 导出工具
    ├── jwt/                   # JWT 工具
    └── i18n/                  # 国际化
```

## Tech Stack
- **Web**: `github.com/kataras/iris/v12`
- **ORM**: `gorm.io/gen`（代码生成）+ `gorm.io/gorm`
- **Cache**: FreeCache（内存）+ Redis
- **Auth**: JWT + Redis session
- **Password**: `golang.org/x/crypto/bcrypt`
- **Excel**: `github.com/xuri/excelize/v2`
- **Logging**: `go.uber.org/zap`
- **I18n**: `github.com/nicksnyder/go-i18n/v2`

## Architecture: DDD Hexagonal (Ports & Adapters)

```
Handler (adapters/handler) → Service (ports/input → application/usecase) → Repository (ports/output → persistence) → DB
```

```
Middleware (filter/) → Handler → UseCase → Repository
     ├─ JWT auth                          ├─ TransactionManager
     ├─ Permission check                  ├─ Domain Service calls
     ├─ Operation log                     └─ Cache layer
     └─ LoginUser in ctx.Values()
```

## Coding Conventions

### Naming
- **Receiver**: always `this`（例外：`SysDeptHandler` 用 `h`）
- **Handlers**: `func (this *XxxHandler) Method(ctx iris.Context)`
- **UseCase receivers**: use `this`
- **Repository receivers**: use `this`
- **Variables**: lowercase, camelCase
- **Go files**: snake_case.go

### Imports
```go
"RuoYi-Go/internal/xxx"  // internal packages
ryjwt "RuoYi-Go/pkg/jwt" // pkg alias when name conflicts
```

### Key Patterns

**Handler** reads loginUser from context:
```go
loginUser, _ := ctx.Values().Get(common.LOGINUSER).(*model.UserInfoStruct)
```

**Repository** uses GORM Gen `.WithContext(context.Background())`:
```go
structEntity, err := this.db.Gen.SysUser.WithContext(context.Background()).
    Where(this.db.Gen.SysUser.DelFlag.Eq("0")).
    Find()
```

**Data scope** computed in handler:
```go
scope := filter.ComputeDataScope(ctx, this.roleService, this.deptService)
if !scope.IsAdmin && scope.ScopeType != filter.DataScopeAll {
    req.DeptIDs = append(req.DeptIDs, scope.DeptIds...)
}
```

**Transaction** via `txManager.Execute`:
```go
err = this.txManager.Execute(func(tx *gorm.DB) error {
    if err := tx.Create(record).Error; err != nil { return err }
    return nil
})
```

**Operation log** auto-recorded in middleware (POST/PUT/DELETE only):
- No per-route config needed; URL+Method auto-mapped in `autoLogOperation()`

**Permission middleware**:
```go
app.Get("/system/user/list", ms.PermissionMiddleware("system:user:list"), handler.UserPage)
```

### Business Type Constants (filter/http_middleware.go)
```go
BusinessTypeOther=0, Insert=1, Update=2, Delete=3, Grant=4,
Export=5, Import=6, Force=7, Clean=8, GenCode=9
```

### Data Scope Constants (filter/data_scope.go)
```go
DataScopeAll="1", DataScopeCustom="2", DataScopeDept="3",
DataScopeDeptAndChild="4", DataScopeSelf="5"
```

## Build & Run
```powershell
# Build
cd D:\Projects\RuoYi\RuoYi-Go
go build ./...

# Run (requires MySQL/PostgreSQL/SQLite + Redis)
go run main.go

# Config: config/config.yaml (server.port, database.*, redis.*)
```

## Current Implementation Status
- ✅ 123 routes registered in `di/container.go`
- ✅ Authentication & Permission middleware
- ✅ Operation log auto-recording
- ✅ Data scope filtering (5 types)
- ✅ Excel export/import
- ✅ Transaction manager
- ✅ Domain services (PasswordService, DataScopeService, OperationLogResolver)
- ❌ Code generator (`/tool/gen/*`) not implemented
- ❌ `monitor/server` system monitor (CPU/memory/disk) is stub
