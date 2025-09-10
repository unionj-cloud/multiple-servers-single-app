package main

import (
	"example/hertz-app/routes"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzZerolog "github.com/hertz-contrib/logger/zerolog"
	"github.com/libp2p/go-reuseport"
	"log"
	"net"
)

func main() {
	var listenConfig = net.ListenConfig{
		Control: reuseport.Control,
	}

	// 创建 Hertz 服务器
	h := server.Default(
		server.WithHostPorts(":6060"),
		server.WithMaxRequestBodySize(4*1024*1024), // 4MB
		server.WithListenConfig(&listenConfig),
		server.WithBasePath("hertz-app"),
	)

	hlog.SetLogger(hertzZerolog.New())

	// 注册路由
	routes.RegisterRoutes(h)

	// 启动服务器
	log.Println("Hertz server starting on :6060")
	h.Spin()
}
