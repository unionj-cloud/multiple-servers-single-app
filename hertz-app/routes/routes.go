package routes

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"time"
)

func RegisterRoutes(h *server.Hertz) {
	// 健康检查路由
	h.GET("/health", func(ctx context.Context, c *app.RequestContext) {
		hlog.Info("health check")
		c.JSON(consts.StatusOK, utils.H{
			"status":  "ok",
			"message": "Hertz server is running",
			"time":    time.Now().Format(time.RFC3339),
		})
	})

	// 欢迎页面
	h.GET("/", func(ctx context.Context, c *app.RequestContext) {
		hlog.Info("hello world")
		c.JSON(consts.StatusOK, utils.H{
			"message": "Welcome to Hertz!",
			"version": "1.0.0",
			"docs":    "/docs",
		})
	})

	// API 路由组
	api := h.Group("/api/v1")
	{
		// 用户相关路由
		api.GET("/users", getUsers)
		api.GET("/users/:id", getUserByID)
		api.POST("/users", createUser)
		api.PUT("/users/:id", updateUser)
		api.DELETE("/users/:id", deleteUser)

		// 示例业务路由
		api.GET("/ping", ping)
		api.POST("/echo", echo)
	}

	// 静态文件服务
	h.StaticFS("/static", &app.FS{Root: "./static", IndexNames: []string{"index.html"}})
}

// 获取用户列表
func getUsers(ctx context.Context, c *app.RequestContext) {
	users := []map[string]interface{}{
		{"id": 1, "name": "Alice", "email": "alice@example.com"},
		{"id": 2, "name": "Bob", "email": "bob@example.com"},
		{"id": 3, "name": "Charlie", "email": "charlie@example.com"},
	}

	c.JSON(consts.StatusOK, utils.H{
		"data":  users,
		"count": len(users),
	})
}

// 根据ID获取用户
func getUserByID(ctx context.Context, c *app.RequestContext) {
	id := c.Param("id")

	// 模拟数据库查询
	user := map[string]interface{}{
		"id":    id,
		"name":  "User " + id,
		"email": "user" + id + "@example.com",
	}

	c.JSON(consts.StatusOK, utils.H{
		"data": user,
	})
}

// 创建用户
func createUser(ctx context.Context, c *app.RequestContext) {
	var user map[string]interface{}
	if err := c.BindAndValidate(&user); err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{
			"error": "Invalid request body",
		})
		return
	}

	// 模拟创建用户
	user["id"] = time.Now().Unix()
	user["created_at"] = time.Now().Format(time.RFC3339)

	c.JSON(consts.StatusCreated, utils.H{
		"message": "User created successfully",
		"data":    user,
	})
}

// 更新用户
func updateUser(ctx context.Context, c *app.RequestContext) {
	id := c.Param("id")

	var updateData map[string]interface{}
	if err := c.BindAndValidate(&updateData); err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{
			"error": "Invalid request body",
		})
		return
	}

	// 模拟更新用户
	updateData["id"] = id
	updateData["updated_at"] = time.Now().Format(time.RFC3339)

	c.JSON(consts.StatusOK, utils.H{
		"message": "User updated successfully",
		"data":    updateData,
	})
}

// 删除用户
func deleteUser(ctx context.Context, c *app.RequestContext) {
	id := c.Param("id")

	c.JSON(consts.StatusOK, utils.H{
		"message": "User " + id + " deleted successfully",
	})
}

// Ping 接口
func ping(ctx context.Context, c *app.RequestContext) {
	c.JSON(consts.StatusOK, utils.H{
		"message": "pong",
		"time":    time.Now().Format(time.RFC3339),
	})
}

// Echo 接口
func echo(ctx context.Context, c *app.RequestContext) {
	var data map[string]interface{}
	if err := c.BindAndValidate(&data); err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{
			"error": "Invalid request body",
		})
		return
	}

	c.JSON(consts.StatusOK, utils.H{
		"echo": data,
		"time": time.Now().Format(time.RFC3339),
	})
}
