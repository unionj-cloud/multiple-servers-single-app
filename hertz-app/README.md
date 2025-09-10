# Hertz 应用示例

这是一个使用 CloudWeGo Hertz 框架构建的简单 Web 应用示例。

## 功能特性

- 🚀 基于 CloudWeGo Hertz 框架
- 📝 完整的 RESTful API 示例
- 🔍 健康检查接口
- 📊 性能分析支持 (pprof)
- 📋 结构化日志记录
- 🎯 用户管理 CRUD 操作

## 快速开始

### 1. 安装依赖

```bash
go mod tidy
```

### 2. 运行应用

```bash
go run main.go
```

应用将在 `http://localhost:8080` 启动。

### 3. 测试接口

#### 健康检查
```bash
curl http://localhost:8080/health
```

#### 获取用户列表
```bash
curl http://localhost:8080/api/v1/users
```

#### 根据ID获取用户
```bash
curl http://localhost:8080/api/v1/users/1
```

#### 创建用户
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john@example.com"}'
```

#### 更新用户
```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Jane Doe", "email": "jane@example.com"}'
```

#### 删除用户
```bash
curl -X DELETE http://localhost:8080/api/v1/users/1
```

#### Ping 接口
```bash
curl http://localhost:8080/api/v1/ping
```

#### Echo 接口
```bash
curl -X POST http://localhost:8080/api/v1/echo \
  -H "Content-Type: application/json" \
  -d '{"message": "Hello Hertz!"}'
```

## API 接口文档

### 基础接口

| 方法 | 路径 | 描述 |
|------|------|------|
| GET | `/` | 欢迎页面 |
| GET | `/health` | 健康检查 |

### 用户管理 API

| 方法 | 路径 | 描述 |
|------|------|------|
| GET | `/api/v1/users` | 获取用户列表 |
| GET | `/api/v1/users/:id` | 根据ID获取用户 |
| POST | `/api/v1/users` | 创建用户 |
| PUT | `/api/v1/users/:id` | 更新用户 |
| DELETE | `/api/v1/users/:id` | 删除用户 |

### 工具接口

| 方法 | 路径 | 描述 |
|------|------|------|
| GET | `/api/v1/ping` | Ping 接口 |
| POST | `/api/v1/echo` | Echo 接口 |

## 性能分析

应用集成了 pprof 性能分析工具，可以通过以下方式访问：

- 在浏览器中访问 `http://localhost:8080/debug/pprof/`
- 使用 `go tool pprof` 命令进行性能分析

## 项目结构

```
hertz-app/
├── main.go          # 主程序文件
├── go.mod           # Go 模块文件
└── README.md        # 项目说明文档
```

## 技术栈

- **框架**: CloudWeGo Hertz
- **日志**: zerolog
- **性能分析**: pprof
- **Go 版本**: 1.21+

## 开发说明

这个示例应用展示了如何使用 Hertz 框架构建一个完整的 Web 应用，包括：

1. 服务器配置和启动
2. 中间件集成（日志、性能分析）
3. 路由注册和分组
4. 请求处理和响应
5. 参数绑定和验证
6. 静态文件服务

你可以基于这个示例进行扩展，添加数据库连接、认证授权、缓存等功能。

