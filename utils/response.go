package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(ctx *gin.Context, httpStatus int, code int, message string, data interface{}) {

	ctx.JSON(httpStatus, gin.H{
		"code": code,
		"msg":  message,
		"data": data,
	})
}

// Success 成功的请求
func Success(ctx *gin.Context, data interface{}) {
	Response(ctx, http.StatusOK, 1, "请求成功", data)
}

// Fail 失败的请求
func Fail(ctx *gin.Context, message string) {
	Response(ctx, http.StatusOK, 0, message, nil)
}
