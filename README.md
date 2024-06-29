# RuoYi-Go(DDD)

### 1. 关于我(在找远程工作，给机会的老板可以联系)
[个人介绍](https://github.com/Kun-GitHub)

<br>

### 2. 后端
后端是用Go写的RuoYi权限管理系统 (功能正在持续实现)   
用DDD领域驱动设计(六边形架构)做实践

后端 [GitHub地址](https://github.com/Kun-GitHub/RuoYi-Go)    

后端 [Gitee地址](https://gitee.com/gitee_kun/RuoYi-Go)

<br>

### 3. 前端
本项目没有自研前端，前端代码为 [RuoYi-Vue3 官方前端Vue3版](https://github.com/Kun-GitHub/RuoYi-Vue3)

<br>

### 4. Go后端技术栈（持续在对齐项目，在补充）
<table>
<thead>
<tr>
<th>功能</th>
<th>框架</th>
<th>是否采用</th>
<th>备注</th>
</tr>
</thead>
<tbody>
<tr>
<td rowspan="2">配置管理</td>
<td><a href="https://github.com/spf13/viper">Viper</a></td>
<td>✅</td>
<td>功能丰富，支持动态重载</td>
</tr>
<tr>
<td><a href="https://github.com/kelseyhightower/envconfig">Envconfig</a></td>
<td></td>
<td>轻量级库</td>
</tr>
<tr>
<td>验证码</td>
<td><a href="https://github.com/mojocn/base64Captcha">base64Captcha</a></td>
<td>✅</td>
<td>提供了生成各种类型验证码的功能</td>
</tr>
<tr>
<td rowspan="4">Web</td>
<td><a href="https://github.com/kataras/iris">Iris</a></td>
<td>✅</td>
<td>高性能、灵活且易于使用的Go Web框架</td>
</tr>
<tr>
<td><a href="https://github.com/gin-gonic/gin">Gin</a></td>
<td></td>
<td>快速且高效的Go Web框架</td>
</tr>
<tr>
<td><a href="https://github.com/gogf/gf">goFrame</a></td>
<td></td>
<td>高性能、模块化和企业级的全栈开发框架</td>
</tr>
<tr>
<td><a href="https://github.com/beego/beego">beego</a></td>
<td></td>
<td>全功能的MVC框架</td>
</tr>
<tr>
<td rowspan="3">ORM</td>
<td><a href="https://github.com/go-gorm/gorm">gorm</a></td>
<td>✅</td>
<td>Go语言中一个非常流行的ORM框架</td>
</tr>
<tr>
<td><a href="https://github.com/go-xorm/xorm">Xorm</a></td>
<td></td>
<td>简洁、易用且功能强大的Go语言ORM库，不过没维护了</td>
</tr>
<tr>
<td><a href="https://github.com/volatiletech/sqlboiler">SQLBoiler</a></td>
<td></td>
<td>通过Go的代码生成器来实现的ORM工具</td>
</tr>
<tr>
<td rowspan="3">内存缓存</td>
<td><a href="https://github.com/allegro/bigcache">Bigcache</a></td>
<td></td>
<td>高性能、持久化的键值存储库<br>
适合存储永不过期或者生命周期非常长的数据</td>
</tr>
<tr>
<td><a href="https://github.com/coocood/freecache">freecache</a></td>
<td>✅</td>
<td>高性能的内存缓存库</td>
</tr>
<tr>
<td><a href="https://github.com/golang/groupcache">Groupcache</a></td>
<td></td>
<td>Google开源的一个分布式缓存和缓存填充系统<br>
主要用于大型系统的缓存共享</td>
</tr>
<tr>
<td rowspan="4">日志记录</td>
<td><a href="https://github.com/rs/zerolog">zerolog</a></td>
<td></td>
<td>高性能的结构化日志库，专为JSON输出优化，支持零分配日志记录<br>
适合微服务和云原生应用</td>
</tr>
<tr>
<td><a href="https://github.com/uber-go/zap">Zap</a></td>
<td>✅</td>
<td>高性能、结构化的日志库，特别强调速度和效率<br>
项目配合用了lumberjack，实现日志文件的自动切割和管理功能</td>
</tr>
<tr>
<td><a href="https://github.com/sirupsen/logrus">Logrus</a></td>
<td></td>
<td>以其易用性和灵活性著称</td>
</tr>
<tr>
<td><a href="https://github.com/cihub/seelog">seelog</a></td>
<td></td>
<td>支持复杂的过滤规则、多级日志处理管道和多种输出目标</td>
</tr>
<tr>
<td rowspan="2">依赖注入</td>
<td><a href="https://github.com/google/wire">wire</a></td>
<td></td>
<td>由Google开源的依赖注入工具，它通过代码生成的方式，在编译时期完成依赖注入</td>
</tr>
<tr>
<td><a href="https://github.com/uber-go/dig">dig</a></td>
<td></td>
<td>提供了高性能和可读性，支持构造函数注入、函数参数注入和结构体字段注入</td>
</tr>
<tr>
<td>Redis</td>
<td><a href="https://github.com/redis/go-redis">go-redis/redis</a></td>
<td>✅</td>
<td></td>
</tr>
<tr>
<td>ORM 代码生成工具</td>
<td><a href="https://github.com/go-gorm/gen">go-gorm/gen</a></td>
<td>✅</td>
<td>Friendly & Safer GORM powered by Code Generation</td>
</tr>
<tr>
<td rowspan="2">JWT</td>
<td><a href="https://github.com/golang-jwt/jwt">jwt</a></td>
<td>✅</td>
<td><a href="https://github.com/dgrijalva/jwt-go">jwt-go</a> 衍生版</td>
</tr>
<tr>
<td><a href="https://github.com/lestrrat-go/jwx">jwx</a></td>
<td></td>
<td>实现各种 JWx（JWA/JWE/JWK/JWS/JWT，也称为 JOSE）技术的 Go 模块</td>
</tr>
<tr>
<td rowspan="2">参数校验</td>
<td><a href="https://github.com/go-playground/validator">validator</a></td>
<td>✅</td>
<td>提供了一种优雅的方式来定义和执行各种数据验证规则</td>
</tr>
<tr>
<td><a href="https://github.com/asaskevich/govalidator">govalidator</a></td>
<td></td>
<td>提供了多种内置的验证标签和自定义标签支持</td>
</tr>
</tbody>
</table>
  
功能模块对应的开源库，还有很多我未知的(基于个人认知局限)，以上只列了一部分，大佬有其他更好的欢迎提issue一起分享试用  

<br>

### 5. 数据库（后面再考虑要不要支持多几个数据库）
<table>
<thead>
<tr>
<th>ORM框架</th>
<th>数据库</th>
<th>是否采用</th>
<th>备注</th>
</tr>
</thead>
<tbody>
<tr>
<td rowspan="3">gorm</td>
<td><a href="https://www.postgresql.org">PostgreSQL</a></td>
<td>✅</td>
<td>默认</td>
</tr>
<tr>
<td><a href="https://www.mysql.com/">Mysql</a></td>
<td>✅</td>
<td>不用说的，很赞</td>
</tr>
<tr>
<td><a href="https://www.sqlite.org/">Sqlite</a></td>
<td>✅</td>
<td>如果用这个的话，需要重新用gorm生成模型文件</td>
</tr>
</tbody>
</table>

[RuoYi 数据库脚本](https://github.com/yangzongzhuan/RuoYi-Vue/blob/master/sql/ry_20240529.sql)

<br>

### 6. 项目目录（持续在对齐项目，在补充）
```项目结构
RuoYi-Go/
├── cmd/
│   └── api/
│       └── main.go
├── config/
│   └── config.yaml
├── internal/
│   ├── domain/
│   │   ├── model/
│   │   │   └── demo.go
│   ├── application/
│   │   └── usecase/
│   │       └── demo_usecase.go
│   ├── ports/
│   │   ├── input/
│   │   │   └── demo_service.go
│   │   └── output/
│   │       └── demo_repository.go
│   ├── adapters/
│   │   ├── api/
│   │   │   └── demo_handler.go
│   │   ├── persistence/
│   │   │   └── demo_repository.go
├── di/
│   └── container.go
├── pkg/
│   │   ├── db/
│   │   │   └── database.go
│   │   ├── jwt/
│   │   │   └── jwt.go
│   │   ├── logger/
│   │   │   └── logger.go
│   │   ├── config/
│   │   │   └── config.go
└── go.mod
```
<br>

### 7. 环境(工具)  
[Go 1.22.2](https://go.dev/doc/install) 

[Visual Studio Code](https://code.visualstudio.com/) 神器

[JetBrains Fleet](https://www.jetbrains.com/fleet) （目前还是免费用，类似微软的VS Code，不喜勿喷）
PS:发现暂不支持安装插件，不太好用

[DBeaver Community](https://dbeaver.io/) （SQL客户端和数据库管理工具）

[Another Redis Desktop Manager](https://github.com/qishibo/AnotherRedisDesktopManager) （Redis 客户端）

<br>

### 8. 致谢
致谢 [RuoYi](https://ruoyi.vip)  

致谢以上项目使用到的开源库，不分先后哈   

致谢以上开发用到的工具

<br>

### 9. 缺陷
1. 本项目是纯后端项目，前端是用RuoYi前端，所有为了适配RuoYi前端，有些写法会不太遵循Go语言的规范，不过不影响使用，只是为了适配RuoYi前端而已 
2. 有些工具类的引用没有使用依赖注入，或者上下文，目前用的是全局变量，暂时先这样，等后续个人经验丰富了，有更好的做法可能会改上去。
3. 当你有每一步都自己写的时候，你会发现若依的接口代码真***，还只是为了适配前端数据而已
<br>

### 10. 最后
目前项目还是一个人写，边工作边写，主要是下班后写，所以可能会慢一些哈，如果不介意的话，点个 Start 持续关注，谢谢啦，有什么建议可以提issue哈。




