# 后端开发指南

## 配置文件

项目使用 YAML 配置文件，支持多环境：

| 文件 | 用途 |
|------|------|
| `configs/config.yaml` | 默认配置 |
| `configs/config.dev.yaml` | 开发环境 |
| `configs/config.prod.yaml` | 生产环境 |

启动时通过 `-config` 参数指定：

```bash
go run cmd/server/main.go -config configs/config.dev.yaml
```

---

## 数据库配置

支持三种数据库驱动：**MySQL**、**PostgreSQL**、**SQLite**，通过 `database.driver` 切换。

### PostgreSQL（当前开发环境）

```yaml
database:
  driver: postgres
  host: localhost
  port: 5432
  username: postgres
  password: "tiger"
  database: infra
  max_idle_conns: 10      # 最大空闲连接数
  max_open_conns: 100     # 最大打开连接数
  conn_max_lifetime: 3600 # 连接最大生命周期（秒）
```

连接 DSN 格式（代码自动拼接）：

```
host=localhost port=5432 user=postgres password=tiger dbname=infra sslmode=disable TimeZone=Asia/Shanghai
```

### MySQL

```yaml
database:
  driver: mysql
  host: localhost
  port: 3306
  username: root
  password: ""
  database: infra
  charset: utf8mb4
  max_idle_conns: 10
  max_open_conns: 100
  conn_max_lifetime: 3600
```

### SQLite

```yaml
database:
  driver: sqlite
  path: data/infra.db      # 数据库文件路径
```

### 驱动对照表

| 驱动       | `database.driver` | 连接器                        |
|------------|-------------------|-------------------------------|
| PostgreSQL | `postgres`        | `gorm.io/driver/postgres`     |
| MySQL      | `mysql`           | `gorm.io/driver/mysql`        |
| SQLite     | `sqlite`          | `gorm.io/driver/sqlite`       |

---

## 服务配置

```yaml
server:
  port: 8080              # 监听端口
  mode: debug             # Gin模式: debug / release / test
  read_timeout: 60        # 读取超时（秒）
  write_timeout: 60       # 写入超时（秒）
```

---

## Redis 配置

```yaml
redis:
  host: localhost
  port: 6379
  password: ""
  db: 0                   # 数据库编号
  pool_size: 10           # 连接池大小
```

---

## JWT 配置

```yaml
jwt:
  secret: infra-jwt-secret-key-dev   # 签名密钥（生产环境需更换）
  expire: 7200                       # Token 过期时间（秒）
```

---

## 日志配置

```yaml
log:
  level: debug            # 级别: debug / info / warn / error
  filename: logs/app.log  # 日志文件路径
  max_size: 100           # 单个日志文件最大大小（MB）
  max_backups: 3          # 最大保留的旧日志文件数
  max_age: 7              # 日志文件保留天数
  compress: true          # 是否压缩旧日志
```

---

## OSS 配置（阿里云对象存储）

```yaml
oss:
  endpoint: oss-cn-hangzhou.aliyuncs.com
  access_key_id: ""
  access_key_secret: ""
  bucket: ""
  base_path: uploads/
```

---

## 快速开始（PostgreSQL）

### 1. 创建数据库

```bash
createdb -U postgres infra
# 或者
psql -U postgres -c "CREATE DATABASE infra;"
```

### 2. 修改配置

编辑 `configs/config.dev.yaml`，确认数据库段如下：

```yaml
database:
  driver: postgres
  host: localhost
  port: 5432
  username: postgres
  password: "tiger"
  database: infra
```

### 3. 启动服务

```bash
cd server
make run
```

服务启动后访问 `http://localhost:8080`，Swagger 文档在 `http://localhost:8080/swagger/index.html`。

---

## 配置热加载

项目使用 Viper + fsnotify，修改配置文件后服务会自动重载，无需重启。
