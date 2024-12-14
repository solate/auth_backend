# API设计规范

## 1. API定义规范

### 1.1 基本规范

- 使用go-zero的api文件定义接口
- 遵循RESTful设计原则 （改为全Post请求）
- 版本号放在URL中 （考虑考虑使用Header）
- 使用蛇形命名(snake_case)
- 统一响应格式

### 1.2 API文件结构

```
// api/auth/auth.api
syntax = "v1"

info(
    title: "认证服务接口"
    desc: "提供用户认证相关接口"
    author: "team"
    email: "team@example.com"
    version: "v1"
)

type (
    LoginRequest {
        Username string `json:"username" validate:"required,min=3,max=32"`
        Password string `json:"password" validate:"required,min=6,max=32"`
    }

    LoginResponse {
        AccessToken  string `json:"access_token"`
        TokenType   string `json:"token_type"`
        ExpiresIn   int64  `json:"expires_in"`
        RefreshToken string `json:"refresh_token,optional"`
    }
)

@server(
    jwt: Auth
    group: auth
    middleware: Cors, Metrics
)

service auth-api {
    @doc "用户登录"
    @handler login
    post /api/v1/auth/login (LoginRequest) returns (LoginResponse)
}
```


## 2. 请求规范

### 2.1 HTTP方法使用

- GET: 查询资源
- POST: 创建资源
- PUT: 更新资源(全量更新)
- PATCH: 更新资源(部分更新)
- DELETE: 删除资源


### 2.2 URL规范
```
# 资源集合
GET    /api/v1/users           # 获取用户列表
POST   /api/v1/users           # 创建用户

# 具体资源
GET    /api/v1/users/:id       # 获取指定用户
PUT    /api/v1/users/:id       # 更新指定用户
DELETE /api/v1/users/:id       # 删除指定用户

# 子资源
GET    /api/v1/users/:id/roles # 获取用户角色
POST   /api/v1/users/:id/roles # 分配用户角色

# 操作
POST   /api/v1/auth/login      # 用户登录
POST   /api/v1/auth/logout     # 用户登出
```



### 2.3 请求参数规范
```go
type (
    // 查询参数
    ListUserRequest {
        Page     int    `form:"page,default=1"`
        PageSize int    `form:"page_size,default=20"`
        Keyword  string `form:"keyword,optional"`
        Status   string `form:"status,optional"`
        OrderBy  string `form:"order_by,optional"`
    }

    // 创建参数
    CreateUserRequest {
        Username string   `json:"username" validate:"required,min=3,max=32"`
        Password string   `json:"password" validate:"required,min=6,max=32"`
        Name     string   `json:"name" validate:"required"`
        Email    string   `json:"email" validate:"required,email"`
        Phone    string   `json:"phone" validate:"required,phone"`
        RoleIds  []int64  `json:"role_ids" validate:"required,min=1"`
    }
)
```

## 3. 响应规范

### 3.1 统一响应格式
```go
// pkg/response/response.go
type Response struct {
    Code    int         `json:"code"`              // 业务码
    Message string      `json:"message"`           // 提示信息
    Data    interface{} `json:"data,omitempty"`    // 数据
    TraceId string      `json:"trace_id,omitempty"` // 链路ID
}

// 分页响应
type PageResponse struct {
    List     interface{} `json:"list"`       // 数据列表
    Total    int64      `json:"total"`      // 总数
    Page     int        `json:"page"`       // 当前页
    PageSize int        `json:"page_size"`  // 每页大小
}
```

### 3.2 错误码规范
```go
// pkg/errorx/code.go
const (
    Success = 0    // 成功
    
    // 客户端错误
    InvalidParams  = 40001  // 参数错误
    Unauthorized   = 40101  // 未授权
    TokenExpired   = 40102  // 令牌过期
    Forbidden      = 40301  // 禁止访问
    NotFound       = 40401  // 资源不存在
    
    // 服务端错误
    InternalError  = 50001  // 内部错误
    DBError       = 50002  // 数据库错误
    CacheError    = 50003  // 缓存错误
    RpcError      = 50004  // RPC调用错误
)
```


## 4. 接口文档


### 4.1 Swagger配置
```go
// api/auth/auth.api
@doc(
    summary: "用户登录接口"
    description: "通过用户名密码进行登录认证"
    tags: ["auth"]
)
```



### 4.2 接口注释规范
```go
// internal/handler/auth/login.go

// Login 用户登录
// @Summary 用户登录接口
// @Description 通过用户名密码进行登录认证
// @Tags auth
// @Accept json
// @Produce json
// @Param data body LoginRequest true "登录参数"
// @Success 200 {object} Response{data=LoginResponse} "成功"
// @Failure 400 {object} Response "参数错误"
// @Failure 401 {object} Response "认证失败"
// @Router /api/v1/auth/login [post]
```

## 5. 安全规范


### 5.1 认证中间件
```go
// pkg/middleware/auth.go
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // 1. 获取token
        token := r.Header.Get("Authorization")
        if token == "" {
            httpx.Error(w, errorx.NewAuthError("missing token"))
            return
        }

        // 2. 验证token
        claims, err := jwt.ParseToken(token)
        if err != nil {
            httpx.Error(w, errorx.NewAuthError("invalid token"))
            return
        }

        // 3. 将用户信息写入上下文
        ctx := context.WithValue(r.Context(), "user_id", claims.UserId)
        next(w, r.WithContext(ctx))
    }
}
```


### 5.2 权限中间件
```go
// pkg/middleware/permission.go
func PermissionMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        userId := r.Context().Value("user_id").(int64)
        
        // 检查用户是否有权限访问当前接口
        if !checkPermission(userId, r.URL.Path, r.Method) {
            httpx.Error(w, errorx.NewForbiddenError("permission denied"))
            return
        }
        
        next(w, r)
    }
}
```


## 6. 性能优化


### 6.1 接口限流
```go
// pkg/middleware/ratelimit.go
func RateLimitMiddleware(limit float64) func(http.HandlerFunc) http.HandlerFunc {
    limiter := rate.NewLimiter(rate.Limit(limit), int(limit))
    return func(next http.HandlerFunc) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {
            if !limiter.Allow() {
                httpx.Error(w, errorx.NewTooManyRequestsError())
                return
            }
            next(w, r)
        }
    }
}
```

### 6.2 响应压缩
```go
// pkg/middleware/gzip.go
func GzipMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
            next(w, r)
            return
        }
        
        gz := gzip.NewWriter(w)
        defer gz.Close()
        
        w.Header().Set("Content-Encoding", "gzip")
        next(gzipResponseWriter{Writer: gz, ResponseWriter: w}, r)
    }
}
```
```

