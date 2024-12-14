# 系统详细设计文档

## 1. 认证服务设计

### 1.1 认证流程 

```

// auth.api
type (
LoginRequest {
Username string json:"username"
Password string json:"password"
}
LoginResponse {
AccessToken string json:"accessToken"
RefreshToken string json:"refreshToken"
ExpireTime int64 json:"expireTime"
}
)
service auth-api {
@handler login
post /api/v1/auth/login (LoginRequest) returns (LoginResponse)
}
```


### 1.2 Token管理

```
// internal/logic/auth/token.go
type TokenLogic struct {
ctx context.Context
svcCtx svc.ServiceContext
}
func (l TokenLogic) GenerateToken(uid int64) (types.LoginResponse, error) {
// 生成访问令牌
accessToken, err := l.generateAccessToken(uid)
if err != nil {
return nil, err
}
// 生成刷新令牌
refreshToken, err := l.generateRefreshToken(uid)
if err != nil {
return nil, err
}
return &types.LoginResponse{
AccessToken: accessToken,
RefreshToken: refreshToken,
ExpireTime: time.Now().Add(24 time.Hour).Unix(),
}, nil
}
```


## 2. 用户服务设计

### 2.1 用户模型


```
// ent/schema/user.go
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
field.String("username").Unique(),
field.String("password"),
field.String("name"),
field.String("email"),
field.String("phone"),
field.Time("created_at"),
field.Time("updated_at"),
field.Bool("is_deleted").Default(false),
}
}
func (User) Edges() []ent.Edge {
return []ent.Edge{
// 定义与角色的关系
edge.To("roles", Role.Type),
// 定义与组织的关系
edge.To("organizations", Organization.Type),
}
}

```



### 2.2 用户管理接口

```
// user.api
type (
CreateUserRequest {
Username string json:"username"
Password string json:"password"
Name string json:"name"
Email string json:"email"
Phone string json:"phone"
}
UserInfo {
Id int64 json:"id"
Username string json:"username"
Name string json:"name"
Email string json:"email"
Phone string json:"phone"
}
)
service user-api {
@handler createUser
post /api/v1/users (CreateUserRequest) returns (UserInfo)
@handler getUserById
get /api/v1/users/:id returns (UserInfo)
}
```


## 3. 权限服务设计

### 3.1 RBAC模型设计

```
// ent/schema/role.go
type Role struct {
ent.Schema
}
func (Role) Fields() []ent.Field {
return []ent.Field{
field.Int64("id"),
field.String("name").Unique(),
field.String("code").Unique(),
field.String("description"),
field.Time("created_at"),
field.Time("updated_at"),
}
}
func (Role) Edges() []ent.Edge {
return []ent.Edge{
edge.From("users", User.Type).Ref("roles"),
edge.To("permissions", Permission.Type),
}
}
// ent/schema/permission.go
type Permission struct {
ent.Schema
}
func (Permission) Fields() []ent.Field {
return []ent.Field{
field.Int64("id"),
field.String("name"),
field.String("code").Unique(),
field.String("type"), // API, MENU, BUTTON等
field.String("description"),
field.Time("created_at"),
field.Time("updated_at"),
}
}
```


### 3.2 权限检查中间件

```
// pkg/middleware/auth.go
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
return func(w http.ResponseWriter, r http.Request) {
// 1. 获取并验证token
token := r.Header.Get("Authorization")
claims, err := parseToken(token)
if err != nil {
httpx.Error(w, errorx.NewAuthError("invalid token"))
return
}
// 2. 检查权限
if !checkPermission(claims.UserId, r.URL.Path, r.Method) {
httpx.Error(w, errorx.NewAuthError("permission denied"))
return
}
next(w, r)
}
}
```


## 4. 缓存设计

### 4.1 缓存模型

