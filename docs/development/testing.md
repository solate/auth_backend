
# 测试规范文档

## 1.1 测试规范
- 测试文件命名：xxx_test.go
- 测试函数命名：Test<Function>
- 表格驱动测试
- 合理使用mock
- 测试覆盖率要求：>80%


### 1.2 测试示例
```go
// internal/logic/user/create_user_logic_test.go
func TestCreateUserLogic(t *testing.T) {
    tests := []struct {
        name    string
        req     *types.CreateUserRequest
        wantErr bool
    }{
        {
            name: "valid user",
            req: &types.CreateUserRequest{
                Username: "testuser",
                Password: "password123",
                Name:     "Test User",
            },
            wantErr: false,
        },
        {
            name: "duplicate username",
            req: &types.CreateUserRequest{
                Username: "existing_user",
                Password: "password123",
                Name:     "Test User",
            },
            wantErr: true,
        },
        {
            name: "invalid password",
            req: &types.CreateUserRequest{
                Username: "testuser",
                Password: "123", // too short
                Name:     "Test User",
            },
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // 准备测试环境
            ctx := context.Background()
            svcCtx := &svc.ServiceContext{
                DB:    newTestDB(t),
                Cache: newTestCache(t),
            }

            // 执行测试
            logic := NewCreateUserLogic(ctx, svcCtx)
            got, err := logic.CreateUser(tt.req)

            // 验证结果
            if tt.wantErr {
                assert.Error(t, err)
                assert.Nil(t, got)
            } else {
                assert.NoError(t, err)
                assert.NotNil(t, got)
                assert.Equal(t, tt.req.Username, got.Username)
                assert.Equal(t, tt.req.Name, got.Name)
            }
        })
    }
}
```

### 1.3 Mock示例
```go
// internal/logic/user/mock_user_repo.go
type MockUserRepo struct {
    mock.Mock
}

func (m *MockUserRepo) Create(ctx context.Context, user *ent.User) (*ent.User, error) {
    args := m.Called(ctx, user)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).(*ent.User), args.Error(1)
}

// 使用mock进行测试
func TestCreateUserWithMock(t *testing.T) {
    // 创建mock对象
    mockRepo := new(MockUserRepo)
    
    // 设置期望
    mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*ent.User")).
        Return(&ent.User{
            ID:       1,
            Username: "testuser",
            Name:     "Test User",
        }, nil)

    // 使用mock对象进行测试
    ctx := context.Background()
    logic := &CreateUserLogic{
        ctx:    ctx,
        svcCtx: &svc.ServiceContext{UserRepo: mockRepo},
    }

    req := &types.CreateUserRequest{
        Username: "testuser",
        Password: "password123",
        Name:     "Test User",
    }

    // 执行测试
    resp, err := logic.CreateUser(req)

    // 验证结果
    assert.NoError(t, err)
    assert.NotNil(t, resp)
    assert.Equal(t, "testuser", resp.Username)

    // 验证mock调用
    mockRepo.AssertExpectations(t)
}
```


## 2. 集成测试


### 2.1 测试环境
```go
// tests/setup.go
type TestEnv struct {
    DB    *ent.Client
    Redis *redis.Client
    API   *httptest.Server
}

func SetupTestEnv() (*TestEnv, error) {
    // 初始化测试数据库
    db, err := ent.Open("mysql", "root:root@tcp(localhost:3306)/auth_test?parseTime=true")
    if err != nil {
        return nil, err
    }

    // 初始化Redis
    rdb := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
        DB:   1, // 使用单独的数据库
    })

    // 初始化API服务
    handler := setupAPIHandler(db, rdb)
    server := httptest.NewServer(handler)

    return &TestEnv{
        DB:    db,
        Redis: rdb,
        API:   server,
    }, nil
}

func (env *TestEnv) Cleanup() {
    env.DB.Close()
    env.Redis.Close()
    env.API.Close()
}
```

