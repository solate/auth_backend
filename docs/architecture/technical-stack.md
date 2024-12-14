
docs/
├── architecture/              # 架构设计文档
│   ├── overview.md           # 架构概述
│   ├── system-design.md      # 系统设计
│   ├── technical-stack.md    # 技术栈选型
│   ├── security.md           # 安全架构
│   └── deployment.md         # 部署架构
├── development/              # 开发文档
│   ├── guide.md             # 开发指南
│   ├── standards.md         # 开发规范
│   ├── api-design.md        # API设计规范
│   └── database.md          # 数据库设计
└── README.md                # 文档说明 





# 技术栈选型文档

## 1. 核心技术栈

### 1.1 后端技术栈
- 开发语言: Go 1.21+
- 微服务框架: go-zero
- ORM框架: ent
- 数据库: MySQL 8.0+
- 缓存: Redis 7.0+

### 1.2 基础组件
- 配置中心: go-zero内置配置管理
- 服务发现: etcd
- 消息队列: Redis Stream (轻量级消息队列)
- 日志收集: go-zero内置日志 + filebeat
- 链路追踪: go-zero内置trace
- 监控告警: prometheus + grafana

## 2. 技术选型理由

### 2.1 go-zero框架
- 高性能微服务框架
- 内置服务治理能力
- 自动生成API代码
- 丰富的中间件支持
- 完善的监控指标

### 2.2 ent框架
- 类型安全的ORM
- 强大的图查询能力
- 自动生成CRUD代码
- Schema即代码
- 事务支持完善

### 2.3 自研组件
- 认证中心: 基于go-zero自研
- 权限引擎: 自研RBAC+ABAC引擎
- 审计日志: 基于ent hooks实现
- 数据权限: 基于ent interceptor实现

## 3. 系统分层

### 3.1 API层 (api)
- 接口定义: api/*.api
- 协议: HTTP/gRPC
- 参数校验: go-zero validator
- 接口文档: swagger

### 3.2 业务层 (service)
- 业务逻辑
- 事务处理
- 领域模型
- 业务规则

### 3.3 数据层 (model)
- ent Schema定义
- 数据模型
- 数据访问
- 缓存处理

## 4. 代码结构

```
project/
├── api/ # API定义和服务入口
│ ├── auth/ # 认证服务API
│ ├── user/ # 用户服务API
│ └── permission/ # 权限服务API
├── service/ # 微服务实现
│ ├── auth/ # 认证服务
│ ├── user/ # 用户服务
│ └── permission/ # 权限服务
├── model/ # 数据模型
│ ├── ent/ # ent实体定义
│ └── cache/ # 缓存模型
├── pkg/ # 公共包
│ ├── errorx/ # 错误处理
│ ├── middleware/ # 中间件
│ └── utils/ # 工具函数
└── deploy/ # 部署配置
├── kubernetes/ # k8s配置
└── docker/ # docker配置
```


## 5. 开发规范

### 5.1 API设计规范
- 使用go-zero api文件定义接口
- RESTful风格API
- 统一错误码
- 统一响应格式

### 5.2 数据库规范
- 使用ent管理数据库Schema
- 统一字段命名
- 必备基础字段
- 索引设计规范

### 5.3 缓存规范
- 统一缓存键前缀
- 合理的过期时间
- 缓存穿透防护
- 缓存击穿防护

## 6. 性能优化

### 6.1 API层优化
- 接口限流
- 请求合并
- 响应压缩
- 连接复用

### 6.2 业务层优化
- 本地缓存
- 批量处理
- 异步处理
- 并发控制

### 6.3 数据层优化
- 索引优化
- 分库分表
- 读写分离
- 缓存优化