```
// model/cache/user.go
type UserCache struct {
redis redis.Client
prefix string
}
func NewUserCache(redis redis.Client) UserCache {
return &UserCache{
redis: redis,
prefix: "cache:user:",
}
}
func (c UserCache) GetUser(id int64) (ent.User, error) {
key := fmt.Sprintf("%s%d", c.prefix, id)
// 1. 从缓存获取
data, err := c.redis.Get(context.Background(), key).Bytes()
if err == nil {
var user ent.User
if err := json.Unmarshal(data, &user); err == nil {
return &user, nil
}
}
// 2. 从数据库获取
user, err := entClient.User.Get(context.Background(), id)
if err != nil {
return nil, err
}
// 3. 写入缓存
if data, err := json.Marshal(user); err == nil {
c.redis.Set(context.Background(), key, data, 24time.Hour)
}
return user, nil
}
```


## 5. 审计日志设计

### 5.1 审计日志模型


```
// ent/schema/audit_log.go
type AuditLog struct {
ent.Schema
}
func (AuditLog) Fields() []ent.Field {
return []ent.Field{
field.Int64("id"),
field.Int64("user_id"),
field.String("action"),
field.String("resource_type"),
field.Int64("resource_id"),
field.String("detail"),
field.String("ip"),
field.String("user_agent"),
field.Time("created_at"),
}
}
```


### 5.2 审计日志中间件

```
// pkg/middleware/audit.go
func AuditMiddleware(next http.HandlerFunc) http.HandlerFunc {
return func(w http.ResponseWriter, r http.Request) {
// 1. 获取用户信息
userId := getUserFromContext(r.Context())
// 2. 记录请求信息
log := &ent.AuditLog{
UserId: userId,
Action: r.Method,
ResourceType: getResourceType(r.URL.Path),
IP: getClientIP(r),
UserAgent: r.UserAgent(),
CreatedAt: time.Now(),
}
// 3. 异步写入审计日志
go func() {
if err := createAuditLog(log); err != nil {
logx.Error("failed to create audit log:", err)
}
}()
next(w, r)
}
}
```


## 6. 服务监控设计

### 6.1 Prometheus指标


```
// pkg/metrics/metrics.go
var (
RequestCounter = prometheus.NewCounterVec(
prometheus.CounterOpts{
Name: "http_requests_total",
Help: "Total number of HTTP requests",
},
[]string{"method", "path", "status"},
)
RequestDuration = prometheus.NewHistogramVec(
prometheus.HistogramOpts{
Name: "http_request_duration_seconds",
Help: "HTTP request duration in seconds",
Buckets: []float64{0.1, 0.3, 0.5, 0.7, 0.9, 1.0},
},
[]string{"method", "path"},
)
)
```


### 6.2 监控中间件

```
// pkg/middleware/metrics.go
func MetricsMiddleware(next http.HandlerFunc) http.HandlerFunc {
return func(w http.ResponseWriter, r http.Request) {
start := time.Now()
// 包装ResponseWriter以获取状态码
ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
next(ww, r)
// 记录请求数量
RequestCounter.WithLabelValues(
r.Method,
r.URL.Path,
strconv.Itoa(ww.Status()),
).Inc()
// 记录请求耗时
RequestDuration.WithLabelValues(
r.Method,
r.URL.Path,
).Observe(time.Since(start).Seconds())
}
}
```

# 数据库设计文档

## 1. 数据库设计原则

### 1.1 基本原则
- 使用ent作为ORM框架
- 统一字段命名规范
- 合理使用索引
- 考虑分库分表扩展
- 软删除而非物理删除

### 1.2 基础字段规范
所有表都应包含以下基础字段:


```
// ent/schema/mixin/base.go
package mixin
import (
"time"
"entgo.io/ent"
"entgo.io/ent/schema/field"
"entgo.io/ent/schema/mixin"
)
type BaseMixin struct {
mixin.Schema
}
func (BaseMixin) Fields() []ent.Field {
return []ent.Field{
field.Int64("id").
Positive().
Immutable().
Comment("主键ID"),
field.Time("created_at").
Immutable().
Default(time.Now).
Comment("创建时间"),
field.Time("updated_at").
Default(time.Now).
UpdateDefault(time.Now).
Comment("更新时间"),
field.Bool("is_deleted").
Default(false).
Comment("是否删除"),
}
}

```


