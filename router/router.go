package router

import (
	"gin-demo/middleware"
	"github.com/gin-gonic/gin"
)

func Create() *gin.Engine {
	router := gin.Default()
	// 注册中间件
	router.Use(
		middleware.LoggerMiddleWare(),  // 日志
		middleware.RecoverMiddleWare(), // 异常的
	)
	// 配置全局路径
	router.Group("/api/v1")
	{
		router.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong"})
		})
	}
	return router
}
