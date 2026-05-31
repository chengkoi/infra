# Infra

基于 Go + Gin + GORM + Casbin 的模块化单体后端项目。

## 目录结构

```text
infra/
│
├── server/                     # 后端项目
│   ├── cmd/
│   │   └── server/
│   │       └── main.go         # 程序入口
│   │
│   ├── configs/
│   │   ├── config.yaml         # 默认配置
│   │   ├── config.dev.yaml     # 开发环境
│   │   └── config.prod.yaml    # 生产环境
│   │
│   ├── internal/
│   │
│   │   ├── system/             # 系统模块
│   │   │
│   │   │   ├── auth/
│   │   │   │   ├── handler.go  # 登录/注册/忘记密码/重置密码
│   │   │   │   ├── service.go  # 认证业务
│   │   │   │   ├── dto.go      # 请求响应
│   │   │   │   └── route.go    # 路由注册
│   │   │   │
│   │   │   ├── user/
│   │   │   │   ├── handler.go  # CRUD接口
│   │   │   │   ├── service.go  # 用户业务
│   │   │   │   ├── repository.go # 数据访问
│   │   │   │   ├── model.go    # Gorm实体
│   │   │   │   ├── dto.go      # 请求响应对象
│   │   │   │   └── route.go    # 路由注册
│   │   │   │
│   │   │   ├── operlog/        # 操作日志
│   │   │   │   ├── handler.go
│   │   │   │   ├── service.go
│   │   │   │   ├── repository.go
│   │   │   │   ├── model.go
│   │   │   │   ├── dto.go
│   │   │   │   └── route.go
│   │   │   │
│   │   │   ├── loginlog/       # 登录日志
│   │   │   │   ├── handler.go
│   │   │   │   ├── service.go
│   │   │   │   ├── repository.go
│   │   │   │   ├── model.go
│   │   │   │   ├── dto.go
│   │   │   │   └── route.go
│   │   │   │
│   │   │   ├── dept/           # 部门
│   │   │   ├── post/           # 岗位
│   │   │   ├── role/           # 角色
│   │   │   ├── menu/           # 菜单
│   │   │   ├── dict/           # 字典
│   │   │   ├── config/         # 配置
│   │   │   └── notice/         # 通知
│   │   │
│   │   ├── monitor/            # 监控模块
│   │   │   ├── online/         # 在线用户
│   │   │   ├── job/            # 定时任务
│   │   │   ├── cache/          # 缓存监控
│   │   │   ├── datasource/     # 数据源监控
│   │   │   └── server/         # 服务监控
│   │   │
│   │   ├── tool/               # 工具模块
│   │   │
│   │   │   ├── codegen/        # 代码生成
│   │   │
│   │   ├── middleware/         # 中间件
│   │   │   ├── auth.go         # JWT认证
│   │   │   ├── cors.go         # 跨域
│   │   │   ├── logger.go       # 请求日志
│   │   │   └── recovery.go     # panic恢复
│   │   │
│   │   ├── router/
│   │   │   └── router.go       # 总路由入口
│   │   │
│   │   └── shared/             # 共享组件
│   │
│   │       ├── database/
│   │       │   ├── mysql.go    # mysql连接
│   │       │   ├── postgres.go # postgres连接
│   │       │   ├── sqlite.go   # sqlite连接
│   │       │   └── gorm.go     # gorm初始化
│   │       │
│   │       ├── response/
│   │       │   ├── response.go # Success/Fail
│   │       │   └── page.go     # 分页结构
│   │       │
│   │       ├── errno/
│   │       │   └── code.go     # 错误码定义
│   │       │
│   │       ├── logger/
│   │       │   └── logger.go   # zap日志
│   │       │
│   │       ├── jwt/
│   │       │   └── jwt.go      # JWT工具
│   │       │
│   │       ├── validator/
│   │       │   └── validator.go # 参数校验
│   │       │
│   │       ├── constants/
│   │       │   └── constants.go # 常量
│   │       │
│   │       └── utils/
│   │           ├── time.go
│   │           ├── string.go
│   │           ├── captcha.go   # 验证码
│   │           └── ip.go
│   │
│   ├── Makefile
│   ├── go.mod
│   └── go.sum
│
├── web/                        # 前端项目
│   ├── src/
│   │   ├── api/                # API接口
│   │   ├── assets/             # 静态资源
│   │   ├── components/         # 公共组件
│   │   ├── layouts/            # 布局组件
│   │   ├── router/             # 路由配置
│   │   ├── stores/             # 状态管理
│   │   ├── styles/             # 全局样式
│   │   ├── utils/              # 工具函数
│   │   └── views/              # 页面视图
│   ├── index.html
│   ├── package.json
│   ├── tsconfig.json
│   └── vite.config.ts
│
├── scripts/
│   ├── build.sh
│   └── deploy.sh
│
├── LICENSE
└── README.md
```

## 技术栈

### 后端技术栈

- **语言**: Go
- **框架**: Gin
- **ORM**: GORM
- **权限**: Casbin
- **数据库**: MySQL / PostgreSQL / SQLite（可切换）
- **缓存**: Redis
- **存储**: 阿里云 OSS
- **日志**: Zap
- **配置**: Viper + fsnotify
- **文档**: OpenAPI 3.1
- **JWT**: golang-jwt

### 前端技术栈

- **框架**: Vue 3
- **构建**: Vite
- **路由**: Vue Router
- **状态**: Pinia
- **HTTP**: Axios

## 开发指南

### 后端开发

#### 初始化（go.mod 已存在则跳过）

```bash
cd server
go mod init github.com/chengjin/infra/server
```

#### 运行

```bash
make run
```

### 前端开发

#### 安装依赖

```bash
cd web
npm install
```

#### 开发

```bash
npm run dev
```

#### 构建

```bash
npm run build
```

## RESTful API 规范

```text
GET     /api/v1/users           # 获取用户列表
GET     /api/v1/users/:id       # 获取用户详情
POST    /api/v1/users           # 创建用户
PUT     /api/v1/users/:id       # 更新用户
DELETE  /api/v1/users/:id       # 删除用户
```

## 统一返回结构

```json
{
  "code": 0,
  "msg": "success",
  "data": {}
}
```

## 分页结构

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "list": [],
    "total": 100,
    "page": 1,
    "pageSize": 10
  }
}
```

## 错误码规范

| 错误码 | 说明           |
|--------|----------------|
| 0      | 成功           |
| 40000  | 请求参数错误   |
| 40100  | 未授权         |
| 40300  | 禁止访问       |
| 40001  | 用户不存在     |
| 40002  | 角色不存在     |
| 50000  | 内部服务器错误 |

## License

[Apache License 2.0](LICENSE)
