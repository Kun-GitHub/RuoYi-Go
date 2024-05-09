# RuoYi-Go

### 1. 关于我
[个人介绍](https://github.com/Kun-GitHub)

### 2. 介绍
后端用Go写的RuoYi权限管理系统 (功能正在持续实现)

### 3. 前端
[RuoYi-Vue3 官方前端Vue3版](https://github.com/Kun-GitHub/RuoYi-Vue3)

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
</tbody>
</table>

### 5. 项目目录（持续在对齐项目，在补充）
```项目结构
time-management-tool/
|-- build/
|   |-- build.cmd     // go build 文件
|-- cmd/
|   |-- main.go       // 主入口文件
|-- config/           // 配置文件目录
|   |-- config.yml    // 示例配置文件
|-- internal/
|   |-- handler/      // HTTP请求处理器
|   |-- service/      // 业务逻辑层
|   |-- repository/   // 数据访问层
|-- pkg/
|   |-- i18n/         // 国际化
|-- middleware/       // 中间件（如JWT验证）
|-- model/            // 数据模型
|-- websocket/        // WebSocket处理逻辑
|-- go.mod            // Go模块依赖管理
|-- go.sum            // Go模块依赖哈希
```

### 6. 编写环境
[JetBrains Fleet](https://www.jetbrains.com/fleet) （目前还是免费用，类似微软的VS Code，个人刚用这个Fleet，不喜勿喷）

### 7. 最后
目前项目还是一个人写，边工作边写，主要是下班后写，所以可能会慢一些哈，如果不介意的话，点个 Start 持续跟踪，谢谢啦