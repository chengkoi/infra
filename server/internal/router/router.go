package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "server/docs"
	"server/internal/middleware"
	"server/internal/shared/jwt"
	"server/internal/system/auth"
	"server/internal/system/user"
)

// Register 注册所有路由
func Register(r *gin.Engine, jwtManager *jwt.JWT) {
	// 中间件
	r.Use(middleware.Recovery())
	r.Use(middleware.Cors())
	r.Use(middleware.Logger())

	// Swagger 文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API路由组
	api := r.Group("/api/v1")
	{
		// 公开路由 - 不需要认证
		public := api.Group("")
		{
			auth.Register(public, jwtManager)
		}

		// 需要认证的路由
		protected := api.Group("")
		protected.Use(middleware.Auth(jwtManager))
		{
			// 系统模块
			system := protected.Group("")
			{
				user.Register(system)
				// TODO: dept.Register(system)
				// TODO: role.Register(system)
				// TODO: menu.Register(system)
				// TODO: dict.Register(system)
				// TODO: config.Register(system)
				// TODO: notice.Register(system)
			}

			// 监控模块
			// monitor := protected.Group("")
			// TODO: operlog.Register(monitor)
			// TODO: loginlog.Register(monitor)
			// TODO: online.Register(monitor)
		}
	}
}
