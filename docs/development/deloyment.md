# 部署文档

## 1. 部署架构

### 1.1 系统组件

- API服务
- MySQL数据库
- Redis缓存
- Prometheus监控
- Grafana可视化
- ELK日志系统


### 1.2 部署拓扑
```ascii
                        [负载均衡]
                            |
                    +--------------+
                    |              |
                [API服务集群]  [API服务集群]
                    |              |
            +-------+------+-------+------+
            |              |              |
        [MySQL主从]    [Redis集群]    [监控系统]
```

## 2. 环境要求


### 2.1 硬件要求

- CPU: 4核+
- 内存: 8GB+
- 磁盘: 100GB+
- 网络: 千兆网卡

### 2.2 软件要求

- Docker 20.10+
- Docker Compose 2.0+
- Kubernetes 1.20+ (可选)


## 3. 部署步骤


### 3.1 Docker Compose部署
```yaml
# deploy/docker-compose.yml
version: '3'

services:
  api:
    build: 
      context: .
      dockerfile: Dockerfile
    ports:
      - "8888:8888"
    environment:
      - MYSQL_HOST=mysql
      - REDIS_HOST=redis
    depends_on:
      - mysql
      - redis
    restart: always

  mysql:
    image: mysql:8.0
    ports:
      - "3306:3306"
    environment:
      - MYSQL_ROOT_PASSWORD=root
      - MYSQL_DATABASE=auth_system
    volumes:
      - mysql_data:/var/lib/mysql
    restart: always

  redis:
    image: redis:7.0
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data
    restart: always

  prometheus:
    image: prom/prometheus
    ports:
      - "9090:9090"
    volumes:
      - ./prometheus.yml:/etc/prometheus/prometheus.yml
    restart: always

  grafana:
    image: grafana/grafana
    ports:
      - "3000:3000"
    depends_on:
      - prometheus
    restart: always

volumes:
  mysql_data:
  redis_data:
```


### 3.2 Kubernetes部署
```yaml
# deploy/kubernetes/api-deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: auth-api
spec:
  replicas: 3
  selector:
    matchLabels:
      app: auth-api
  template:
    metadata:
      labels:
        app: auth-api
    spec:
      containers:
      - name: auth-api
        image: auth-api:latest
        ports:
        - containerPort: 8888
        env:
        - name: MYSQL_HOST
          valueFrom:
            configMapKeyRef:
              name: auth-config
              key: mysql_host
        - name: REDIS_HOST
          valueFrom:
            configMapKeyRef:
              name: auth-config
              key: redis_host
        resources:
          requests:
            cpu: 200m
            memory: 256Mi
          limits:
            cpu: 500m
            memory: 512Mi
        livenessProbe:
          httpGet:
            path: /health
            port: 8888
          initialDelaySeconds: 30
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /ready
            port: 8888
          initialDelaySeconds: 5
          periodSeconds: 5
```


## 4. 配置管理


### 4.1 配置文件
```yaml
# configs/config.yaml
Name: auth-api
Host: 0.0.0.0
Port: 8888

MySQL:
  Host: ${MYSQL_HOST}
  Port: 3306
  Database: auth_system
  Username: ${MYSQL_USER}
  Password: ${MYSQL_PASSWORD}
  
Redis:
  Host: ${REDIS_HOST}
  Port: 6379
  Password: ${REDIS_PASSWORD}
  
Auth:
  AccessSecret: ${JWT_SECRET}
  AccessExpire: 7200
  
Log:
  ServiceName: auth-api
  Level: info
  Mode: file
  Path: /var/log/auth-api
```

### 4.2 环境变量
```bash
# deploy/.env
MYSQL_HOST=localhost
MYSQL_USER=root
MYSQL_PASSWORD=root
REDIS_HOST=localhost
REDIS_PASSWORD=
JWT_SECRET=your-secret-key
```

## 5. 监控告警


### 5.1 Prometheus配置
```yaml
# deploy/prometheus.yml
global:
  scrape_interval: 15s
  evaluation_interval: 15s

scrape_configs:
  - job_name: 'auth-api'
    static_configs:
      - targets: ['auth-api:8888']
```


### 5.2 Grafana仪表板
```json
{
  "dashboard": {
    "title": "Auth System Dashboard",
    "panels": [
      {
        "title": "API请求量",
        "type": "graph",
        "datasource": "Prometheus",
        "targets": [
          {
            "expr": "rate(http_requests_total[5m])"
          }
        ]
      },
      {
        "title": "响应时间",
        "type": "graph",
        "datasource": "Prometheus",
        "targets": [
          {
            "expr": "rate(http_request_duration_seconds_sum[5m])/rate(http_request_duration_seconds_count[5m])"
          }
        ]
      }
    ]
  }
}
```


## 6. 运维手册


### 6.1 日常运维
```bash
# 查看服务状态
docker-compose ps
kubectl get pods

# 查看日志
docker-compose logs -f api
kubectl logs -f deployment/auth-api

# 重启服务
docker-compose restart api
kubectl rollout restart deployment auth-api

# 扩容服务
docker-compose up -d --scale api=3
kubectl scale deployment auth-api --replicas=3
```


### 6.2 备份恢复
```bash
# MySQL备份
mysqldump -h localhost -u root -p auth_system > backup.sql

# MySQL恢复
mysql -h localhost -u root -p auth_system < backup.sql

# Redis备份
redis-cli save

# Redis恢复
cp dump.rdb /data/
docker-compose restart redis
```



### 6.3 故障处理
```bash
# 1. 服务不可用
- 检查服务状态和日志
- 检查数据库连接
- 检查Redis连接
- 检查磁盘空间

# 2. 性能问题
- 检查CPU和内存使用
- 检查慢查询日志
- 检查连接池配置
- 检查缓存命中率

# 3. 数据问题
- 检查数据一致性
- 执行数据修复脚本
- 恢复数据备份
```