### 2.2 API测试
```go
// tests/api/user_test.go
func TestUserAPI(t *testing.T) {
    // 设置测试环境
    env, err := SetupTestEnv()
    require.NoError(t, err)
    defer env.Cleanup()

    // 测试创建用户
    t.Run("create user", func(t *testing.T) {
        resp, err := http.Post(
            env.API.URL+"/api/v1/users",
            "application/json",
            strings.NewReader(`{
                "username": "testuser",
                "password": "password123",
                "name": "Test User"
            }`),
        )
        require.NoError(t, err)
        assert.Equal(t, http.StatusOK, resp.StatusCode)

        var result struct {
            Code int             `json:"code"`
            Data types.UserInfo `json:"data"`
        }
        err = json.NewDecoder(resp.Body).Decode(&result)
        require.NoError(t, err)
        assert.Equal(t, 0, result.Code)
        assert.Equal(t, "testuser", result.Data.Username)
    })

    // 测试用户登录
    t.Run("user login", func(t *testing.T) {
        resp, err := http.Post(
            env.API.URL+"/api/v1/auth/login",
            "application/json",
            strings.NewReader(`{
                "username": "testuser",
                "password": "password123"
            }`),
        )
        require.NoError(t, err)
        assert.Equal(t, http.StatusOK, resp.StatusCode)

        var result struct {
            Code int                `json:"code"`
            Data types.LoginResponse `json:"data"`
        }
        err = json.NewDecoder(resp.Body).Decode(&result)
        require.NoError(t, err)
        assert.Equal(t, 0, result.Code)
        assert.NotEmpty(t, result.Data.AccessToken)
    })
}
```



## 3. 性能测试


### 3.1 基准测试
```go
// tests/benchmark/auth_test.go
func BenchmarkLogin(b *testing.B) {
    env, err := SetupTestEnv()
    require.NoError(b, err)
    defer env.Cleanup()

    // 准备测试数据
    user := &types.CreateUserRequest{
        Username: "benchmark_user",
        Password: "password123",
        Name:     "Benchmark User",
    }
    _, err = createUser(env.API.URL, user)
    require.NoError(b, err)

    // 执行基准测试
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        resp, err := http.Post(
            env.API.URL+"/api/v1/auth/login",
            "application/json",
            strings.NewReader(`{
                "username": "benchmark_user",
                "password": "password123"
            }`),
        )
        require.NoError(b, err)
        require.Equal(b, http.StatusOK, resp.StatusCode)
        resp.Body.Close()
    }
}
```


### 3.2 负载测试
```go
// tests/load/main.go
func main() {
    // 设置测试参数
    rate := vegeta.Rate{Freq: 100, Per: time.Second} // 每秒100请求
    duration := 5 * time.Minute
    targeter := vegeta.NewStaticTargeter(vegeta.Target{
        Method: "POST",
        URL:    "http://localhost:8888/api/v1/auth/login",
        Body:   []byte(`{"username":"test","password":"test123"}`),
    })

    // 执行测试
    attacker := vegeta.NewAttacker()
    var metrics vegeta.Metrics
    for res := range attacker.Attack(targeter, rate, duration, "Load Test") {
        metrics.Add(res)
    }
    metrics.Close()

    // 输出结果
    fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
    fmt.Printf("95th percentile: %s\n", metrics.Latencies.P95)
    fmt.Printf("Mean: %s\n", metrics.Latencies.Mean)
    fmt.Printf("Max: %s\n", metrics.Latencies.Max)
    fmt.Printf("Requests: %d\n", metrics.Requests)
    fmt.Printf("Success rate: %.2f%%\n", metrics.Success*100)
}
```

## 4. 测试工具


### 4.1 测试框架
- testing: Go标准测试包
- testify: 断言和mock工具
- gomock: 接口mock工具
- httptest: HTTP测试工具
- vegeta: 负载测试工具


### 4.2 测试覆盖率
```bash
# 运行测试并生成覆盖率报告
go test -coverprofile=coverage.out ./...

# 查看覆盖率报告
go tool cover -html=coverage.out

# 检查覆盖率是否达标
go test -cover ./... | grep -v "no test files" | awk '{if($5 != "100.0%") print $0}'
```


### 4.3 测试脚本
```bash
#!/bin/bash
# scripts/test.sh

# 运行单元测试
echo "Running unit tests..."
go test -v ./internal/...

# 运行集成测试
echo "Running integration tests..."
go test -v ./tests/api/...

# 运行基准测试
echo "Running benchmark tests..."
go test -v -bench=. ./tests/benchmark/...

# 生成测试覆盖率报告
echo "Generating coverage report..."
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html

# 检查测试覆盖率
coverage=$(go tool cover -func=coverage.out | grep total | awk '{print $3}')
if (( $(echo "$coverage < 80.0" | bc -l) )); then
    echo "Test coverage is below 80%: $coverage"
    exit 1
fi

echo "All tests passed with coverage: $coverage"
```
