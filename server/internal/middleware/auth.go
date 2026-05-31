package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"

	"server/internal/shared/constants"
	"server/internal/shared/errno"
	"server/internal/shared/jwt"
	"server/internal/shared/response"
)

// Auth JWT认证中间件
func Auth(jwtManager *jwt.JWT) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Fail(c, errno.Unauthorized, "请先登录")
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Fail(c, errno.Unauthorized, "认证格式错误")
			c.Abort()
			return
		}

		claims, err := jwtManager.ParseToken(parts[1])
		if err != nil {
			response.Fail(c, errno.Unauthorized, "token已过期或无效")
			c.Abort()
			return
		}

		// 设置上下文
		c.Set(constants.ContextKeyUserID, claims.UserID)
		c.Set(constants.ContextKeyUsername, claims.Username)
		c.Set(constants.ContextKeyRoleID, claims.RoleID)

		c.Next()
	}
}