/*
*
User: cr-mao
Date: 2024/2/25 09:39
Email: crmao@qq.com
Desc: ctx_timeout.go
*/
package middlewares

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CtxTimeout(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 用超时context wrap request的context
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer func() {
			// 检查是否超时
			if ctx.Err() == context.DeadlineExceeded {
				c.Writer.WriteHeader(http.StatusGatewayTimeout)
				c.Abort()
			}
			//清理资源
			cancel()
		}()
		// 替换
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
