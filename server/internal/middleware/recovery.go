package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/internal/shared/errno"
	"server/internal/shared/logger"
	"server/internal/shared/response"
)

// Recovery panic恢复中间件
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				stack := string(debug.Stack())

				logger.Error("panic recovered",
					zap.Any("error", err),
					zap.String("stack", stack),
				)

				response.Fail(c, errno.InternalError, fmt.Sprintf("内部错误: %v", err))
				c.AbortWithStatus(http.StatusOK)
			}
		}()

		c.Next()
	}
}