## 2. 核心表结构设计

### 2.1 用户表(users)


```
// ent/schema/user.go
type User struct {
ent.Schema
}
func (User) Mixin() []ent.Mixin {
return []ent.Mixin{
mixin.BaseMixin{},
}
}
func (User) Fields() []ent.Field {
return []ent.Field{
field.String("username").
Unique().
NotEmpty().
Comment("用户名"),
field.String("password").
Sensitive().
NotEmpty().
Comment("密码"),
field.String("name").
NotEmpty().
Comment("姓名"),
field.String("email").
Optional().
Comment("邮箱"),
field.String("phone").
Optional().
Comment("手机号"),
field.Enum("status").
Values("active", "inactive", "locked").
Default("active").
Comment("用户状态"),
field.Time("last_login_at").
Optional().
Comment("最后登录时间"),
field.String("last_login_ip").
Optional().
Comment("最后登录IP"),
}
}
func (User) Edges() []ent.Edge {
return []ent.Edge{
edge.To("roles", Role.Type).
Comment("用户角色"),
edge.To("organizations", Organization.Type).
Comment("所属组织"),
}
}
func (User) Indexes() []ent.Index {
return []ent.Index{
index.Fields("username"),
index.Fields("phone"),
index.Fields("email"),
}
}
```

### 2.2 角色表(roles)


```
// ent/schema/role.go
type Role struct {
ent.Schema
}
func (Role) Mixin() []ent.Mixin {
return []ent.Mixin{
mixin.BaseMixin{},
}
}
func (Role) Fields() []ent.Field {
return []ent.Field{
field.String("name").
NotEmpty().
Comment("角色名称"),
field.String("code").
Unique().
NotEmpty().
Comment("角色编码"),
field.String("description").
Optional().
Comment("角色描述"),
field.Bool("is_system").
Default(false).
Comment("是否系统角色"),
}
}
func (Role) Edges() []ent.Edge {
return []ent.Edge{
edge.From("users", User.Type).
Ref("roles"),
edge.To("permissions", Permission.Type).
Through("role_permissions", RolePermission.Type),
}
}

```

### 2.3 权限表(permissions)

```
// ent/schema/permission.go
type Permission struct {
ent.Schema
}
func (Permission) Mixin() []ent.Mixin {
return []ent.Mixin{
mixin.BaseMixin{},
}
}
func (Permission) Fields() []ent.Field {
return []ent.Field{
field.String("name").
NotEmpty().
Comment("权限名称"),
field.String("code").
Unique().
NotEmpty().
Comment("权限编码"),
field.String("type").
NotEmpty().
Comment("权限类型(API/MENU/BUTTON)"),
field.String("parent_id").
Optional().
Comment("父级权限ID"),
field.String("path").
Optional().
Comment("资源路径"),
field.String("method").
Optional().
Comment("HTTP方法"),
field.String("description").
Optional().
Comment("权限描述"),
}
}
func (Permission) Edges() []ent.Edge {
return []ent.Edge{
edge.From("roles", Role.Type).
Ref("permissions").
Through("role_permissions", RolePermission.Type),
}
}
func (Permission) Indexes() []ent.Index {
return []ent.Index{
index.Fields("code"),
index.Fields("type", "path", "method"),
}
}
```


### 2.4 组织表(organizations)

```
// ent/schema/organization.go
type Organization struct {
ent.Schema
}
func (Organization) Mixin() []ent.Mixin {
return []ent.Mixin{
mixin.BaseMixin{},
}
}
func (Organization) Fields() []ent.Field {
return []ent.Field{
field.String("name").
NotEmpty().
Comment("组织名称"),
field.String("code").
Unique().
NotEmpty().
Comment("组织编码"),
field.Int64("parent_id").
Optional().
Comment("父级组织ID"),
field.String("path").
NotEmpty().
Comment("组织路径"),
field.Int("sort").
Default(0).
Comment("排序"),
field.String("leader").
Optional().
Comment("负责人"),
field.String("description").
Optional().
Comment("描述"),
}
}
func (Organization) Edges() []ent.Edge {
return []ent.Edge{
edge.From("users", User.Type).
Ref("organizations"),
edge.To("children", Organization.Type).
From("parent").
Unique(),
}
}
func (Organization) Indexes() []ent.Index {
return []ent.Index{
index.Fields("code"),
index.Fields("parent_id"),
index.Fields("path"),
}
}
```


