package main

import (
	"fmt"
	"gin-demo/mode"
	"gin-demo/router"
	"gin-demo/runner"
	"os"
	"strconv"
)

var (
	// Mode the build mode.
	Mode = mode.Dev
	PORT = 8888
)

func main() {
	// 设置模式
	mode.Set(Mode)

	// 获取端口号
	port := strconv.Itoa(PORT)
	// 启动服务
	fmt.Println(fmt.Sprintf("HTTP服务启动:localhost:%s", port))

	// 创建路由
	engine := router.Create()

	// 启动
	if err := runner.Run(engine, port); err != nil {
		fmt.Println("Server error: ", err)
		os.Exit(1)
	}
}
