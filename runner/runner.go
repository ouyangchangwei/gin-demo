package runner

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
)

// Run 启动服务
func Run(router *gin.Engine, PORT string) error {
	// 初始化关闭信号通道shutdown
	shutdown := make(chan error)
	// 启动一个 goroutine，监听系统信号（如中断信号）来触发程序关闭
	go doShutdownOnSignal(shutdown)

	// 启动 HTTP 服务
	go func() {
		err := router.Run(fmt.Sprintf(":%s", PORT))
		// 如果服务出错，调用 doShutdown 将错误发送到 shutdown 通道
		doShutdown(shutdown, err)
	}()

	// 阻塞等待从 shutdown 通道接收到错误或关闭信号
	err := <-shutdown
	fmt.Println("Shutting down:", err)

	return err
}

func doShutdownOnSignal(shutdown chan<- error) {
	onSignal := make(chan os.Signal, 1)
	signal.Notify(onSignal, os.Interrupt, syscall.SIGTERM)
	// 阻塞等待，直到 onSignal 通道接收到信号
	sig := <-onSignal
	doShutdown(shutdown, fmt.Errorf("received signal %s", sig))
}

func doShutdown(shutdown chan<- error, err error) {
	select {
	case shutdown <- err:
		// 将错误信息发送到 shutdown 通道
	default:
		// 如果没有监听者正在等待接收 shutdown 通道消息，执行默认分支
		// 在这种情况下，什么也不做，因为关闭已经开始，可以忽略这些错误
	}
}