## 3. 关联表设计

### 3.1 角色权限关联表(role_permissions)

```
// ent/schema/role_permission.go
type RolePermission struct {
ent.Schema
}
func (RolePermission) Mixin() []ent.Mixin {
return []ent.Mixin{
mixin.BaseMixin{},
}
}
func (RolePermission) Fields() []ent.Field {
return []ent.Field{
field.Int64("role_id").
Comment("角色ID"),
field.Int64("permission_id").
Comment("权限ID"),
}
}
func (RolePermission) Edges() []ent.Edge {
return []ent.Edge{
edge.To("role", Role.Type).
Unique().
Required().
Field("role_id"),
edge.To("permission", Permission.Type).
Unique().
Required().
Field("permission_id"),
}
}
func (RolePermission) Indexes() []ent.Index {
return []ent.Index{
index.Fields("role_id", "permission_id").
Unique(),
}
}
```


## 4. 缓存设计

### 4.1 缓存键设计

```
// model/cache/keys.go
const (
// 用户信息缓存
UserCacheKey = "user:%d"
// 用户权限缓存
UserPermissionsCacheKey = "user:permissions:%d"
// 角色权限缓存
RolePermissionsCacheKey = "role:permissions:%d"
// 组织架构缓存
OrganizationTreeCacheKey = "org:tree"
)
// 缓存过期时间
const (
UserCacheExpiration = 24 time.Hour
PermissionCacheExpiration = 12 time.Hour
OrganizationCacheExpiration = 24 time.Hour
)
```


### 4.2 缓存接口设计

```
// model/cache/interface.go
type Cache interface {
Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
Get(ctx context.Context, key string, value interface{}) error
Del(ctx context.Context, key string) error
DelByPattern(ctx context.Context, pattern string) error
}
```

## 5. 数据库操作规范

### 5.1 事务处理

```
// 使用ent的事务
func CreateUserWithRoles(ctx context.Context, client ent.Client, user ent.User, roleIDs []int) error {
return client.WithTx(ctx, func(tx ent.Tx) error {
// 创建用户
u, err := tx.User.Create().
SetUsername(user.Username).
SetPassword(user.Password).
SetName(user.Name).
Save(ctx)
if err != nil {
return fmt.Errorf("failed creating user: %w", err)
}
// 关联角色
roles, err := tx.Role.Query().
Where(role.IDIn(roleIDs...)).
All(ctx)
if err != nil {
return fmt.Errorf("failed querying roles: %w", err)
}
return u.AddRoles(ctx, roles...)
})
}

```



### 5.2 批量操作

```
// 批量创建用户
func BatchCreateUsers(ctx context.Context, client ent.Client, users []ent.User) error {
bulk := make([]ent.UserCreate, len(users))
for i, user := range users {
bulk[i] = client.User.Create().
SetUsername(user.Username).
SetPassword(user.Password).
SetName(user.Name)
}
return client.User.CreateBulk(bulk...).Exec(ctx)
}
```



### 5.3 软删除实现

```
// 软删除中间件
func (BaseMixin) Hooks() []ent.Hook {
return []ent.Hook{
// 查询时自动过滤已删除记录
hooks.On(
func(next ent.Querier) ent.Querier {
return ent.QuerierFunc(func(ctx context.Context, q ent.Query) (ent.Value, error) {
q.Where(predicate.And(
predicate.Bool("is_deleted", false),
))
return next.Query(ctx, q)
})
},
ent.OpRead|ent.OpDelete|ent.OpUpdate,
),
}
}
```

