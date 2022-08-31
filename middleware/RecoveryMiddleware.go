package middleware

import (
	"Essential/response"
	"fmt"

	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil { //recover用于“拦截”运行时恐慌的内建函数,防止程序崩溃
				response.Fail(ctx, nil, fmt.Sprint(err))
			}
		}()
		ctx.Next()
	}
}
