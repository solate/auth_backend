# 开发指南

## 1. 开发环境搭建

### 1.1 必要软件
- Go 1.23+
- MySQL 8.0+
- Redis 7.0+
- Docker & Docker Compose
- Git
- Make

### 1.2 开发工具推荐
- VSCode
- Navicat (数据库管理工具)
- Postman (API测试工具)
- Redis Desktop Manager (Redis管理工具)

### 1.3 环境初始化


```
1. 克隆项目
git clone https://github.com/solate/auth_backend.git
2. 安装依赖
go mod tidy
3. 安装开发工具
go install github.com/zeromicro/go-zero/tools/goctl@latest
go install entgo.io/ent/cmd/ent@latest
4. 初始化数据库
make init-db
5. 生成代码
make gen
```


## 2. 项目结构说明

```
├── api # API定义目录
│ ├── auth # 认证服务API
│ ├── user # 用户服务API
│ └── permission # 权限服务API
├── cmd # 主程序入口
│ ├── api # API服务
│ └── job # 定时任务
├── configs # 配置文件
├── docs # 文档
├── internal # 内部代码
│ ├── config # 配置结构
│ ├── handler # 请求处理器
│ ├── logic # 业务逻辑
│ ├── middleware # 中间件
│ ├── svc # 服务上下文
│ └── types # 类型定义
├── model # 数据模型
│ ├── ent # Ent实体
│ └── cache # 缓存模型
├── pkg # 公共包
│ ├── errorx # 错误处理
│ ├── utils # 工具函数
│ └── middleware # 通用中间件
├── scripts # 脚本文件
└── deploy # 部署配置
```


## 3. 开发流程

### api 初始化




### 3.1 定义API

```
1. 创建API文件
api/user/user.api
syntax = "v1"
type (
CreateUserRequest {
Username string json:"username"
Password string json:"password"
Name string json:"name"
}
UserInfo {
Id int64 json:"id"
Username string json:"username"
Name string json:"name"
}
)
service user-api {
@handler createUser
post /api/v1/users (CreateUserRequest) returns (UserInfo)
}


2. 生成API代码
goctl api go -api user.api -dir . -style go_zero
```



### 3.2 定义数据模型

```
// model/ent/schema/user.go
package schema
import (
"entgo.io/ent"
"entgo.io/ent/schema/field"
)
type User struct {
ent.Schema
}
func (User) Fields() []ent.Field {
return []ent.Field{
field.Int64("id"),
field.String("username"),
field.String("password"),
field.String("name"),
}
}
// 生成Ent代码
go generate ./ent
```


### 3.3 实现业务逻辑

```
// internal/logic/user/create_user_logic.go
func (l CreateUserLogic) CreateUser(req types.CreateUserRequest) (types.UserInfo, error) {
// 1. 参数校验
if err := l.validateRequest(req); err != nil {
return nil, err
}
// 2. 创建用户
user, err := l.svcCtx.DB.User.Create().
SetUsername(req.Username).
SetPassword(hash(req.Password)).
SetName(req.Name).
Save(l.ctx)
if err != nil {
return nil, errorx.NewDBError(err)
}
// 3. 清除缓存
l.svcCtx.Cache.Del(fmt.Sprintf("user:%d", user.ID))
// 4. 返回结果
return &types.UserInfo{
Id: user.ID,
Username: user.Username,
Name: user.Name,
}, nil
}
```


### 3.4 编写测试


```
// internal/logic/user/create_user_logic_test.go
func TestCreateUserLogic(t testing.T) {
// 1. 准备测试数据
req := &types.CreateUserRequest{
Username: "test_user",
Password: "password123",
Name: "Test User",
}
// 2. 创建测试上下文
ctx := context.Background()
svcCtx := &svc.ServiceContext{
DB: newTestDB(t),
Cache: newTestCache(t),
}
// 3. 执行测试
logic := NewCreateUserLogic(ctx, svcCtx)
resp, err := logic.CreateUser(req)
// 4. 验证结果
assert.NoError(t, err)
assert.NotNil(t, resp)
assert.Equal(t, req.Username, resp.Username)
assert.Equal(t, req.Name, resp.Name)
}
```



## 4. 开发规范

### 4.1 代码规范
- 遵循Go官方代码规范
- 使用gofmt格式化代码
- 使用golangci-lint进行代码检查
- 编写单元测试,保持测试覆盖率

### 4.2 提交规范

```
Commit消息格式
<type>(<scope>): <subject>
type类型
feat: 新功能
fix: 修复Bug
docs: 文档变更
style: 代码格式
refactor: 代码重构
test: 测试相关
chore: 构建过程或辅助工具的变动


示例
feat(user): add user creation api
fix(auth): fix token validation
```



### 4.3 分支管理
- master: 主分支,用于生产环境
- develop: 开发分支,用于测试环境
- feature/*: 功能分支
- hotfix/*: 紧急修复分支

## 5. 调试技巧

### 5.1 日志调试

```
// 设置日志级别
logx.SetLevel(logx.InfoLevel)
// 打印日志
logx.Info("processing request", logx.Field("req", req))
logx.Error("failed to create user", logx.Field("error", err))
// 使用上下文日志
logx.WithContext(ctx).Info("user created", logx.Field("id", user.ID))
```


### 5.2 性能分析

```
// 开启pprof
import "net/http/pprof"
go func() {
http.ListenAndServe(":6060", nil)
}()
// 使用go tool pprof分析
go tool pprof http://localhost:6060/debug/pprof/heap
go tool pprof http://localhost:6060/debug/pprof/profile
```


### 5.3 调试工具

```
go
// 使用Delve调试
dlv debug main.go
// 使用air进行热重载
air -c .air.toml
```

