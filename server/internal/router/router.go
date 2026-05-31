package router

import (
	"github.com/gin-gonic/gin"

	"server/internal/middleware"
	"server/internal/shared/jwt"
	"server/internal/system/auth"
	"server/internal/system/loginlog"
	"server/internal/system/operlog"
	"server/internal/system/user"
)

// Register 注册所有路由
func Register(r *gin.Engine, jwtManager *jwt.JWT) {
	// 中间件
	r.Use(middleware.Recovery())
	r.Use(middleware.Cors())
	r.Use(middleware.Logger())

	// API路由组
	api := r.Group("/api/v1")
	{
		public := api.Group("")
		{
			auth.Register(public, jwtManager)
		}

		protected := api.Group("")
		protected.Use(middleware.Auth(jwtManager))
		{
			system := protected.Group("")
			{
				user.Register(system)
				operlog.Register(system)
				loginlog.Register(system)
			}
		}
	}
}
